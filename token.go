package interpreter

type Token int

const (
	//Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	//Literals
	IDENTIFIER

	//Misc characters
	COMMA // ,
	POINT // .
	SP // (
	EP // )
)

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

var eof = rune(0)

