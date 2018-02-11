package main

// Min 取小者
func Min(data1, data2 int64) int64 {
	if data1 < data2 {
		return data1
	}
	return data2
}

// Max 取大者
func Max(data1, data2 int64) int64 {
	if data1 > data2 {
		return data1
	}
	return data2
}

// IsNormal 判断是否正常
func IsNormal(data1, data2 int64) (is bool) {
	return data1 == Min(data1, data2)
}

func TestCase3(data1, data2 int64) {
	is := IsNormal(data1, data2)
	println(data1, data2, "=>", is)
}

func TestIsNormal() {
	TestCase3(111, 222)
	TestCase3(222, 111)
	TestCase3(333, 333)
}
