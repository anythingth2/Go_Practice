// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"strings"
)

var ANSWER = [...]string{
	"Sure.",
	"Whoa, chill out!",
	"Calm down, I know what I'm doing!",
	"Fine. Be that way!",
	"Whatever."}

func IsContainAll(remark string, remark2 byte) bool {
	for i := range remark {
		if remark[i] != remark2 {
			return false
		}
	}
	return true
}
func IsOnlyNumber(text string) bool {
	for i := range text {
		if (text[i] > "A"[0] && text[i] < "Z"[0]) || (text[i] > "a"[0] && text[i] < "z"[0]) {
			return false
		}

	}
	return true
}

func GetEndingCharacter(text string) byte {
	if len(text) != 0 {
		ending := text[len(text)-1]
		if ending == " "[0] || ending == "\n"[0] || ending == "\r"[0] || ending == "\t"[0] {
			return GetEndingCharacter(text[:len(text)-1])
		} else {
			return ending
		}
	} else {
		return 0
	}
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	ans := ""
	fmt.Println(string(GetEndingCharacter(remark)))
	if len(remark) > 0 && GetEndingCharacter(remark) == "?"[0] {
		if strings.Compare(remark, strings.ToUpper(remark)) == 0 && !IsOnlyNumber(remark) {
			ans = ANSWER[2]
		} else {
			ans = ANSWER[0]
		}
	} else if strings.Compare(remark, strings.ToUpper(remark)) == 0 && !IsOnlyNumber(remark) {
		ans = ANSWER[1]
	} else if IsContainAll(remark, " "[0]) || IsContainAll(remark, "\t"[0]) || remark == "\n\r \t" {
		ans = ANSWER[3]
	} else {
		ans = ANSWER[4]
	}
	return ans
}
