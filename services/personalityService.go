package services

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jim-nnamdi/Personality-traits.git/models"
)

var tmpl = template.Must(template.ParseGlob("forms/*"))

func DatabaseConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/personalitytest")
	models.ErrorCheck(err)
	return db
}

/**  we can return all instances of the test taken, but i'll return the latest one*/
func ReturnAllPersonalityQuestions(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	results, err := db.Query("select * from questions order by id desc limit 1")
	models.ErrorCheck(err)

	personalityQuestion := models.Personality{}
	personalityQuestions := []models.Personality{}

	for results.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline string

		err := results.Scan(&id, &answer1, &answer2, &scoreline)
		models.ErrorCheck(err)

		personalityQuestion.Id = id
		personalityQuestion.Answer1 = answer1
		personalityQuestion.Answer2 = answer2
		personalityQuestion.Scoreline = scoreline

		personalityQuestions = append(personalityQuestions, personalityQuestion)
	}
	tmpl.ExecuteTemplate(w, "result", personalityQuestions)
	defer db.Close()
}

func ReturnSinglePersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()

	id := r.URL.Query().Get("id")
	result, err := db.Query("select * from questions where id = ?", id)
	models.ErrorCheck(err)

	singleQuestion := models.Personality{}
	for result.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline string

		err := result.Scan(&id, &answer1, &answer2, &scoreline)
		models.ErrorCheck(err)

		singleQuestion.Id = id
		singleQuestion.Answer1 = answer1
		singleQuestion.Answer2 = answer2
		singleQuestion.Scoreline = scoreline
	}
	tmpl.ExecuteTemplate(w, "show", singleQuestion)
	defer db.Close()
}

func EditPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	results, err := db.Query("select * from questions where id = ?", id)
	models.ErrorCheck(err)

	result := models.Personality{}
	for results.Next() {
		var id int
		var answer1 string
		var answer2 string
		var scoreline string

		err := results.Scan(&id, &answer1, &answer2, &scoreline)
		models.ErrorCheck(err)

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

		if answer1 == "true" && answer2 == "false" || answer1 == "false" && answer2 == "true" {
			scorelineresult := "introvert"

			stmt, err := db.Prepare("insert into questions (answer1, answer2, scoreline) values(?, ?, ?)")
			models.ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult)
		} else {

			scorelineresult := "extrovert"

			stmt, err := db.Prepare("insert into questions (answer1, answer2, scoreline) values(?, ?, ?)")
			models.ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult)
		}
		log.Println("data submitted successfully")
	}
	defer db.Close()
	http.Redirect(w, r, "result", 301)
}

func UpdatePersonalityTraitData(w http.ResponseWriter, r *http.Request) {
	db := DatabaseConnection()
	id := r.URL.Query().Get("id")

	if r.Method == "PUT" {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")

		// instead of this i should have done a proper validation
		if answer1 == "" || answer2 == "" {
			log.Fatal(" answers not filled! ")
		}

		if answer1 == "true" && answer2 == "false" || answer1 == "false" && answer2 == "true" {
			scorelineresult := "introvert"
			stmt, err := db.Prepare("update questions set answer1 = ?, answer2 =?, scoreline=? where id=?")
			models.ErrorCheck(err)
			stmt.Exec(answer1, answer2, scorelineresult, id)
		} else {
			scorelineresult := "extrovert"
			stmt, err := db.Prepare("update questions set answer1 =?, answer2=?, scoreline=? where id =?")
			models.ErrorCheck(err)
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
	models.ErrorCheck(err)
	stmt.Exec(id)
	log.Println("resource trait deleted !")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func LandingScreen(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index", nil)
}
