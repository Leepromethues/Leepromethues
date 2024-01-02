package main

import "fmt"

var i = 0

func autoCount() {
	i++
	fmt.Println(i)
}
func main() {
	autoCount()
	autoCount()
	autoCount()

}
