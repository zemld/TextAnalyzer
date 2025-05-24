package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	wordCloudUrl string = "https://quickchart.io/wordcloud"
)

// @description Analyze file.
// @tag.name File operations
// @param text body string true "Text to analyze."
// @produce json
// @success 200 {object} Analysis
// @failure 500 {string} string
// @router /files/analyze [post]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	text, ok := parseTextFromRequest(r)
	if !ok {
		writeResponse(w, "Cannot read file.", http.StatusInternalServerError)
		return
	}
	log.Println(text)
	analysis := Analysis{}
	analysis.Id = -1
	analysis.ParagraphsAmount = countParagraphs(text)
	analysis.SentencesAmount = countSentences(text)
	analysis.WordsAmount = countWords(text)
	analysis.SymbolsAmount = countSymbols(text)
	if len(text) == 0 {
		analysis.AverageSentencesPerParagraph = 0
		analysis.AverageWordsPerSentence = 0
		analysis.AverageLengthOfWords = 0
	} else {
		analysis.AverageSentencesPerParagraph = float64(analysis.SentencesAmount) / float64(analysis.ParagraphsAmount)
		analysis.AverageWordsPerSentence = float64(analysis.WordsAmount) / float64(analysis.SentencesAmount)
		analysis.AverageLengthOfWords = float64(analysis.SymbolsAmount) / float64(analysis.WordsAmount)
	}
	codedAnalysis, _ := json.Marshal(analysis)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(codedAnalysis)
}

// @description Get word cloud.
// @tag.name File operations
// @param text body string true "Text to analyze."
// @produce png
// @success 200 {file} blob
// @failure 500 {string} string
// @router /files/wordcloud [post]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	text, ok := parseTextFromRequest(r)
	if !ok {
		writeResponse(w, "Cannot read file.", http.StatusInternalServerError)
		return
	}
	log.Println(text)
	body := createRequestForWordCloud(text)
	resp, err := http.Post(wordCloudUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		writeResponse(w, "Cannot get word cloud.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		writeResponse(w, "Cannot get word cloud.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "inline; filename=\"wordcloud.png\"")
	io.Copy(w, resp.Body)
}
