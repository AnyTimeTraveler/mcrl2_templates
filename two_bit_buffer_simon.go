package main

func main() {
	var input = make(chan bool)
	var output = make(chan bool)

	tbb(input, output)

	println("should not deadlock:")
	input <- false
	input <- true
	println(<-output)
	input <- false

	println("emptying buffer:")
	println(<-output)
	println(<-output)

	println("\n\n\nshould deadlock:")
	input <- false
	input <- true
	println("deadlock after this line")
	input <- false

}

func obb(in <-chan bool, out chan<- bool) {
	for {
		out <- <-in
		//	laenger:
		var bit = <-in
		out <- bit
	}
}

func tbb(input <-chan bool, output chan<- bool) {
	var transfer = make(chan bool)
	go obb(input, transfer)
	go obb(transfer, output)
}
