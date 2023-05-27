package router

import (
	"fmt"
	"github.com/dan-kucherenko/se-school-project/currency_rate_getter"
	"github.com/dan-kucherenko/se-school-project/emails"
	"log"
	"net/http"
	"os"
)

// main routing function to start the server and activate the endpoint handlers
func ActivateRouter() {
	router := http.NewServeMux()

	router.HandleFunc("/rate", rate)
	router.HandleFunc("/subscribe", subscribe)
	router.HandleFunc("/sendEmails", sendEmails)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}

// rate handler function
func rate(rw http.ResponseWriter, _ *http.Request) {
	time, curRate, err := currency_rate_getter.GetRateBtcToUah()
	if err != nil {
		rw.WriteHeader(400)
		http.Error(rw, "Error during getting the currency rate", http.StatusBadRequest)
		return
	}
	resultingString := fmt.Sprintf("Current rate at %s BTC to UAH is %f", time, curRate)
	rw.WriteHeader(200)
	_, _ = fmt.Fprint(rw, resultingString)
}

// subscribe handler function
func subscribe(rw http.ResponseWriter, r *http.Request) {
	candidateEmail := r.FormValue("email")
	if emails.IsEmailSubscribed(candidateEmail) {
		rw.WriteHeader(409)
		_, _ = fmt.Fprintf(rw, "Error! Email %s is already subscribed", candidateEmail)
	} else {
		err := emails.SubscribeNewEmail(candidateEmail)
		if err != nil {
			_, _ = fmt.Fprint(rw, "Error subscribing new email")
			return
		}
		rw.WriteHeader(200)
		_, _ = fmt.Fprintf(rw, "Email %s is successfully subscribed", candidateEmail)
	}
}

// sendEmails handler function
func sendEmails(rw http.ResponseWriter, _ *http.Request) {
	err := emails.SendEmailWithRate()
	if err != nil {
		_, _ = fmt.Fprint(rw, "Error sending the emails with rate")
	}
	rw.WriteHeader(200)
	_, _ = fmt.Fprint(rw, "Emails were successfully sent")
}
