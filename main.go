package main

// https://github.com/shaytan4/mygo/blob/ac21d7da5286d9050ccef5047760754ab09c7aec/01_Mastering_go/ch06_functions/html_tmpl/main.go

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	srvctl "srvctl/pkg"
)

var tpl *template.Template

//var DATA map[string]string

func init() {
	tpl = template.Must(template.ParseGlob("html_templates/*.gohtml"))
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request, mydata map[string]string) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", mydata)
	HandleError(w, err)
	//fmt.Println(mydata)
}

func main() {
	cfgData := srvctl.LoadCfg()

	for k, v := range cfgData {
		fmt.Println(k, v)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, cfgData)
	})

	http.ListenAndServe(":8080", nil)

	// err := tpl.Execute(os.Stdout, cfgData)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
