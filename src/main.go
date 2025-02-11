package main

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", QrGenerator)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func QrGenerator(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data := r.FormValue("data")
	if data == "" {
		data = r.URL.Query().Get("data")
	}

	if data == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	s, err := url.QueryUnescape(data)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	code, err := qr.Encode(s, qr.L, qr.Auto)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	size := r.FormValue("size")
	if size == "" {
		size = r.URL.Query().Get("size")
	}
	if size == "" {
		size = "200"
	}
	intsize, err := strconv.Atoi(size)
	if err != nil {
		intsize = 200
	}

	// Scale the barcode to the appropriate size
	code, err = barcode.Scale(code, intsize, intsize)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, code); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
