// +build !wasm

package main

//go:generate vugugen .

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("Provided Port: %q", port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	port = ":" + port // ":" is used as the required Heroku format

	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	httpl := flag.String("http", port, "Listen for HTTP on this host:port")
	flag.Parse()
	wd, _ := filepath.Abs(*dir)
	os.Chdir(wd)

	// Set Custom Template / Title
	simplehttp.DefaultPageTemplateSource = template_string
	simplehttp.DefaultStaticData["Title"] = title

	// import CSS / fonts
	simplehttp.DefaultStaticData["CSSFiles"] = []string{cssIndex, cssLoader, cssSocial}

	log.Printf("Starting HTTP Server at %q", *httpl)
	h := simplehttp.New(wd, *dev)
	log.Fatal(http.ListenAndServe(*httpl, h))
}
