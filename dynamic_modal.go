package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type ModalModel struct {
	ModalType string
}

func generateModal(w http.ResponseWriter, r *http.Request) {

	log.Println("Entered generateModal")

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	modalLottery := random.Intn(3)

	log.Println("Random Number: ", modalLottery)

	modal := ModalModel{}

	// Choose a random modal type to send to the template
	switch modalLottery {

	case 0:
		modal.ModalType = "A"
	case 1:
		modal.ModalType = "B"
	case 2:
		modal.ModalType = "C"
	default:
		// This should not be possible
		log.Panic()
	}

	log.Println(modal.ModalType)

	t, err := template.ParseFiles("views/modal.html")

	if err != nil {
		panic(err)
	}

	t.Execute(w, &modal)

}

func main() {

	// Add the routes

	http.Handle("/", http.FileServer(http.Dir("./web")))

	http.HandleFunc("/modal", generateModal)

	fmt.Println("Serving on Port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
