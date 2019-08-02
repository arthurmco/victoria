package victoria

import "fmt"

func main() {
	p := Parser{}
	p.Parse("a")

	fmt.Println("Hello world " + p.file.name)
}
