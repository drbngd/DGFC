package repl

import (
	"DGFC/pkg/lexer"
	"DGFC/pkg/parser"
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
		scanLine := lexer.New(line)
		parsePointer := parser.New(scanLine)

		program := parsePointer.ParseProgram()

		if len(parsePointer.GetErrors()) != 0 {
			printParserErrors(out, parsePointer.GetErrors())
			continue
		}

		io.WriteString(out, program.ToString())
		io.WriteString(out, "\n")

		// for lexer testing
		//for tk := scanLine.NextToken(); tk.Type != token.EOF; tk = scanLine.NextToken() {
		//	fmt.Printf("%+v\n", tk)
		//}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
