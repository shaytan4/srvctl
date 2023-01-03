package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

var Tpl *template.Template

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request, mydata map[string]string) {
	err := Tpl.ExecuteTemplate(w, "index.gohtml", mydata)
	HandleError(w, err)

	url_path := r.URL.Path[1:]
	fmt.Println(url_path)
	execCmd := mydata[url_path]
	fmt.Println("execCmd from map", execCmd)

	runCmd := exec.Command(execCmd)
	out, err := runCmd.Output()
	if err != nil {
		log.Println("could not run command: ", err)
	}
	log.Println("Output: ", string(out))

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

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
		viper.AddConfigPath("/tmp")
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
		log.Println(i, viper.Get(i))
		confLines[i] = fmt.Sprintf("%v", viper.Get(i))
	}

	// if viper.IsSet("OS.reboot") {
	// 	fmt.Println("OS.reboot:", viper.Get("OS.reboot"))
	// } else {
	// 	fmt.Println("OS.reboot not set!")
	// }

	// if viper.IsSet("OS.update_os") {
	// 	fmt.Println("OS.update_os:", viper.Get("OS.update_os"))
	// } else {
	// 	fmt.Println("OS.update_os not set!")
	// }
	return confLines
}
