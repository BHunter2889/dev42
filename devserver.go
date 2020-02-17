// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	wd, _ := os.Getwd()
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	// Set Custom Template / Title
	simplehttp.DefaultPageTemplateSource = template_string
	simplehttp.DefaultStaticData["Title"] = title

	// import CSS / fonts
	simplehttp.DefaultStaticData["CSSFiles"] = []string{cssIndex, cssLoader, cssSocial}

	// create handler and serve
	h := simplehttp.New(wd, true)
	log.Fatal(http.ListenAndServe(l, h))
}
