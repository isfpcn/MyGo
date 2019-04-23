package main

import (
	"fmt"
	"regexp"
)

const text = `my eamil is lnmpson_@gamil.com.cn
	dsa@qq.com
isfp@cc.org      
`

func main() {

	//regex, e := regexp.Compile("lnmpsong@gamil.com")
	//if e != nil {
	//	panic(e)
	//}

	//邮箱
	regex := regexp.MustCompile("([a-zA-Z0-9_]+)@([a-zA-z0-9]+)\\.([a-zA-Z0-9.]+)")

	//之匹配第一个
	s := regex.FindString(text)

	//多个匹配 第二个参数 为 -1 匹配所有
	allString := regex.FindAllString(text, -1)

	//取（）内的 内容
	submatch := regex.FindAllStringSubmatch(text, -1)

	fmt.Println(s)

	fmt.Println(allString)

	fmt.Println(submatch)

}
