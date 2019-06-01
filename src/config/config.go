package config

import "fmt"
import "config/subconfig"

func init() {
	fmt.Println("call config init")
}

func LoadConfig() {
	fmt.Println("load config")
	subconfig.LoadConfig()
}
