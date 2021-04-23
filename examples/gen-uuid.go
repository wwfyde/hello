package main

func a() int {

	panic(a())

	return 0

}

func main() {
	a()


}