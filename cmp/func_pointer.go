package cmp

func FuncEquals(p1 any, p2 any) bool {
	if p1 == nil {
		return p2 == nil
	}
	return p1 == p2
	// if reflect.ValueOf(function1).Pointer() == reflect.ValueOf(function2).Pointer() {
	// 	fmt.Println("function1 and function2 are equal")
	// } else {
	// 	fmt.Println("function1 and function2 are not equal")
	// }
}
