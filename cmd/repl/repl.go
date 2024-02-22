package repl

import (
	"DGFC/pkg/lexer"
	"DGFC/pkg/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		scan_line := lexer.New(line)

		for tk := scan_line.NextToken(); tk.Type != token.EOF; tk = scan_line.NextToken() {
			fmt.Printf("%+v\n", tk)
		}

	}
}
