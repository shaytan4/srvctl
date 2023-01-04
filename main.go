package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	srvctl "srvctl/pkg"
)

func init() {
	srvctl.Tpl = template.Must(template.ParseGlob("html_templates/*.gohtml"))
}

func main() {
	var port string
	cfgData := srvctl.LoadCfg()

	customPort := os.Getenv("PORT")
	if customPort == "" {
		port = ":8080"
	} else {
		port = ":" + customPort
	}
	log.Println("used port ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		srvctl.IndexHandler(w, r, cfgData)
	})

	http.ListenAndServe(port, nil)

}
