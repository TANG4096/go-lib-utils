package data_structure

import (
	"fmt"
	"strings"
)

// BIT 树状数组(Binary Indexed Tree/Fenwick Tree)
// 用于高效计算数列的前缀和，区间和
// 注意：索引从1开始
type BIT struct {
	tree []int64
	n    int
}

// NewBIT 创建大小为n的树状数组
// n为数组大小，有效索引范围为[1, n]
func NewBIT(n int) *BIT {
	if n < 0 {
		n = 0
	}
	return &BIT{
		tree: make([]int64, n+1),
		n:    n,
	}
}

// lowbit 返回x的最低位1所代表的值
func lowbit(x int) int {
	return x & (-x)
}

// Update 在index位置增加delta值
// index从1开始，范围为[1, n]
func (b *BIT) Update(index int, delta int64) {
	if index < 1 || index > b.n {
		return
	}
	for i := index; i <= b.n; i += lowbit(i) {
		b.tree[i] += delta
	}
}

// Query 查询前缀和[1, index]
// index从1开始，范围为[1, n]
func (b *BIT) Query(index int) int64 {
	if index < 1 {
		return 0
	}
	if index > b.n {
		index = b.n
	}
	var sum int64 = 0
	for i := index; i > 0; i -= lowbit(i) {
		sum += b.tree[i]
	}
	return sum
}

// RangeQuery 查询区间和[left, right]
// left和right从1开始，范围为[1, n]
func (b *BIT) RangeQuery(left, right int) int64 {
	if left > right {
		return 0
	}
	return b.Query(right) - b.Query(left-1)
}

// Set 设置index位置的值为value
// index从1开始，范围为[1, n]
func (b *BIT) Set(index int, value int64) {
	if index < 1 || index > b.n {
		return
	}
	oldValue := b.Get(index)
	delta := value - oldValue
	b.Update(index, delta)
}

// Get 获取index位置的值
// index从1开始，范围为[1, n]
func (b *BIT) Get(index int) int64 {
	if index < 1 || index > b.n {
		return 0
	}
	return b.RangeQuery(index, index)
}

// String 返回树状数组的字符串表示
func (b *BIT) String() string {
	var sb strings.Builder
	sb.WriteString("BIT{")
	sb.WriteString(fmt.Sprintf("n: %d, ", b.n))
	sb.WriteString("values: [")
	for i := 1; i <= b.n; i++ {
		if i > 1 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", b.Get(i)))
	}
	sb.WriteString("]}")
	return sb.String()
}
