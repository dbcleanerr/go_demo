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
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")
	req := HttpRequest.NewRequest().SetTimeout(30).Debug(false)

	resp, err := req.Get("https://www.google.com")
	defer resp.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := resp.Body()

	fmt.Println(err, string(body))

}
