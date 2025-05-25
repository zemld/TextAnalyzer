package handlers

type FileStatusResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type Analysis struct {
	Id                           int     `json:"id"`
	ParagraphsAmount             int     `json:"paragraphs_amount"`
	SentencesAmount              int     `json:"sentences_amount"`
	WordsAmount                  int     `json:"words_amount"`
	SymbolsAmount                int     `json:"symbols_amount"`
	AverageSentencesPerParagraph float64 `json:"average_sentences_per_paragraph"`
	AverageWordsPerSentence      float64 `json:"average_words_per_sentence"`
	AverageLengthOfWords         float64 `json:"average_length_of_words"`
}

type Comparision struct {
	FirstId            int     `json:"first_id"`
	SecondId           int     `json:"second_id"`
	MatchingPercentage float64 `json:"matching_percentage"`
}
