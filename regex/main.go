package main

import (
	"fmt"
	"regexp"
)

const emailAddress = `My email address is 1201220707@pku.edu.cn
645996795@qq.com
RuixuanHe@qq.com
RuiangHe@qq.com
`

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+.[a-zA-Z0-9]+`)
	// match := re.FindString(emailAddress)
	matchAddres := re.FindAllString(emailAddress, -1)
	for _, v := range matchAddres {
		fmt.Println(v)
	}
}
