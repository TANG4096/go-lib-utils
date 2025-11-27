package data_structure

import (
	"testing"
)

func TestNewBIT(t *testing.T) {
	// 测试正常创建
	bit := NewBIT(10)
	if bit == nil {
		t.Fatal("NewBIT should not return nil")
	}
	if bit.n != 10 {
		t.Errorf("Expected n = 10, got %d", bit.n)
	}
	if len(bit.tree) != 11 {
		t.Errorf("Expected tree length = 11, got %d", len(bit.tree))
	}

	// 测试创建大小为0的树状数组
	bit0 := NewBIT(0)
	if bit0 == nil {
		t.Fatal("NewBIT(0) should not return nil")
	}
	if bit0.n != 0 {
		t.Errorf("Expected n = 0, got %d", bit0.n)
	}

	// 测试负数大小
	bitNeg := NewBIT(-5)
	if bitNeg == nil {
		t.Fatal("NewBIT(-5) should not return nil")
	}
	if bitNeg.n != 0 {
		t.Errorf("Expected n = 0 for negative input, got %d", bitNeg.n)
	}
}

func TestBITUpdate(t *testing.T) {
	bit := NewBIT(10)

	// 测试正常更新
	bit.Update(1, 5)
	if bit.Get(1) != 5 {
		t.Errorf("Expected value at index 1 = 5, got %d", bit.Get(1))
	}

	bit.Update(3, 10)
	if bit.Get(3) != 10 {
		t.Errorf("Expected value at index 3 = 10, got %d", bit.Get(3))
	}

	// 测试多次更新同一位置
	bit.Update(1, 3)
	if bit.Get(1) != 8 {
		t.Errorf("Expected value at index 1 = 8, got %d", bit.Get(1))
	}

	// 测试边界索引更新
	bit.Update(0, 100)  // 无效索引，不应更新
	bit.Update(11, 100) // 超出范围，不应更新
	if bit.Query(10) != 18 {
		t.Errorf("Expected total sum = 18, got %d", bit.Query(10))
	}
}

func TestBITQuery(t *testing.T) {
	bit := NewBIT(5)

	// 初始化值
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)

	// 测试前缀和查询
	tests := []struct {
		index    int
		expected int64
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
	}

	for _, tc := range tests {
		result := bit.Query(tc.index)
		if result != tc.expected {
			t.Errorf("Query(%d): expected %d, got %d", tc.index, tc.expected, result)
		}
	}

	// 测试边界情况
	if bit.Query(0) != 0 {
		t.Errorf("Query(0) should return 0, got %d", bit.Query(0))
	}
	if bit.Query(-1) != 0 {
		t.Errorf("Query(-1) should return 0, got %d", bit.Query(-1))
	}
	if bit.Query(10) != 15 {
		t.Errorf("Query(10) should return 15 (clamped to n), got %d", bit.Query(10))
	}
}

func TestBITRangeQuery(t *testing.T) {
	bit := NewBIT(5)

	// 初始化值
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)

	// 测试区间和查询
	tests := []struct {
		left     int
		right    int
		expected int64
	}{
		{1, 1, 1},
		{1, 3, 6},
		{2, 4, 9},
		{3, 5, 12},
		{1, 5, 15},
		{4, 4, 4},
	}

	for _, tc := range tests {
		result := bit.RangeQuery(tc.left, tc.right)
		if result != tc.expected {
			t.Errorf("RangeQuery(%d, %d): expected %d, got %d", tc.left, tc.right, tc.expected, result)
		}
	}

	// 测试无效区间
	if bit.RangeQuery(4, 2) != 0 {
		t.Errorf("RangeQuery(4, 2) should return 0 for invalid range, got %d", bit.RangeQuery(4, 2))
	}
}

func TestBITSetAndGet(t *testing.T) {
	bit := NewBIT(5)

	// 测试设置值
	bit.Set(1, 10)
	if bit.Get(1) != 10 {
		t.Errorf("Expected Get(1) = 10, got %d", bit.Get(1))
	}

	bit.Set(3, 30)
	if bit.Get(3) != 30 {
		t.Errorf("Expected Get(3) = 30, got %d", bit.Get(3))
	}

	// 测试覆盖设置
	bit.Set(1, 100)
	if bit.Get(1) != 100 {
		t.Errorf("Expected Get(1) = 100 after override, got %d", bit.Get(1))
	}

	// 验证前缀和仍然正确
	if bit.Query(3) != 130 {
		t.Errorf("Expected Query(3) = 130, got %d", bit.Query(3))
	}

	// 测试边界索引
	bit.Set(0, 50)  // 无效索引
	bit.Set(6, 50)  // 超出范围
	if bit.Get(0) != 0 {
		t.Errorf("Get(0) should return 0 for invalid index, got %d", bit.Get(0))
	}
	if bit.Get(6) != 0 {
		t.Errorf("Get(6) should return 0 for out of range index, got %d", bit.Get(6))
	}
}

