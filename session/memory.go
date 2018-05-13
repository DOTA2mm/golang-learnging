package memory

import (
	"time"
)

var pder = &Provider{list: list.new()}

type SessionStore struct {
	sid          string                      // session id 唯一标识
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} // session 里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}
