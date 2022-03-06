package routes

import (
	"log"
	"net/http"

	"github.com/jim-nnamdi/Personality-traits.git/controller"
)

func Routes() {
	log.Println("server started running at port :9400")

	/** handle all the routes here */
	http.HandleFunc("/", controller.CreateNewPersonalityQuestion)
	http.HandleFunc("/start", controller.ReturnAllPersonalityQuestions)
	http.HandleFunc("/save", controller.SaveAnswersToPersonalityTest)
	http.HandleFunc("/edit", controller.EditPersonalityQuestion)
	http.HandleFunc("/update", controller.UpdatePersonalityTraitData)
	http.HandleFunc("/delete", controller.DeletePersonalityTraitData)

	/** access the public folder without the leading trail */
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	/** listen to the running Port here */
	http.ListenAndServe(":9400", nil)
}
