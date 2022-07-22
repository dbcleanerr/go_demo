// 把字符串解析到数组
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	jsonstring := "[0, 1, 2, 3, 4]"
	var listoflists []int
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	err := dec.Decode(&listoflists)
	fmt.Println(err, listoflists)
}
