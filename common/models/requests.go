package models

type AnalysisRequest struct {
	// @required
	Id int `json:"id"`
	// @required
	ParagraphsAmount int `json:"paragraphs_amount"`
	// @required
	SentencesAmount int `json:"sentences_amount"`
	// @required
	WordsAmount int `json:"words_amount"`
	// @required
	SymbolsAmount int `json:"symbols_amount"`
	// @required
	AverageSentencesPerParagraph float64 `json:"average_sentences_per_paragraph"`
	// @required
	AverageWordsPerSentence float64 `json:"average_words_per_sentence"`
	// @required
	AverageLengthOfWords float64 `json:"average_length_of_words"`
}
