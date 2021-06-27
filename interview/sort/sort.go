package main

import "fmt"

func main() {
	data := []int{4, 5, 6, 1, 3, 2}
	fmt.Println("sort before:", data)
	// bubbleSort(data)
	// insertionSort(data)
	quickSort(data)
	fmt.Println("sort after:", data)
}

// 冒泡排序
// 最好 O(n) 最坏 O(n^2) 平均 O(n^2)
// 逆序度 = 满有序度 - 有序度
func bubbleSort(data []int) {
	n := len(data)

	if n <= 1 {
		return
	}
	var flag bool
	for i := 0; i < n; i++ {
		// 提前退出
		flag = false
		for j := 0; j < n-i-1; j++ {
			// 稳定排序
			if data[j] > data[j+1] {
				// 原地排序
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 插入排序
// 最好 O(n) 最坏 O(n^2) 平均 O(n^2)
func insertionSort(data []int) {
	n := len(data)

	if n <= 1 {
		return
	}

	var value int
	var j int
	for i := 1; i < n; i++ {
		value = data[i]
		j = i - 1
		for ; j >= 0; j-- {
			// 稳定排序
			if data[j] > value {
				// 原地排序
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = value
	}
}

// 选择排序
// 不稳定排序
// 最好 O(n) 最坏 O(n^2) 平均 O(n^2)
func selectionSort(data []int) {

}

// 归并排序

// 快速排序
func quickSort(data []int) {
	sortScope(data, 0, len(data)-1)
}

// 分区排序
func sortScope(data []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := partition(data, start, end)
	sortScope(data, start, pivot-1)
	sortScope(data, pivot+1, end)
}

// 分区
func partition(data []int, start int, end int) int {
	ref := data[end]
	pivot := start
	for i := start; i < end; i++ {
		if data[i] < ref {
			data[pivot], data[i] = data[i], data[pivot]
			pivot++
		}
	}
	data[pivot], data[end] = data[end], data[pivot]
	return pivot
}

// 如何在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素？

// 桶排序 Bucket sort

// 如果你所在的省有 50 万考生，如何通过成绩快速排序得出名次呢
// 计数排序 Counting sort

// 假设我们有 10 万个手机号码，希望将这 10 万个手机号码从小到大排序
// 基数排序 Radix sort

// 如何根据年龄给 100 万用户排序

// 堆排序 Heap sort
// 优先级队列、求 Top K 和求中位数
