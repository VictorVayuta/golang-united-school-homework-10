package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//NameParam returns "Hello, {PARAM}"
func NameParam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Fprintf(w, "Hello, "+params["PARAM"]+"!")
}

//BadParam returns status 500
func BadParam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

//BodyParam returns "I got message:\n {BODY}"
func BodyParam(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, errors.New("Error reading http response").Error())
		return
	}

	fmt.Fprintf(w, "I got message:\n"+string(body))
}

//HeadersParam returns header with key "a+b" and value {SUM}
func HeadersParam(w http.ResponseWriter, r *http.Request) {
	headerA := r.Header.Get("a")
	headerB := r.Header.Get("b")

	if headerA == "" || headerB == "" {
		fmt.Fprintf(w, errors.New("Empty header!").Error())
		return
	}

	valueA, errA := strconv.Atoi(headerA)
	valueB, errB := strconv.Atoi(headerB)

	if errA == nil && errB == nil {
		w.Header().Set("a+b", strconv.Itoa(valueA+valueB))
	}
}
