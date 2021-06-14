# LetLang
A scanner/tokenizer, parser and evaluator for the Let programming language

Author: Taylor Privat
Interpreter for the Let Language

This is a scanner, parser and evaluator for the LET language written in go.

To Run:
*********************************************
type on the command line (in this directory)
`go run main.go let_scanner.go let_parser.go let_evaluator.go`

You will be propted to enter a file name
THIS FILE MUST BE IN THE CURRECT DIRECTORY WITH THE GO FILES

Output:
********************************************
The interpreter will print out the program read from the text file,
tokens extracted from text,
the AST created from the program,
and finally what the program evaluates to.


Let Language BNF Grammar:
```
1.Expression ::=  Number
2.           ::=  minus (Expression,Expression)
3.           ::=  iszero (Expression)
4.           ::=  if Expression then Expression else Expression
5.           ::=  Identifier
6.           ::=  let Identifier = Expression in Expression
```
