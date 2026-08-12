package main

import (
	"fmt"
)

func ArrayAndSlices(numbers [5]int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func ArrayAndSlicesWithAdv(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum+=num
	}
	return  sum
}

func CaulculateAverageOfArrayItems(mySlice []int) int {
	avg := 0
	for _, num := range mySlice {
		avg+=num
	}
	return (avg/len(mySlice))
}

func SumAll(numbers...[]int)[]int{
	lengthOfNumbers := len(numbers)
	sumAll := make([]int, lengthOfNumbers)
	for i,num := range numbers{
		sumAll[i] = Sum(num)
	}
	return  sumAll
}

func  Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers{
		sum+=num
	}
	return  sum
}

func SumAllTails(numbers...[]int)[]int{
	var tailSum []int
	for _,num := range numbers{
		if len(num) == 0{
			tailSum = append(tailSum, 0)
		}else{
			tail := num[1:]
			tailSum = append(tailSum, Sum(tail))
		}
	}
	return tailSum
}

func main(){
	fmt.Println("trying something with arrays")
}