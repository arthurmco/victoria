package victoria

import "testing"

func TestParserParseHelloWorld(t *testing.T) {

	str := `
fn main() {
	println('hello world')
}`

	parser := Parser{}
	r := parser.Parse(str)


	if r == false {
		t.Errorf("Parser could not parse a valid file")
	}

	if _, ok := parser.file.definedFunctions["main"]; !ok {
		t.Fatalf("Parser could not parse function 'main'")
	}

	if parser.file.definedFunctions["main"].name != "main" {
		t.Errorf("Parser could not parse function 'main'")
	}
	
	if len(parser.file.definedFunctions["main"].args) != 0 {
		t.Errorf("Parser could not get the args of 'main'")
	}

	if parser.file.definedFunctions["main"].ret.name != "" {
		t.Errorf("Parser could not get the return type of 'main'")
	}
}

type FunctionProp struct {
	fname string
	fargs int
	fret string
}

func TestParserParseMultipleFunctions(t *testing.T) {

	str := `
fn main() {
	println(add(77, 33))
	println(neg(100))
}

fn add(x int, y int) int {
	return x + y
}

fn neg(y int) int {
	return -y
}`

	parser := Parser{}
	r := parser.Parse(str)

	if r == false {
		t.Errorf("Parser could not parse a valid file")
	}

	props := [3]FunctionProp{
		FunctionProp{"main", 0, "void"}, FunctionProp{"add", 2, "int"}, FunctionProp{"neg", 1, "int"},
	}

	for i := 0; i < 3; i++ {
		if _, ok := parser.file.definedFunctions[props[i].fname]; !ok {
			t.Fatalf("Parser could not parse function '%s'", props[i].fname)
		}

		if parser.file.definedFunctions[props[i].fname].name != props[i].fname {
			t.Errorf("Parser could not parse function '%s'", props[i].fname)
		}
		
		if len(parser.file.definedFunctions[props[i].fname].args) != props[i].fargs {
			t.Errorf("Parser could not get the args of '%s'", props[i].fname)
		}

		if parser.file.definedFunctions[props[i].fname].ret.name != props[i].fret {
			t.Errorf("Parser could not get the return type of '%s'", props[i].fname)
		}
	}

	
}


func TestParserParseMultipleFunctionsAndComments(t *testing.T) {

	str := `

// Main function
fn main() {
	println(add(77, 33))
	println(sub(100, 50))
}

/* Another function
   with a multiline
   comment */
fn add(x int, y int) int {
	return x + y
}

// Main function
// with two comment lines
fn sub(x, y int) int {
	return x - y
}`

	parser := Parser{}
	r := parser.Parse(str)

	if r == false {
		t.Errorf("Parser could not parse a valid file")
	}

	props := [3]FunctionProp{
		FunctionProp{"main", 0, "void"}, FunctionProp{"add", 2, "int"}, FunctionProp{"sub", 2, "int"},
	}
	
	for i := 0; i < 3; i++ {
		if _, ok := parser.file.definedFunctions[props[i].fname]; !ok {
			t.Fatalf("Parser could not parse function '%s'", props[i].fname)
		}

		if parser.file.definedFunctions[props[i].fname].name != props[i].fname {
			t.Errorf("Parser could not parse function '%s'", props[i].fname)
		}
		
		if len(parser.file.definedFunctions[props[i].fname].args) != props[i].fargs {
			t.Errorf("Parser could not get the args of '%s'", props[i].fname)
		}

		if parser.file.definedFunctions[props[i].fname].ret.name != props[i].fret {
			t.Errorf("Parser could not get the return type of '%s'", props[i].fname)
		}
	}
}
