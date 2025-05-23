package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// @description Analyze file.
// @tag.name File operations
// @param text formData string true "Text to analyze."
// @produce json
// @success 200 {object} Analysis
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analyze [get]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	log.Println(text)
	analysis := Analysis{}
	analysis.Id = -1
	analysis.ParagraphsAmount = countParagraphs(text)
	log.Println(analysis.ParagraphsAmount)
	analysis.SentencesAmount = countSentences(text)
	log.Println(analysis.SentencesAmount)
	analysis.WordsAmount = countWords(text)
	log.Println(analysis.WordsAmount)
	analysis.SymbolsAmount = countSymbols(text)
	log.Println(analysis.SymbolsAmount)
	analysis.AverageSentencesPerParagraph = float64(analysis.SentencesAmount) / float64(analysis.ParagraphsAmount)
	log.Println(analysis.AverageSentencesPerParagraph)
	analysis.AverageWordsPerSentence = float64(analysis.WordsAmount) / float64(analysis.SentencesAmount)
	log.Println(analysis.AverageWordsPerSentence)
	analysis.AverageLengthOfWords = float64(analysis.SymbolsAmount) / float64(analysis.WordsAmount)
	log.Println(analysis.AverageLengthOfWords)
	codedAnalysis, _ := json.Marshal(analysis)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(codedAnalysis)
}

// @description Get word cloud.
// @tag.name File operations
// @param text formData string true "Text to analyze."
// @produce png
// @success 200 {file} blob
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud [get]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {

}
