package config

import "fmt"
import "config/subconfig"

func LoadConfig() {
	fmt.Println("load config")
	subconfig.LoadConfig()
}
