package victoria

// The parser context
type Parser struct {

	// Parsing data for the file being parsed
	// Will be merged with the context after parsing
	file ParsedFile
}

// Parse a text string with the file contents
func (p *Parser) Parse(data string) bool {

	return false
}

func (p *Parser) parseFile() {
	
}
