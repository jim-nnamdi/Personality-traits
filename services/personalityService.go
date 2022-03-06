package services

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("forms/*"))

type Personality struct {
	Id        int
	Answer1   string
	Answer2   string
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

func ReturnAllPersonalityQuestions(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	results, err := db.Query("select * from questions")
	ErrorCheck(err)

	personalityQuestion := Personality{}
	personalityQuestions := []Personality{}

	for results.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline int

		err := results.Scan(&id, &answer1, &answer2, &scoreline)
		ErrorCheck(err)

		personalityQuestion.Id = id
		personalityQuestion.Answer1 = answer1
		personalityQuestion.Answer2 = answer2
		personalityQuestion.Scoreline = scoreline

		personalityQuestions = append(personalityQuestions, personalityQuestion)
	}
	tmpl.ExecuteTemplate(w, "returnAllPersonalityQuestions", nil)
	defer db.Close()
}

func ReturnSinglePersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	id := r.URL.Query().Get("id")
	result, err := db.Query("select * from questions where id = ?", id)
	ErrorCheck(err)

	singleQuestion := Personality{}
	for result.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline int

		err := result.Scan(&id, &answer1, &answer2, &scoreline)
		ErrorCheck(err)

		singleQuestion.Id = id
		singleQuestion.Answer1 = answer1
		singleQuestion.Answer2 = answer2
		singleQuestion.Scoreline = scoreline
	}
	tmpl.ExecuteTemplate(w, "returnAllPersonalityQuestions", nil)
	defer db.Close()
}

func EditPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	results, err := db.Query("select * from questions where id = ?", id)
	ErrorCheck(err)

	result := Personality{}
	for results.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline int

		err := results.Scan(&id, &answer1, &answer2, &scoreline)
		ErrorCheck(err)

		result.Id = id
		result.Answer1 = answer1
		result.Answer2 = answer2
		result.Scoreline = scoreline
	}
	tmpl.ExecuteTemplate(w, "edit", nil)
	defer db.Close()
}

func SaveAnswersToPersonalityTest(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	if r.Method == "POST" {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")

		i, err := strconv.Atoi(answer1)
		ErrorCheck(err)

		j, err := strconv.Atoi(answer2)
		ErrorCheck(err)
		scoreline := (i + j)

		if scoreline < 2 || scoreline < 0 {
			scorelineresult := "introvert"

			stmt, err := db.Prepare("insert into questions (answer1, answer2, scoreline) values(?, ?, ?)")
			ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult)
		} else {

			scorelineresult := "extrovert"

			stmt, err := db.Prepare("insert into questions (answer1, answer2, scoreline) values(?, ?, ?)")
			ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult)
		}

		log.Println("data submitted successfully")
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func UpdatePersonalityTraitData(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	if r.Method == "PUT" {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")

		i, err := strconv.Atoi(answer1)
		ErrorCheck(err)

		j, err := strconv.Atoi(answer2)
		ErrorCheck(err)

		scoreline := i + j

		if scoreline < 2 || scoreline < 0 {
			scorelineresult := "introvert"
			stmt, err := db.Prepare("update questions set answer1 = ?, answer2 =?, scoreline=? where id=?")
			ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult, id)
		} else {
			scorelineresult := "extrovert"
			stmt, err := db.Prepare("update questions set answer1 =?, answer2=?, scoreline=? where id =?")
			ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult, id)
		}
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

func DeletePersonalityTraitData(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("delete * from questions where id = ?")
	ErrorCheck(err)
	stmt.Exec(id)
	log.Println("resource trait deleted !")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
