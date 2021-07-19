package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye from Us")

	d, e := ioutil.ReadAll(r.Body)
	if e != nil {
		http.Error(w, "Oppps", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Goodbye %s \n", d)
}
