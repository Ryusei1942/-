package main

import "fmt"

type Vertex struct {
	X, Y int
}

// Vertexのstructと紐づけている
func (v Vertex) Area() int {
	return v.X * v.Y
}

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	// vの呼び出し関数として利用可能
	fmt.Println(v.Area())
}
