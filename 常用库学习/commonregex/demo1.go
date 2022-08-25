package main

import (
	"fmt"
	"github.com/mingrammer/commonregex"
)

// 内置常用的正则方法
// go get -u github.com/mingrammer/commonregex
// https://github.com/mingrammer/commonregex

func main() {
	//text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
	text2 := `John, to me by 5:00PM on 07-08-2022 abc 2022/08/01 abc Jan 9th 2012`
	text3 := `John, 192.168.1.11 on fe80::a8dd:1c26:ef68:e187/64 abc fe80::a8dd:1c26:ef68:e187 abc 192.168.1.11/32 2012 192.168.1.0/24 abc`

	dateList := commonregex.Date(text2)
	ipList := commonregex.IPs(text3)

	fmt.Println("date list:", dateList)
	fmt.Println("ip list:", ipList)
}
