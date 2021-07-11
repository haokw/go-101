package main

import "fmt"

func main() {
	arr := []int{7, 5, 19, 8, 4, 1, 20, 13, 16}
	larr := len(arr)
	fmt.Println("init arr:", arr)

	heap := Build(arr)
	fmt.Println("build heap:", heap)

	heapp := Constructor(larr)
	for _, v := range arr {
		heapp.Push(v)
	}
	fmt.Println("heap push:", heapp)

	sarr := Sort(arr)
	fmt.Println("sarr:", sarr)
}

// 大顶堆
// 堆为完全二叉树 可用数组存储
// so'n = 2*root, 2*root+1
// root = child/2
type Heap struct {
	arr  []int
	cap  int
	size int
}

// 堆初始化
func Constructor(cap int) *Heap {
	heap := &Heap{
		cap: cap,
	}
	// 堆顶从 1 开始
	heap.arr = make([]int, cap+1)
	return heap
}

// 向堆中插入元素
func (this *Heap) Push(data int) {
	// 堆满退出
	if this.size == this.cap {
		return
	}

	// 新的数据放入末尾节点
	this.size++
	this.arr[this.size] = data

	// 自下而上 调整和根节点的位置
	for i := this.size; i/2 > 0 && this.arr[i] > this.arr[i/2]; i /= 2 {
		this.swap(i, i/2)
	}
}

// 交换根叶子节点元素
func (this *Heap) swap(child, root int) {
	this.arr[child], this.arr[root] = this.arr[root], this.arr[child]
}

// 堆顶弹出
func (this *Heap) Pop() (r int) {
	if this.size == 0 {
		r = -1
		return
	}
	r = this.arr[1]
	this.arr[1], this.arr[this.size] = this.arr[this.size], this.arr[1]
	this.size--
	this.heapfily(1)
	return
}

// 堆化调整 自上而下
func (this *Heap) heapfily(root int) {
	curr := root
	for {
		// 初始化最大节点为当前节点
		max := curr
		// 与左节点比较
		if left := curr * 2; left <= this.size && this.arr[curr] < this.arr[left] {
			max = left
		}
		// 与右节点比较
		if right := curr*2 + 1; right <= this.size && this.arr[max] < this.arr[right] {
			max = right
		}
		// 已满足堆条件 退出
		if max == curr {
			break
		}
		// 不满足条件 与最大子节点交换
		this.swap(max, curr)
		// 从子节点开始继续调整
		curr = max
	}
}

// 构建堆
func Build(arr []int) *Heap {
	larr := len(arr)
	heap := Constructor(larr)
	heap.arr = append([]int{0}, arr...)
	heap.size = larr

	// 从非叶子节点开始
	for i := heap.size / 2; i >= 1; i-- {
		heap.heapfily(i)
	}
	return heap
}

// 排序
func Sort(arr []int) (r []int) {
	heap := Build(arr)
	size := heap.size
	for i := 0; i < size; i++ {
		r = append(r, heap.Pop())
	}
	return
}
