
package main

//
// A definition point
//
// Records the line and column a certain object has been defined
//
type Definition struct {
	file string
	line uint32
	column uint32
}

//
// A type
//
type NamedType struct {
	def Definition
	name string
	methods map[string]Method
}

//
// A function
//
type Function struct {
	def Definition
	args []MethodArg
	ret NamedType
}

type MethodArg struct {
	name string
	ntype *NamedType
}

//
// A method
//
// Methods are functions bound to a type
type Method struct {
	boundType NamedType
	Function
}

// A module
type Module struct {
	def Definition
	name string
	types []NamedType
	functions []Function
}

// A file
type File struct {
	name string // filename relative to the project root
	definedMethods map[string]*Method
	definedFunctions map[string]*Function
	definedTypes map[string]*NamedType
}

// The project context
//
// We organize it like this:
// module -> types -> methods
//   |-> functins
//
// A module owns its types and function, and a type owns its methods
//
// We also have a pointer to the current file and to the current
// module
type Context struct {
	name string // An alias?
	root string // The project root

	currentFile *File
	currentModule *Module

	position Definition // the current cursor position
}
