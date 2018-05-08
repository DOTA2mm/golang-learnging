package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"text/template"
	"time"
)

// 全局的 session 管理器
type Manager struct {
	cookieName  string
	lock        sync.Mutex // 唯一的 sessionid
	provider    Provider
	maxLifeTime int64
}

// Manager Method: 创建全局唯一的 sessionID
func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Manager Method: 用来检测是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()         // 使用互斥锁保证互斥执行
	defer manager.lock.Unlock() // 解锁
	cookie, err := r.Cookie(manager.cookieName)
	// 没找到来自客户端的 sessionID
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		// 服务器创建一个 session
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// Manager Method: Destroy sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// Manager Method: session 销毁
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

// 抽象出一个Provider接口，用以表征session管理器底层存储结构
type Provider interface {
	// SessionInit函数实现Session的初始化，操作成功则返回此新的Session变量
	SessionInit(sid string) (Session, error)
	// SessionRead函数返回sid所代表的Session变量，如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
	SessionRead(sid string) (Session, error)
	// SessionDestroy函数用来销毁sid对应的Session变量
	SessionDestroy(sid string) error
	// SessionGC根据maxLifeTime来删除过期的数据
	SessionGC(maxLifeTime int64)
}

// Session 接口实现对 session 的基本操作
type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}

// 用来随需注册存储session的结构的Register函数的实现
var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

var globalSessions *session.Manager

func init() {
	// 初始化一个全局的session 管理器
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}

// 使用 session
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles()
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
