package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"mycode/chapter02/015.fatrate.refactor/calc"
)

/*
Go Module 使用
*/
func main() {
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "healthcheck",                      // 命令行的名字
		Short: "体重计算器，根据身高、体重、性别、年龄计算体制比，并给出健康建议", // 短描述
		Long:  "体重计算器......",                      // 详细描述
		// 注册回调函数
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)
			// 计算 BMI
			bmi, err := calc.CalcBMI(tall, weight)
			if err != nil {
				fmt.Println("计算异常...")
			}
			fatRate := calc.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate: ", fatRate)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	// 执行
	cmd.Execute()
}