func TestBITEdgeCases(t *testing.T) {
	// 测试大索引
	bit := NewBIT(1000)
	bit.Update(1000, 12345)
	if bit.Get(1000) != 12345 {
		t.Errorf("Expected Get(1000) = 12345, got %d", bit.Get(1000))
	}

	// 测试负数值
	bit2 := NewBIT(5)
	bit2.Update(1, -10)
	if bit2.Get(1) != -10 {
		t.Errorf("Expected Get(1) = -10, got %d", bit2.Get(1))
	}
	bit2.Update(2, 20)
	if bit2.Query(2) != 10 {
		t.Errorf("Expected Query(2) = 10, got %d", bit2.Query(2))
	}

	// 测试设置负值
	bit2.Set(3, -5)
	if bit2.Get(3) != -5 {
		t.Errorf("Expected Get(3) = -5, got %d", bit2.Get(3))
	}

	// 测试大数值
	bit3 := NewBIT(10)
	var largeValue int64 = 9223372036854775800 // 接近int64最大值
	bit3.Update(1, largeValue)
	if bit3.Get(1) != largeValue {
		t.Errorf("Expected Get(1) = %d, got %d", largeValue, bit3.Get(1))
	}

	// 测试空树状数组
	bit4 := NewBIT(0)
	bit4.Update(1, 10) // 不应崩溃
	if bit4.Query(1) != 0 {
		t.Errorf("Expected Query(1) = 0 for empty BIT, got %d", bit4.Query(1))
	}
}

func TestBITString(t *testing.T) {
	bit := NewBIT(5)
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)

	str := bit.String()
	expected := "BIT{n: 5, values: [1, 2, 3, 4, 5]}"
	if str != expected {
		t.Errorf("Expected String() = %q, got %q", expected, str)
	}

	// 测试空树状数组的字符串表示
	bit2 := NewBIT(0)
	str2 := bit2.String()
	expected2 := "BIT{n: 0, values: []}"
	if str2 != expected2 {
		t.Errorf("Expected String() = %q, got %q", expected2, str2)
	}
}

func TestBITComprehensive(t *testing.T) {
	// 综合测试
	bit := NewBIT(10)

	// 依次更新
	for i := 1; i <= 10; i++ {
		bit.Update(i, int64(i*10))
	}

	// 验证所有值
	for i := 1; i <= 10; i++ {
		if bit.Get(i) != int64(i*10) {
			t.Errorf("Get(%d): expected %d, got %d", i, i*10, bit.Get(i))
		}
	}

	// 验证前缀和
	expectedSum := int64(0)
	for i := 1; i <= 10; i++ {
		expectedSum += int64(i * 10)
		if bit.Query(i) != expectedSum {
			t.Errorf("Query(%d): expected %d, got %d", i, expectedSum, bit.Query(i))
		}
	}

	// 测试区间和
	// [1,5] = 10+20+30+40+50 = 150
	if bit.RangeQuery(1, 5) != 150 {
		t.Errorf("RangeQuery(1, 5): expected 150, got %d", bit.RangeQuery(1, 5))
	}
	// [6,10] = 60+70+80+90+100 = 400
	if bit.RangeQuery(6, 10) != 400 {
		t.Errorf("RangeQuery(6, 10): expected 400, got %d", bit.RangeQuery(6, 10))
	}

	// 修改后验证
	bit.Set(5, 500)
	if bit.Get(5) != 500 {
		t.Errorf("After Set(5, 500), Get(5) should be 500, got %d", bit.Get(5))
	}
	// [1,5] = 10+20+30+40+500 = 600
	if bit.RangeQuery(1, 5) != 600 {
		t.Errorf("After Set, RangeQuery(1, 5): expected 600, got %d", bit.RangeQuery(1, 5))
	}
}
