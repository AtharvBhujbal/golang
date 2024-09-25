package main

// Channels
// 1. Channels hold data, anything from range, int or slices
// 2. Thread safe i.e. avoid data erases while reading and writing from databases
// 3. Listen for data i.e. block code execution until read or write data is done
// IMP Point
//	- When a val is set to a channel it blocks the code execution until the data is read from it
// ex:
// func example(){
//	var c = make(chan int)
//	c<-1 // a val 1 is set to channel, the code execution is blocked until the val is read from channel leading to "DEADLOCK"
//	var i = <- c // The code never reaches here
//	fmt.Println(i)
// }

// func main() {
// 	var c = make(chan int)
// 	go process(c)      // asyn call passes the code execution to next line
// 	for i := range c { // waits until the new val is set
// 		fmt.Println(i)
// 	}
// }

func process(c chan int) {
	defer close(c) // tells a compiler to close the channel just before exiting a function
	for i := 0; i < 5; i++ {
		c <- i
	}
}
