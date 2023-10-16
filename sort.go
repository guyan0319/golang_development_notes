package main

import (
	"fmt"
)

func main() {
	a := []int{28, 36, 45, 14, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	//b := bubbleSort(a)//冒泡
	//b := selectSort(a) //选择
	//b := insertSort(a) //插入
	//b := quickSort(a) //快速
	quick1Sort(a) //快速1
	fmt.Println(a)
	//fmt.Println(b)

}
func bubbleSort(arr []int) []int {
	var j, k, m int
	k = len(arr)
	for j = k - 1; j > 0; j-- {
		for m = 0; m < j; m++ {
			if arr[j] < arr[m] {
				arr[j], arr[m] = arr[m], arr[j]
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int {
	var i, j, k int
	k = len(arr)
	for j = 0; j < k-1; j++ {
		maxIndex := j
		for i = j; i < k; i++ {
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		//fmt.Println(arr)
	}
	return arr
}

// 插入排序
func insertSort(arr []int) []int {
	var i, j, k, temp int
	k = len(arr)
	for i = 1; i < k; i++ {
		temp = arr[i]
		for j = i - 1; j >= 0; j-- {
			fmt.Println(temp, arr[j])
			if temp >= arr[j] {
				break
			}
			arr[j+1], arr[j] = arr[j], arr[j+1]
			fmt.Println(arr)
		}
	}
	return arr
}
func quickSort(arr []int) []int {
	var i, k int
	k = len(arr)
	if k < 2 {
		return arr
	}
	var arrL, arrR []int
	for i = 1; i < k; i++ {
		if arr[0] > arr[i] {
			arrL = append(arrL, arr[i])
		} else {
			arrR = append(arrR, arr[i])
		}
	}
	arrL = quickSort(arrL)
	arrR = quickSort(arrR)
	arrL = append(arrL, arr[0])
	return append(arrL, arrR...)
}
func quick1Sort(arr []int){
	recursionSort(arr, 0, len(arr)-1)
}
func recursionSort(arr []int, left, right int) {
	if left<right {
		pivot:=partion(arr,left,right)
		recursionSort(arr,left,pivot-1)
		recursionSort(arr,pivot+1,right)
	}
}
func partion(nums []int, left, right int)  int {
	for left < right {
		for left < right && nums[left] <= nums[right] {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		for left < right && nums[left] <= nums[right] {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return left
}



