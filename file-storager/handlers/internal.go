package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	incorrectIdMsg = "Incorrect ID."
)

const (
	existsPattern    = "/files/exists/{id}"
	getPattern       = "/files/{id}"
	analysisPattern  = "/files/analysis/{id}"
	wordCloudPattern = "/files/wordcloud/{id}"
)

// TODO: имеет смысл добавить кастомный тип для id.

func parseIdFromRequestAndCreateResponse(w http.ResponseWriter, r *http.Request, pattern string) int {
	idStr, err := parseParamFromUrl(r.URL.Path, pattern, "{id}")
	if err != nil {
		log.Printf("Something went wrong while parsing id: %d.\n", -1)
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Something went wrong while parsing id: %d.\n", -1)
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	return id
}

func parseSaveFileRequestAndCreateResponse(w http.ResponseWriter, r *http.Request) (string, bool) {
	body := r.Body
	defer body.Close()

	textBytes, err := io.ReadAll(body)
	if err != nil {
		writeFileStatusResponse(w, -1, "Cannot read file.", http.StatusInternalServerError)
		return "", false
	}
	text := string(textBytes)
	return text, true
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
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	w.Write(repsJson)
}

func setAccessControlForOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func makeResponseFromSelectingAnalysisResult(id int, result map[string]any) Analysis {
	log.Printf("Making response from selecting analysis result: %v\n", result)
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

func getHash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
