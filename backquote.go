package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) != 2 {
		fmt.Println("Error: argument error.")
		return 1
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: open file error.")
		return 1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(bq(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}

var (
	//regDoubleQuoted = regexp.MustCompile(`".*?"`)
	regDoubleQuoted = regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
	regEscChars     = regexp.MustCompile(`\\(.)`)
)

func bq(s string) string {
	if strings.Contains(s, fmt.Sprintf("%c", '`')) {
		return s
	}

	doubleQuotedStr := regDoubleQuoted.FindAllString(s, -1)
	if len(doubleQuotedStr) == 0 {
		return s
	}
	doubleQuotedIdx := regDoubleQuoted.FindAllStringIndex(s, -1)
	if !strings.Contains(s[0:doubleQuotedIdx[0][0]], "script") {
		return s
	}

	srcString := doubleQuotedStr[0]
	// " -> `
	bq := '`'
	bqString := fmt.Sprintf("%c%s%c", bq, srcString[1:len(srcString)-1], bq)
	// \x -> x
	unescapedString := regEscChars.ReplaceAllString(bqString, "${1}")

	return fmt.Sprintf("%v", strings.Replace(s, srcString, unescapedString, 1))
}
