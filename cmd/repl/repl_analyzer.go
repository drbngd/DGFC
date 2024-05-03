package repl

import (
	"DGFC/pkg/analyzer"
	"DGFC/pkg/lexer"
	"DGFC/pkg/parser"
	"fmt"
	"io"
	"os"
)

const PROMPT = ">> "

const fileLocation = "test/correct/test1b.src"

func Start(in io.Reader, out io.Writer) {
	fileContent, err := os.ReadFile(fileLocation)
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

	// analyzing the program
	fmt.Println(out, "\n\n++++++++STARTING ANALYSIS++++++++\n\n")
	anlyz := analyzer.New()
	_, err = anlyz.SymbolTable.Analyze(program, "GLOBAL")
	if err != nil {
		fmt.Println(out, "Error analyzing program: %s\n", err)
		return
	}

	io.WriteString(out, program.ToString())
	io.WriteString(out, "\n")
	anlyz.SymbolTable.PrintSymbolTable()
	io.WriteString(out, "\n")
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
