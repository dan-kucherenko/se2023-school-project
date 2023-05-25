package router

import (
	"log"
	"net/http"
	"os"
)

func ActivateRouter() {
	router := http.NewServeMux()

	router.HandleFunc("/rate", rate)
	router.HandleFunc("/subscribe", subscribe)
	router.HandleFunc("/sendEmails", sendEmails)

	err := http.ListenAndServe(os.Getenv("PATH"), router)
	log.Fatal(err)
}

func rate(rw http.ResponseWriter, r *http.Request)       {}
func subscribe(rw http.ResponseWriter, r *http.Request)  {}
func sendEmails(rw http.ResponseWriter, r *http.Request) {}
