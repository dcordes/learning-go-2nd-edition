-- the full slice expression protects against append
	x1 := make([]string, 0, 5)
	x1 = append(x1, "a", "b", "c", "d")
	y1 := x1[:2:2] // NB
	z1 := x1[2:4:4] // NB
	//y1 := x1[:2]
	//z1 := x1[2:]
	fmt.Println(cap(x1), cap(y1), cap(z1))
	y1 = append(y1, "i", "j", "k")
	x1 = append(x1, "x")
	z1 = append(z1, "y")
	fmt.Println("x:", x1)
	fmt.Println("y:", y1)
	fmt.Println("z:", z1)

-- using a map of booleans to create a "set"
intSet := map[int]bool{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
    intSet[v] = true
}
fmt.Println(len(vals), len(intSet))
fmt.Println(intSet[5])
fmt.Println(intSet[500])
if intSet[100] {
    fmt.Println("100 is in the set")
}
