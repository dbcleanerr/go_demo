package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)

func main() {
	retries := 3
	req := HttpRequest.NewRequest().SetTimeout(30).Debug(false)

	for retries > 0 {
		res, err := req.Get("https://www.baidu.com")
		defer res.Close()
		if err != nil {
			log.Printf("get error, %s", err)
			retries--
		} else {
			break
		}

		body, err := res.Body()
		fmt.Println(err, string(body))
	}
}
