package main

import (
	"fmt"
	"math"
)

func StdDev(arr []int, avg float64) float64 {
	sumSqrd := 0.0
	for index := range arr {
		sumSqrd += math.Pow(float64(arr[index])-avg, 2)
	}
	return math.Sqrt(sumSqrd)
}

func Avg(arr []int) float64 {
	return float64(Sum(arr)) / float64(len(arr))
}

func Sum(arr []int) int {
	sum := 0
	for index := range arr {
		sum += arr[index]
	}
	return sum
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func FloatToString(float float64) string {
	return fmt.Sprintf("%f", float)
}

func IntToString(integer int64) string {
	return fmt.Sprintf("%d", integer)
}
