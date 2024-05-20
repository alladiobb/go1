package main

func main() {
	defer println("1")
	println("2")
	defer println("3")
}
