package repl

import (
	"bufio"
	"fmt"
	"gotoxavier2333.com/monkey/lexer"
	"gotoxavier2333.com/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprintf(out, PROMPT)
		if err != nil {
			return
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				return
			}
		}
	}
}
