// a code to generate a list made from words in their normal and "l33t" writing to be used
// as blacklist for passwords in a domain environment
// by @x_Freed0m
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var replacer = strings.NewReplacer("a", "@", "b", "6", "e", "3", "f", "ph", "g", "9", "i", "1", "o", "0", "q", "9", "s", "5", "t", "7", "0", "o","1", "l")

// readLines reads a whole file into memory
// and returns a slice of its lines.

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening input file: %s\n", path)
	}
	defer Close(file)


	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func charConvert(lines []string) []string {
	var leetLines []string

	for i, leetLine := range lines {
		leetLine = strings.ToLower(leetLine)
		leetLines = append(leetLines, leetLine)
		fmt.Printf("Saving original word #%d\n", i)
		leetLine = replacer.Replace(leetLine)
		leetLines = append(leetLines, leetLine)
		fmt.Printf("Saving manipulated word #%d\n", i)
	}
	return leetLines

}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer Close(file)

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
	}

	return writer.Flush()
}

func banner() {
	println(`
	##        #######   #######  ######## ######## ########  
	##       ##     ## ##     ##    ##    ##       ##     ## 
	##              ##        ##    ##    ##       ##     ## 
	##        #######   #######     ##    ######   ########  
	##              ##        ##    ##    ##       ##   ##   
	##       ##     ## ##     ##    ##    ##       ##    ##  
	########  #######   #######     ##    ######## ##     ##`,
	"\nby @x_Freed0m\n\n")

}

func main() {
	banner()
	inputFile := flag.String("input", "input.txt", "Which input file to process")
	outputFile := flag.String("output", "output.txt", "Which output file to save the manipulated words")
	flag.Parse()

	lines, err := readLines(*inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Printf("Successfully read file %s\n\n", *inputFile)

	if err := writeLines(lines, *outputFile); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

	leetLines := charConvert(lines)
	fmt.Printf("\nSuccessfully manipulated all the words from: %s\n", *inputFile)

	if err := writeLines(leetLines, *outputFile); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	fmt.Printf("Successfully wrote all the words to file: %s\n", *outputFile)

}

//TODO:
// provide more options based on chars not lines
