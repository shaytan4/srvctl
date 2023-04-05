// https://stackoverflow.com/questions/37492544/golang-http-redirect-strange-behavior

package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

var Tpl *template.Template

var adm_name = "admin"
var adm_pass = "parol"

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request, mydata map[string]string) {
	url_path := r.URL.Path[1:]

	err := Tpl.ExecuteTemplate(w, "index.gohtml", mydata)
	HandleError(w, err)

	if url_path != "" {
		log.Println("URL path values is", url_path)
		execCmd := mydata[url_path]
		log.Println("execCmd command got form  map[string]string -------------- ", execCmd)

		s := strings.Split(execCmd, " ")

		//runCmd := exec.Command(execCmd)
		runCmd := exec.Command(s[0], s[1:]...)
		//runCmd.Env = os.Environ()
		out, err := runCmd.Output()
		if err != nil {
			log.Println("could not run command: ", err)
		} else {
			log.Println("Output: ", string(out))
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	err := Tpl.ExecuteTemplate(w, "login.gohtml", adm_name)
	HandleError(w, err)

	// process form submission
	if r.Method == http.MethodPost {
		login := r.FormValue("username")
		pass := r.FormValue("password")

		if login == adm_name && pass == adm_pass {
			log.Println("Login is OK")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			//http.RedirectHandler("/", http.StatusSeeOther)
			//return
		} else {
			log.Printf("Login Err , Login - %s ,Pass - %s", login, pass)
			http.Redirect(w, r, "/errPage", http.StatusFound)
		}
	}

}

func LoadCfg() map[string]string {
	var configFile *string = flag.String("c", "myConfig", "Setting the configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)

	if err == nil {
		log.Println("Using User Specified Configuration file!")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/etc")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		//return
	}
	log.Printf("Using config: %s\n", viper.ConfigFileUsed())

	confLines := make(map[string]string)

	for _, i := range viper.AllKeys() {
		log.Println("key: ", i, "value: ", viper.Get(i))
		x := viper.Get(i)
		n := fmt.Sprintf("%s", x)
		confLines[i] = n[1 : len(n)-1] // truncate []
	}
	return confLines
}
