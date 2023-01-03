package main

// https://github.com/shaytan4/mygo/blob/ac21d7da5286d9050ccef5047760754ab09c7aec/01_Mastering_go/ch06_functions/html_tmpl/main.go
// https://www.sohamkamani.com/golang/exec-shell-command/
// https://zetcode.com/golang/exec-command/

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html

// https://habr.com/ru/post/122095/

// https://stackoverflow.com/questions/15240884/how-can-i-handle-http-requests-of-different-methods-to-in-go

import (
	"html/template"
	"net/http"

	srvctl "srvctl/pkg"
)

func init() {
	srvctl.Tpl = template.Must(template.ParseGlob("html_templates/*.gohtml"))
}

func main() {
	cfgData := srvctl.LoadCfg()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		srvctl.IndexHandler(w, r, cfgData)
	})

	http.ListenAndServe(":8080", nil)

}
