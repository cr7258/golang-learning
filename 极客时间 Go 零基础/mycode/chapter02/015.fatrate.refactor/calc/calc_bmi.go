package calculator

import (
	"fmt"
)

// 计算体脂率
func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return 0, fmt.Errorf("身高不能是0或者负数")
	}
	// todo 验证体重的合法性
	return weight / (tall * tall), nil
}
