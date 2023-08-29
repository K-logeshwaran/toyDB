// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	var (
// 		name    string
// 		age     int
// 		isAdmin bool
// 	)

// 	flag.StringVar(&name, "name", "John Doe", "The name of the user")
// 	flag.IntVar(&age, "age", 25, "The age of the user")
// 	flag.BoolVar(&isAdmin, "admin", false, "Whether the user is an admin")

// 	flag.Parse()

// 	// Your code here using the flag values
// 	fmt.Printf("Name: %s\n", name)
// 	fmt.Printf("Age: %d\n", age)
// 	fmt.Printf("Admin: %v\n", isAdmin)
// }

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TokenType int

const (
	NUMBER TokenType = 0
	PLUS   TokenType = 1
	MINUS  TokenType = 2
)

type Token struct {
	Type  TokenType
	Value string
}

func lex(input string) []Token {
	var tokens []Token
	parts := strings.Split(input, " ")
	fmt.Println(len(parts))
	for _, part := range parts {
		if part == "+" {
			tokens = append(tokens, Token{Type: PLUS, Value: part})
		} else if part == "-" {
			tokens = append(tokens, Token{Type: MINUS, Value: part})
		} else {
			tokens = append(tokens, Token{Type: NUMBER, Value: part})
		}
	}
	fmt.Println("Token len", len(tokens))
	return tokens
}

func parse(tokens []Token) int {
	pos := 0
	currentToken := tokens[pos]
	pos++
	result := parseTerm(tokens, &pos)
	for currentToken.Type == PLUS || currentToken.Type == MINUS {
		if currentToken.Type == PLUS {
			currentToken = tokens[pos]
			pos++
			result += parseTerm(tokens, &pos)
		} else if currentToken.Type == MINUS {
			currentToken = tokens[pos]
			pos++
			result -= parseTerm(tokens, &pos)
		}
	}
	return result
}

func parseTerm(tokens []Token, pos *int) int {
	currentToken := tokens[*pos]
	*pos++
	num, _ := strconv.Atoi(currentToken.Value)
	return num
}

func main() {
	input := "5 + 3 - 3"
	tokens := lex(input)
	result := parse(tokens)
	fmt.Println("Result:", result)
	for _, v := range tokens {
		fmt.Println("Result:", v)
	}
}
