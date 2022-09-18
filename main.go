package main

func main() {
	square := square{}
	square.sideLength = 5

	triangle := triangle{}
	triangle.base = 10
	triangle.height = 5

	printArea(square)
	printArea(triangle)
}
