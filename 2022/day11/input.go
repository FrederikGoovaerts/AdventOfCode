package main

func getEx1() []Monkey {
	return []Monkey{
		{[]int{79, 98}, func(in int) int { return in * 19 }, 23, 2, 3},
		{[]int{54, 65, 75, 74}, func(in int) int { return in + 6 }, 19, 2, 0},
		{[]int{79, 60, 97}, func(in int) int { return in * in }, 13, 1, 3},
		{[]int{74}, func(in int) int { return in + 3 }, 17, 0, 1},
	}
}

func getInput() []Monkey {
	return []Monkey{
		{[]int{75, 63}, func(in int) int { return in * 3 }, 11, 7, 2},
		{[]int{65, 79, 98, 77, 56, 54, 83, 94}, func(in int) int { return in + 3 }, 2, 2, 0},
		{[]int{66}, func(in int) int { return in + 5 }, 5, 7, 5},
		{[]int{51, 89, 90}, func(in int) int { return in * 19 }, 7, 6, 4},
		{[]int{75, 94, 66, 90, 77, 82, 61}, func(in int) int { return in + 1 }, 17, 6, 1},
		{[]int{53, 76, 59, 92, 95}, func(in int) int { return in + 2 }, 19, 4, 3},
		{[]int{81, 61, 75, 89, 70, 92}, func(in int) int { return in * in }, 3, 0, 1},
		{[]int{81, 86, 62, 87}, func(in int) int { return in + 8 }, 13, 3, 5},
	}
}
