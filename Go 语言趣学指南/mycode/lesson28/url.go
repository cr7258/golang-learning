package main

import (
	"fmt"
	"net/url"
	"os"
)

func urlParse(s string) {
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println(u)           // <nil>
		fmt.Println(err)         // parse "https://a b.com": invalid character " " in host name
		fmt.Printf("%#v\n", err) // &url.Error{Op:"parse", URL:"https://a b.com", Err:" "}
	}

	// 对错误执行 *url.Error 类型断言以便访问并打印错误的底层字段
	if e, ok := err.(*url.Error); ok {
		fmt.Println(e.Op)
		fmt.Println(e.URL)
		fmt.Println(e.Err)
	}
	os.Exit(1)

}

func main() {
	s := "https://a b.com"
	urlParse(s)
}
