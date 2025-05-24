package handlers

import (
	"encoding/json"
	"io"
	"net/http"
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

func writeResponse(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(msg))
}

func parseTextFromRequest(r *http.Request) (string, bool) {
	body := r.Body
	defer body.Close()

	textBytes, err := io.ReadAll(body)
	if err != nil {
		return "", false
	}
	text := string(textBytes)
	return text, true
}

func createRequestForWordCloud(text string) []byte {
	requestData := map[string]interface{}{
		"text":       text,
		"format":     "png",
		"width":      800,
		"height":     400,
		"fontFamily": "sans-serif",
		"fontScale":  15,
		"scale":      "linear",
		"padding":    5,
		"colors":     []string{"#1f77b4", "#ff7f0e", "#2ca02c", "#d62728", "#9467bd"},
	}
	jsonData, _ := json.Marshal(requestData)
	return jsonData
}
