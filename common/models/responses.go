package models

type FileExistsResponse struct {
	Exists bool   `json:"exists"`
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type FileStatusResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type AnalysisResponse struct {
	Id                           int     `json:"id"`
	ParagraphsAmount             int     `json:"paragraphs_amount"`
	SentencesAmount              int     `json:"sentences_amount"`
	WordsAmount                  int     `json:"words_amount"`
	SymbolsAmount                int     `json:"symbols_amount"`
	AverageSentencesPerParagraph float64 `json:"average_sentences_per_paragraph"`
	AverageWordsPerSentence      float64 `json:"average_words_per_sentence"`
	AverageLengthOfWords         float64 `json:"average_length_of_words"`
}

type CompareResponse struct {
	FirstId                          int     `json:"first_id"`
	SecondId                         int     `json:"second_id"`
	ParagraphsAmountDiff             int     `json:"paragraphs_amount_diff"`
	SentencesAmountDiff              int     `json:"sentences_amount_diff"`
	WordsAmountDiff                  int     `json:"words_amount_diff"`
	SymbolsAmountDiff                int     `json:"symbols_amount_diff"`
	AverageSentencesPerParagraphDiff float64 `json:"average_sentences_per_paragraph_diff"`
	AverageWordsPerSentenceDiff      float64 `json:"average_words_per_sentence_diff"`
	AverageLengthOfWordsDiff         float64 `json:"average_length_of_words_diff"`
}
