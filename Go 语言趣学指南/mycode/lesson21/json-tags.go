package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// ``可以将结构标签从原始字符串字面量改成普通字符串字面量
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"` // 打标签允许字段名是小写，可以正常被 JSON 编码
	}
	// 也可以使用 \ 转义
	//type location struct {
	//	Lat  float64 "json:\"latitude\""
	//	Long float64 "json:\"longitude\""
	//}

	curiosity := location{-4.5895, 137.4417}
	bytes, _ := json.Marshal(curiosity)
	fmt.Println(string(bytes)) //{"latitude":-4.5895,"longitude":137.4417}
}
