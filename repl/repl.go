package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

// prompt >>
const prompt = ">> "

// Start start REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			printIOError(err)
			_, err = io.WriteString(out, "\n")
			printIOError(err)
		}
	}
}

const monkeyFace = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	_, err := io.WriteString(out, monkeyFace)
	printIOError(err)
	_, err = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	printIOError(err)
	_, err = io.WriteString(out, " parser errors:\n")
	printIOError(err)
	for _, msg := range errors {
		_, err = io.WriteString(out, "\t"+msg+"\n")
		printIOError(err)
	}
}

func printIOError(err error) {
	if err != nil {
		fmt.Println("io.Writer error:" + err.Error())
	}
}
