package main

/*
* 变长参数
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.3.md
 */

// minimum 取最小数
func minimum(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
