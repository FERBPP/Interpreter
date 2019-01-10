package interpreter 

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	
}