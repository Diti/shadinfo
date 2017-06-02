package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
)

var (
	tmplPath string
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"prettyDisk": func(i uint64) string {
			return humanize.Bytes(i)
		},
		"prettyMem": func(i uint64) string {
			return humanize.IBytes(i)
		},
	}).ParseFiles(tmplPath))

	info := PopulateInfo()
	if err := tmpl.ExecuteTemplate(w, filepath.Base(tmplPath), info); err != nil {
		logPanic(err)
	}
}

func logPanic(err error) {
	log.Panic(err.Error)
}

func handlePanic(w http.ResponseWriter, _ *http.Request, err interface{}) {
	errFmt := "Internal Server Error\n\n%s"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, errFmt, err)
	//fmt.Fprintf(w, "\n\n%s", debug.Stack())
}

func init() {
	flag.StringVar(&tmplPath, "template", "./index.html.tmpl", "Path to the Go Template file to use.")
}

func main() {
	flag.Parse()
	if _, err := os.Stat(tmplPath); err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s: Template file not found. Check your -template flag.", tmplPath)
		} else {
			log.Fatal(err)
		}
		flag.PrintDefaults()
		os.Exit(1)
	}

	router := httprouter.New()
	router.PanicHandler = handlePanic
	router.GET("/", index)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
