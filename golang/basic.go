package main

import "fmt"

func main() {
	var num int = 123

	fmt.Printf("number is %d\n", num)

	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	fmt.Printf("number is %d\n", Monday)

	var bytenum uint8 = 255

	fmt.Println("num is ", bytenum)

	a := "Hello, world"

	for i, c := range a {
		fmt.Println(i, string(c))
	}

	fmt.Println("a len", len(a))

	b := `hiii` + a
	fmt.Println(b)

	ctable := "user"
	c := fmt.Sprintf("select * from %s", ctable)
	fmt.Println(c)

	d := len(b) > len(a)
	fmt.Println(d)

	e := [3]int{1,2,3}
	fmt.Println(e)

	f := [...]int{1,2,3,4}
	fmt.Println(f)

	g := [2][3]int{
		{1,2},
		{3,4,5},
	}
	fmt.Println(g)

	const h [2]int = [2]int{1,2}
	h[1] = 3
}
