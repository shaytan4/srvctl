package main

import (
	"log"
	"os"
	srvctl "srvctl/pkg"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("html/tpl.gohtml"))
}

func main() {
	cfgData := srvctl.LoadCfg()
	// for k, v := range cfgData {
	// 	fmt.Println(k, v)
	// }

	err := tpl.Execute(os.Stdout, cfgData)
	if err != nil {
		log.Fatalln(err)
	}

}
