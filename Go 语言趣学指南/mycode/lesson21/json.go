package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64 // 字段必须以大写字母开头，否则 JSON 编码结果将会是 {}
	}

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	ExitOnError(err)
	fmt.Println(string(bytes)) //{"Lat":-4.5895,"Long":137.4417}
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
