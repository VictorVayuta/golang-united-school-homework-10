package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
)

func NameParam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Fprintf(w, "Hello, "+params["PARAM"]+"!")
}

func BadParam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func BodyParam(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, errors.New("Error reading http response").Error())
		return
	}

	fmt.Fprintf(w, "I got message:\n"+string(body))
}

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
