package tools

import (
	"errors"
	"strings"
)

// findIndentChar find character used for indentation and returns error if
// indentation is not consistent, meaning "space"[32] and "tab"[9] are used
// interchangeably throughout the string, empty lines will be ignored
// when evaluating only first character is being used
func findIndentChar(lines []string) (rune, error) {
	var indentChar rune
	for _, line := range lines {
		// skip evaluation on any line that is empty
		if len(line) == 0 {
			continue
		}

		fc := rune(line[0]) // first character
		if fc == 32 || fc == 9 {
			// if there was no indentation detected previously record character used
			// and use it for comparing to upcomming indent characters
			if indentChar == 0 {
				indentChar = fc
				continue

				// check if different indentation character has been used
			} else if fc != indentChar {
				return 0, errors.New("mixed indentation detected")
			}
		}
	}

	// check if indentation was detected at all
	if indentChar == 0 {
		return 0, errors.New("no indentation detected")
	}

	return indentChar, nil
}

// findGreatestCommonIndentLength finds greatest number of consecutive
// indentation character common for all lines, it will skip non-indented lines
// and count length only for those indented
func findGreatestCommonIndentLength(lines []string, indentChar rune) int {
	var greatestCommonIndent int
	for _, line := range lines {
		// find indentation length
		lineIndent := findLeftRepetitions(line, indentChar)

		// ignore lines with no indentation
		if lineIndent == 0 {
			continue
		}

		if lineIndent < greatestCommonIndent || greatestCommonIndent == 0 {
			greatestCommonIndent = lineIndent
		}
	}
	return greatestCommonIndent
}

// findLeftRepetitions finds number of character repeatitions in the string
// looking up from left to right
func findLeftRepetitions(s string, repeatChar rune) int {
	var count int
	for _, char := range []rune(s) {
		if char == repeatChar {
			count++
			continue
		}
		return count // return on first miss
	}
	return count
}

// DeIndent strips greatest common indent for each line, this will enable us
// to keep some degree of formatting after striping
//
// indenting characters are "space"[32] and "tab"[9]
// we would first find greatest common indent by counting successive indent
// characters in each line and store the minimal one
// if first line contains only indent characters or is empty it will be removed
// remove newline after join since strings.Join will add leading \n
func DeIndent(inString string) string {
	inLines := strings.Split(inString, "\n")
	var outLines []string

	// check if first line is empty or contains only indents
	if len(strings.TrimSpace(inLines[0])) == 0 {
		inLines = inLines[1:]
	}

	// check if last line is empty or contains only indents
	if len(strings.TrimSpace(inLines[len(inLines)-1])) == 0 {
		inLines = inLines[:len(inLines)-1]
	}

	// attempt to find character used for indentation, return on error
	indentChar, err := findIndentChar(inLines)
	if err != nil {
		return inString
	}

	// greatest common indent
	gci := findGreatestCommonIndentLength(inLines, indentChar)
	indent := strings.Repeat(string(indentChar), gci)

	// add lines to output and strip indenting prefix if found
	for _, line := range inLines {
		if strings.HasPrefix(line, indent) {
			outLines = append(outLines, line[gci:])
			continue
		}
		outLines = append(outLines, line)
	}

	// join strings and remove leading newline
	output := strings.Join(outLines, "\n")
	output = strings.TrimPrefix(output, "\n")

	return output
}
