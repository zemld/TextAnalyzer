package handlers

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	doesHashesTableExist   = false
	doesAnalysisTableExist = false
)

const (
	pgStr           = "postgres://postgres:postgres@potgres_db:5432/postgres"
	createHashTable = "CREATE TABLE IF NOT EXISTS hashes " +
		"(id SERIAL PRIMARY KEY, hash TEXT)"
	createAnalysisTable = "CREATE TABLE IF NOT EXISTS analysis" +
		"(id INTEGER PRIMARY KEY, paragraphs_amount INTEGER, " +
		"sentences_amount INTEGER, words_amount INTEGER, " +
		"symbols_amount INTEGER, average_sentences_per_paragraph FLOAT, " +
		"average_words_per_sentence FLOAT, average_length_of_words FLOAT)"
	insertIntoHashTable     = "INSERT INTO hashes (hash) VALUES ($1)"
	insertIntoAnalysisTable = "INSERT INTO analysis " +
		"(id, paragraphs_amount, sentences_amount, words_amount, symbols_amount, " +
		"average_sentences_per_paragraph, average_words_per_sentence, average_length_of_words) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	selectForCheckThatFileExists = "SELECT id FROM hashes WHERE id = $1"
	selectForDuplicates          = "SELECT id FROM hashes WHERE hash = $1"
	selectFromAnalysisTable      = "SELECT * FROM analysis WHERE id = $1"
	hashesTableName              = "hashes"
	analysisTableName            = "analysis"
)

func checkFileExistance(id int) bool {
	if !doesHashesTableExist {
		err := createTable(hashesTableName)
		if err != nil {
			return false
		}
		doesHashesTableExist = true
		return false
	}

	db, err := sql.Open("pgx", pgStr)
	if err != nil {
		return false
	}
	defer db.Close()
	log.Println("Connected to db.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundId int
	err = db.QueryRowContext(ctx, selectForCheckThatFileExists, id).Scan(&foundId)
	if err != nil {
		return false
	}
	log.Println("Found id: ", foundId)
	return foundId == id
}

func createTable(tableName string) error {
	db, err := sql.Open("pgx", pgStr)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("Connected to db.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if tableName == hashesTableName {
		_, err = db.ExecContext(ctx, createHashTable)
	} else if tableName == analysisTableName {
		_, err = db.ExecContext(ctx, createAnalysisTable)
	}
	log.Println("Created table: ", tableName)
	if err != nil {
		return err
	}
	if tableName == hashesTableName {
		doesHashesTableExist = true
	} else if tableName == analysisTableName {
		doesAnalysisTableExist = true
	}
	return nil
}

func getAnalysisResult(id int) (map[string]any, error) {
	db, err := sql.Open("pgx", pgStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	log.Println("Connected to db.")
	if !doesAnalysisTableExist {
		err := createTable(analysisTableName)
		if err != nil {
			return nil, err
		}
		doesAnalysisTableExist = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var idFromTable, paragraphsAmount, sentencesAmount, wordsAmount, symbolsAmount int
	var avgSentencePerParagraph, avgWordsPerSentence, avgLengthOfWords float64

	err = db.QueryRowContext(ctx, selectFromAnalysisTable, id).Scan(
		&idFromTable,
		&paragraphsAmount,
		&sentencesAmount,
		&wordsAmount,
		&symbolsAmount,
		&avgSentencePerParagraph,
		&avgWordsPerSentence,
		&avgLengthOfWords)
	if err != nil {
		return nil, err
	}
	log.Println("Got analysis result for id: ", id)
	result := make(map[string]any)

	log.Println("Result: ", result)
	return result, nil
}

func storeAnalysisResult(analysis Analysis) error {
	db, err := sql.Open("pgx", pgStr)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("Connected to db.")
	if !doesAnalysisTableExist {
		err := createTable(analysisTableName)
		if err != nil {
			log.Println("Error creating analysis table: ", err)
			return err
		}
		doesAnalysisTableExist = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, insertIntoAnalysisTable,
		analysis.Id,
		analysis.ParagraphsAmount,
		analysis.SentencesAmount,
		analysis.WordsAmount,
		analysis.SymbolsAmount,
		analysis.AverageSentencesPerParagraph,
		analysis.AverageWordsPerSentence,
		analysis.AverageLengthOfWords)
	if err != nil {
		log.Printf("Error storing analysis result for id: %d, error: %v", analysis.Id, err)
		return err
	}
	log.Println("Stored analysis result for id: ", analysis.Id)
	return nil
}

func storeHash(hash string) int {
	db, err := sql.Open("pgx", pgStr)
	if err != nil {
		log.Println("Error connecting to db: ", err)
		return -1
	}
	defer db.Close()
	log.Println("Connected to db.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !doesHashesTableExist {
		err := createTable(hashesTableName)
		if err != nil {
			log.Println("Error creating hashes table: ", err)
			return -1
		}
		doesHashesTableExist = true
	}

	_, err = db.ExecContext(ctx, insertIntoHashTable, hash)
	if err != nil {
		log.Printf("Error storing hash, error: %v", err)
		return -1
	}
	var id int
	err = db.QueryRowContext(ctx, selectForDuplicates, hash).Scan(&id)
	if err != nil {
		log.Printf("Error getting id, error: %v", err)
		return -1
	}
	return id
}
