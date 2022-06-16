package delivery

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	ErrNotFound    = errors.New("page does not exist")
	ErrBadRequest  = errors.New("bad request")
	ErrEmptyInput  = errors.New("no input given")
	ErrWrongMethod = errors.New("method not allowed")
)

var MainTmpl, ErrTmpl, ArtistTmpl *template.Template

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ErrorHandler(page http.ResponseWriter, status int, err error) {
	page.WriteHeader(status)
	errorText := fmt.Sprintf("%d %s\n%v", status, http.StatusText(status), err)
	ErrTmpl.Execute(page, errorText)
}
