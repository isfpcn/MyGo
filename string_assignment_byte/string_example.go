package main

import "fmt"

func main() {
	var b [64]byte
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// 这么赋值肯定是错误的 类型不同
	//b = s

	//遍历赋值 这种是可以的
	//for i := 0; i < len(s); i++ {
	//	b[i] = s[i]
	//}

	//用copy方式，前提是必须指定[]byte的len
	copy(b[:], s)

	fmt.Println(b)

	name := "Hello，中国。"

	length := len(name)

	bb := make([]byte, length)

	copy(bb, name)

	fmt.Println(string(bb))
}
