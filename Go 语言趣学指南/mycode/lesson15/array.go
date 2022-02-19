package main

func main() {
	var planets [8]string // 声明长度为 8 的 string 类型的数组
	//planets[8] = "Pluto" // 编译错误
	i := 8
	planets[i] = "Pluto" //运行时错误 panic
}
