package main

import "strings"

const (
	openBracket  = "("
	closeBracket = ")"
)

func findFirstStringInBracket(str string) string {
	if len(str) < 0 {
		return str
	}

	if strings.HasPrefix(str, openBracket) {
		str = strings.TrimPrefix(str, openBracket)
	}

	if strings.HasSuffix(str, closeBracket) {
		str = strings.TrimSuffix(str, closeBracket)
	}

	return str
}

/***
// Previous Code
	/*
		if len(str) > 0 {
			indexFirstBracketFound := strings.Index(str, "(")
			if indexFirstBracketFound >= 0 {
				runes := []rune(str)
				wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
				indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
				if indexClosingBracketFound >= 0 {
					runes := []rune(wordsAfterFirstBracket)
					return string(runes[1 : indexClosingBracketFound-1])
				} else {
					return ""
				}
			} else {
				return ""
			}
		} else {
			return ""
		}
		return ""
**/
