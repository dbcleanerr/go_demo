// 对字符串进行多次替换

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string = "['abc' '123']"

	replacer := strings.NewReplacer("[", "", "]", "", "'", "")

	output := replacer.Replace(input)
	fmt.Println(output)
}
