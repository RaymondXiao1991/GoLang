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

func TestCase2(data1, data2 int64) {
	is := IsNormal(data1, data2)
	println(data1, data2, "=>", is)
}

func TestIsNormal() {
	TestCase2(111, 222)
	TestCase2(222, 111)
	TestCase2(333, 333)
}
