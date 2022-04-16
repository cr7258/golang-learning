module mycode

go 1.18

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.4.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

// 使用自己扩展的包，本地替换
replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
// github 上的替换
// replace github.com/armstrongli/go-bmi => github.com/armstrongli/go-bmi v0.0.2
