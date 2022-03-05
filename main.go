package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Personality struct {
	Id        int
	Question  string
	Answer    string
	Scoreline int
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func DatabaseConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/personalitytest")
	ErrorCheck(err)
	return db
}

func returnAllPersonalityQuestions(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	results, err := db.Query("select * from questions")
	ErrorCheck(err)

	personalityQuestion := Personality{}
	personalityQuestions := []Personality{}

	for results.Next() {
		var id int
		var question string
		var answer string
		var scoreline int

		err := results.Scan(&id, &question, &answer, &scoreline)
		ErrorCheck(err)

		personalityQuestion.Id = id
		personalityQuestion.Question = question
		personalityQuestion.Answer = answer
		personalityQuestion.Scoreline = scoreline

		personalityQuestions = append(personalityQuestions, personalityQuestion)
	}
	json.NewEncoder(w).Encode(personalityQuestions)
	defer db.Close()
}

func main() {
	log.Println("server started running at port :9400")

	/** handle all the routes here */
	http.HandleFunc("/", returnAllPersonalityQuestions)

	/** listen to the running Port here */
	http.ListenAndServe(":9400", nil)
}
