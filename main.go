package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("forms/*"))

type Personality struct {
	Id        int
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
		var answer string
		var scoreline int

		err := results.Scan(&id, &answer, &scoreline)
		ErrorCheck(err)

		personalityQuestion.Id = id
		personalityQuestion.Answer = answer
		personalityQuestion.Scoreline = scoreline

		personalityQuestions = append(personalityQuestions, personalityQuestion)
	}
	tmpl.ExecuteTemplate(w, "returnAllPersonalityQuestions", nil)
	defer db.Close()
}

func returnSinglePersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	id := r.URL.Query().Get("id")
	result, err := db.Query("select * from questions where id = ?", id)
	ErrorCheck(err)

	singleQuestion := Personality{}
	for result.Next() {
		var id int
		var answer string
		var scoreline int

		err := result.Scan(&id, &answer, &scoreline)
		ErrorCheck(err)

		singleQuestion.Id = id
		singleQuestion.Answer = answer
		singleQuestion.Scoreline = scoreline
	}
	tmpl.ExecuteTemplate(w, "returnSinglePersonalityQuestion", nil)
	defer db.Close()
}

func createNewPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "create", nil)
}

func editPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	results, err := db.Query("select * from questions where id = ?", id)
	ErrorCheck(err)

	result := Personality{}
	for results.Next() {
		var id int
		var answer string
		var scoreline int

		err := results.Scan(&id, &answer, &scoreline)
		ErrorCheck(err)

		result.Id = id
		result.Answer = answer
		result.Scoreline = scoreline
	}
	tmpl.ExecuteTemplate(w, "edit", nil)
	defer db.Close()
}

func saveAnswersToPersonalityTest(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	if r.Method == "POST" {
		answer
	}
}

func main() {
	log.Println("server started running at port :9400")

	/** handle all the routes here */
	http.HandleFunc("/", returnAllPersonalityQuestions)

	/** listen to the running Port here */
	http.ListenAndServe(":9400", nil)
}
