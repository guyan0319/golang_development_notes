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
	//quick1Sort(a) //快速1
	Quick3Sort(a,0,len(a)-1) //快速3
	fmt.Println(a)
	//fmt.Println(b)

}
//冒泡排序
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
//选择排序
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


// 第三种写法 快速排序
func Quick3Sort(a []int,left int, right int)  {
	if left >= right {
		return
	}
	mid := left
	for i := left + 1; i <= right ; i++ {
		if a[left] >= a[i]{
			//分割位定位++
			mid ++;
			a[i],a[mid] = a[mid],a[i]
		}
	}
	//起始位和分割位
	a[left], a[mid] = a[mid],a[left]

	Quick3Sort(a,left,mid - 1)
	Quick3Sort(a,mid + 1,right)

}
//取乘积最大子数组
func maxGroup(nums []int) int {
	var res, maxvalue,minValue int
	res = nums[0]
	maxvalue=nums[0]
	minValue=nums[0]

	for index := 1; index < len(nums); index++ {
		if nums[index] == 0 {
			maxvalue = 0
			minValue = 0
		} else {
			var tmpMax int
			var tmpMin int
			if nums[index] > 0 {
				tmpMax = nums[index] * maxvalue
				tmpMin = nums[index] * minValue
			} else {
				tmpMax = nums[index] *minValue
				tmpMin = nums[index] *maxvalue
			}
			if tmpMax > nums[index] {
				maxvalue = tmpMax
			} else {
				maxvalue = nums[index]
			}
			if tmpMin < nums[index] {
				minValue = tmpMin
			} else {
				minValue = nums[index]
			}
		}
		if res < maxvalue {
			res =maxvalue
		}
	}
	return res
}
//最大连续子数组和
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for _, num := range nums[1:] {
		if currentSum >= 0 {
			currentSum += num
		} else {
			currentSum = num
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

//最大连续不重复子数组  返回长度
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s){
		// 判断KEY是否存在
		if lastI,ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI + 1
		}
		if lastOccurred[ch] >= start{
			start = lastOccurred[ch] + 1
		}

		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
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



