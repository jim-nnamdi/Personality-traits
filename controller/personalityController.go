package controller

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/jim-nnamdi/Personality-traits.git/services"
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
	services.ReturnAllPersonalityQuestions(w, r)
}

func ReturnSinglePersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	services.ReturnSinglePersonalityQuestion(w, r)
}

func CreateNewPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "create", nil)
}

func EditPersonalityQuestion(w http.ResponseWriter, r *http.Request) {
	services.EditPersonalityQuestion(w, r)
}

func SaveAnswersToPersonalityTest(w http.ResponseWriter, r *http.Request) {
	services.SaveAnswersToPersonalityTest(w, r)
}

func UpdatePersonalityTraitData(w http.ResponseWriter, r *http.Request) {
	services.UpdatePersonalityTraitData(w, r)
}

func DeletePersonalityTraitData(w http.ResponseWriter, r *http.Request) {
	services.DeletePersonalityTraitData(w, r)
}
