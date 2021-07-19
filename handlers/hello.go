package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with a given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("HELLO FROM HANDLER")
	d, e := ioutil.ReadAll(r.Body)
	if e != nil {
		// log.Println("Could not read body")
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oppp!!!"))
		http.Error(rw, "Oppps", http.StatusBadRequest)
		return
	}
	log.Printf(" Data sent is %s \n", d)
	// write the response
	fmt.Fprintf(rw, " Hello %s \n", d)
	log.Println("Hello from Micro Service!!")
}
