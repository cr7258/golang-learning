package main

import "fmt"

func main() {
	avgBmi := calculateAvg(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(avgBmi)
	avgBmi = calculateAvgOnSlice([]float64{3, 4, 3, 2, 1})
	fmt.Println(avgBmi)
	getScoresOfStudent("Tom")
}

// ... 表示不定长参数
// 命名返回值
func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent(stdName string) (chinese int, match int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
