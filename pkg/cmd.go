package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadCfg() {
	var configFile *string = flag.String("c", "myConfig", "Setting the configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)

	if err == nil {
		fmt.Println("Using User Specified Configuration file!")
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
		return
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("OS.reboot") {
		fmt.Println("OS.reboot:", viper.Get("OS.reboot"))
	} else {
		fmt.Println("OS.reboot not set!")
	}

	if viper.IsSet("OS.update_os") {
		fmt.Println("OS.update_os:", viper.Get("OS.update_os"))
	} else {
		fmt.Println("OS.update_os not set!")
	}

}
