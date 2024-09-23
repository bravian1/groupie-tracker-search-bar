package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	ErrMsg string
	Code   int
}

func errorMsg(w http.ResponseWriter, errMsg string, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
		errMsg = "Internal Server Error"
	}
	tmpl, err := template.ParseFiles("frontend/errors.html")
	if err != nil {
		log.Printf("error parsing template files: %v", err)
		http.Error(w, "Oops! Cannot load page.\nTry again later\n", http.StatusInternalServerError)
		return
	}

	data := &Data{
		ErrMsg: errMsg,
		Code:   statusCode,
	}

	tmpl.Execute(w, data)
}

func Static(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	info, err := os.Stat("." + url)
	if err != nil {
		errorMsg(w, "File not found", http.StatusNotFound)
		return
	}

	fmt.Println("Is a directory?", info.IsDir())
	if info.IsDir() {
		errorMsg(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, "."+url)
}
