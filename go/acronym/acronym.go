// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	// regex := regexp.MustCompile(`[A-Z]`)
	split_regex := regexp.MustCompile(`[\s-]`)
	// words := strings.Split(s, " ")
	words := split_regex.Split(s, 5)

	tla := ""[:]

	if true {
		for _, char := range words {
			tla += strings.ToUpper(string(char[0]))
			// if regex.MatchString(char) {

			// }else{

			// }
		}
		fmt.Println(tla)
		return tla
	} else {
		return ""
	}

}

func main() {
	test := "Portable Network Graphics"
	Abbreviate(test)
}
