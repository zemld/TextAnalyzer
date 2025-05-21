package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

const (
	incorrectIdMsg = "Incorrect ID."
)

// TODO: имеет смысл добавить кастомный тип для id.

func parseIdFromRequestAndCreateResponse(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	return id
}

func writeFileExistsResponse(w http.ResponseWriter, resp FileExistsResponse) {
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(resp)
	w.Write(repsJson)
}

func writeBadFileExistsResponse(w http.ResponseWriter) {
	resp := FileExistsResponse{Exists: false, Id: -1, Status: "File doesn't exist"}
	writeFileExistsResponse(w, resp)
}

func writeGoodFileExistsResponse(w http.ResponseWriter, id int) {
	resp := FileExistsResponse{Exists: true, Id: id, Status: "File exists"}
	writeFileExistsResponse(w, resp)
}

func writeFileStatusResponse(w http.ResponseWriter, id int, msg string, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	w.Write(repsJson)
	w.WriteHeader(statusCode)
}

func setAccessControlForOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func makeResponseFromSelectingAnalysisResult(id int, result map[string]any) Analysis {
	return Analysis{
		Id:                           id,
		ParagraphsAmount:             result["paragraphs_amount"].(int),
		SentencesAmount:              result["sentences_amount"].(int),
		WordsAmount:                  result["words_amount"].(int),
		SymbolsAmount:                result["symbols_amount"].(int),
		AverageSentencesPerParagraph: result["average_sentences_per_paragraph"].(float64),
		AverageWordsPerSentence:      result["average_words_per_sentence"].(float64),
		AverageLengthOfWords:         result["average_length_of_words"].(float64),
	}
}

func writeAnalysisResponse(w http.ResponseWriter, analysis Analysis) {
	w.Header().Add("Content-Type", "application/json")
	encodedAnalysis, _ := json.Marshal(analysis)
	w.Write(encodedAnalysis)
	w.WriteHeader(http.StatusOK)
}
