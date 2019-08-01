
package main

import "fmt"

func main() {
	p := Parser{}
	p.Parse()
	
	fmt.Println("Hello world " + p.data)
}
