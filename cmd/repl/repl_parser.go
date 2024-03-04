package repl

import (
	"DGFC/pkg/lexer"
	"DGFC/pkg/parser"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const PROMPT = ">> "

const fileLocation = "/Users/drbngd/Documents/UC/14_2024_SPR/EECE6083/DGFC/reference_programs/correct/iterativeFib.src"

func Start(in io.Reader, out io.Writer) {
	fileContent, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Printf(string(fileContent))
	scanLine := lexer.New(string(fileContent))
	parsePointer := parser.New(scanLine)

	fmt.Printf("Current Token is: %+v\n", parsePointer.GetCurrentToken().Value)

	program := parsePointer.ParseProgram()

	if len(parsePointer.GetErrors()) != 0 {
		printParserErrors(out, parsePointer.GetErrors())
		return
	}

	io.WriteString(out, program.ToString())
	io.WriteString(out, "\n")
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
