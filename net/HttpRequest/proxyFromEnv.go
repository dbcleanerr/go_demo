// 使用环境变量来使用proxy
// HTTPS_PROXY, HTTP_PROXY
// go get github.com/kirinlabs/HttpRequest

package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
	"os"
)

func main() {
	// 如果访问的是 https 就要用 https_proxy 环境变量
	//os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	req := HttpRequest.NewRequest().SetTimeout(30).Debug(false)

	resp, err := req.Get("http://192.168.1.91:9971/t01", map[string]interface{}{
		"id": "eq.1",
	})
	defer resp.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := resp.Body()

	fmt.Println(err, string(body))

}
