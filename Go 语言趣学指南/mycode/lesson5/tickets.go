package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("太空航行公司    飞行天数          飞行类型         价格（百万美元）")
	const secondsPerDay = 60 * 60 * 24
	const distance = 62100000
	company := ""
	trip := ""
	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 1:
			company = "Space Adventures"
		case 2:
			company = "SpaceX"
		case 3:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(16) + 15                  //速度
		price := 20.0 + speed                        // 价格
		duration := distance / speed / secondsPerDay // 天数
		if rand.Intn(2) == 1 {
			trip = "往返"
			price *= price
		} else {
			trip = "单程"
		}
		fmt.Printf("%-18v%-18v%-18v%-18v\n", company, duration, trip, price)
	}
}
