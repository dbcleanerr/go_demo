package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

type Option struct {
	//D:\git\go\go_demo\常用库学习\go-flags-demo>go_flags_demo.exe -n 张三丰 -H 打架 -H 喝酒 --birthday=20220801 --score=语文:80 --score=数学:70
	//name:张三丰
	//hobby:[打架 喝酒]
	//birthday:20220801
	//birthday:map[string]int{"数学":70, "语文":80}

	// required: 是否必须
	// default: 默认值

	// 基本类型
	Name string `short:"n" long:"name" description:"姓名" required:"true"`
	// 切片, h 是帮助
	Hobby []string `short:"H" long:"hobby" description:"爱好"`
	// 函数类型
	Birthday func(string) `long:"birthday" description:"出生日期"`
	// map类型，key只能是字符串
	Score map[string]int `long:"score" description:"各科成绩"`
}

func main() {
	var opt Option
	var birthday string
	opt.Birthday = func(s string) {
		birthday = s
	}

	_, err := flags.Parse(&opt)
	if err != nil {
		log.Fatal("parse error:", err)
	}

	fmt.Printf("name:%s\n", opt.Name)
	fmt.Printf("hobby:%s\n", opt.Hobby)
	fmt.Printf("birthday:%s\n", birthday)
	fmt.Printf("birthday:%#v\n", opt.Score)
}
