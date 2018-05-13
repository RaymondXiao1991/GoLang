package main

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func Sum(vals...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

