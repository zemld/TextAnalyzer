package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// @description Analyze file.
// @tag.name File operations
// @param text body string true "Text to analyze."
// @produce json
// @success 200 {object} Analysis
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analyze [post]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	textBytes, err := io.ReadAll(body)
	if err != nil {
		log.Println(err)
		return
	}
	text := string(textBytes)
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
// @param text formData string true "Text to analyze."
// @produce png
// @success 200 {file} blob
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud [post]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {

}
