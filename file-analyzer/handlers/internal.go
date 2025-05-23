package handlers

import (
	"strings"
)

func countSymbols(text string) int {
	words := strings.Split(text, " ")
	var symbolsCnt int
	for _, word := range words {
		symbolsCnt += len(word)
	}
	return symbolsCnt
}

func countWords(text string) int {
	return len(strings.Split(text, " "))
}

func countSentences(text string) int {
	separatedByDot := strings.Split(text, ".")
	var separatedByDotAndQuestionMark []string
	for _, sentence := range separatedByDot {
		separatedByDotAndQuestionMark = append(separatedByDotAndQuestionMark, strings.Split(sentence, "?")...)
	}
	var separatedByDotQuestionMarkAndExclamationMark []string
	for _, sentence := range separatedByDotAndQuestionMark {
		separatedByDotQuestionMarkAndExclamationMark = append(separatedByDotQuestionMarkAndExclamationMark, strings.Split(sentence, "!")...)
	}
	return len(separatedByDotQuestionMarkAndExclamationMark)
}

func countParagraphs(text string) int {
	return len(strings.Split(text, "\n"))
}
