package sum

// Ints returns the sum of a list of integers
func Ints(vs ...int) int {
	// if len(vs) == 0 {
	// 	return 0
	// }
	// return Ints(vs[1:]...) + vs[0]

	return ints(vs)
}

// IntsArray bla comment
func IntsArray(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	// fmt.Print(vs[1:])
	return IntsArray(vs[1:]) + vs[0]
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}
