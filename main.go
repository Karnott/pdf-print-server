package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	http.HandleFunc("/pdf", handlePDF)
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Launch
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handlePDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, fmt.Sprintf("Expect POST http method, got: %v", r.Method))
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendError(w, err.Error())
		return
	}
	html := bytes.NewBuffer(data)
	pdf, err := generatePdf(html)
	if err != nil {
		sendError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	_, err = w.Write(pdf.Bytes())
	if err != nil {
		sendError(w, err.Error())
		return
	}
}

//
//
//
func sendError(w http.ResponseWriter, errorMessage string) {
	log.Printf(" Error : %v", errorMessage)
	w.WriteHeader(400)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(map[string]string{
		"error": errorMessage,
	})
}

//
//
//
func generatePdf(html *bytes.Buffer) (*bytes.Buffer, error) {
	p, err := pdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	p.Dpi.Set(600)
	page := pdf.NewPageReader(html)
	p.AddPage(page)

	err = p.Create()
	if err != nil {
		return nil, err
	}

	return p.Buffer(), nil
}
