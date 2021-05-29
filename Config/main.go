package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

//var DevEnvName, StageEnvName, ProdEnvName string = getEnvironmentName()

type Configurations struct {
	DevEnvName struct {
		DbName string `yaml:"db_name"`
	} `yaml:"dev"`
	StageEnvName struct {
		DbName string `yaml:"db_name"`
	} `yaml:"stage"`
	ProdEnvName struct {
		DbName string `yaml:"db_name"`
	} `yaml:"prod"`
}

const (
	dev   = "dev"
	stage = "stage"
	prod  = "prod"
)

/*type Configurations struct {
	DevEnvName struct {
		DbName string `yaml:"db_name"`
	} `yaml:"dev" yaml:"stage" yaml:"prod"`
}*/

func main() {
	//fmt.Printf("VAlue received from method: %v %v %v \n", DevEnvName, StageEnvName, ProdEnvName)
	configViaViper()
	//var cfg Configurations
	//readFile(&cfg)
	//envDbName := getEnvironmentNameForDb()
	//fmt.Println(envDbName)
	//if dev == envDbName {
	//	fmt.Printf("%v \n", cfg.DevEnvName.DbName)
	//} else if stage == envDbName {
	//	fmt.Printf("%v \n", cfg.StageEnvName.DbName)
	//} else if prod == envDbName {
	//	fmt.Printf("%v \n", cfg.ProdEnvName.DbName)
	//}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Configurations) {
	f, err := os.Open("src/main/go/Config/config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)

	fmt.Println(*cfg)
	if err != nil {
		processError(err)
	}

}

/*func getEnvironmentName() (string, string, string) {
	os.Setenv("ENV_NAME", "usstagingms")
	fmt.Printf("Value fetched from env %v \n", os.Getenv("ENV_NAME"))
	var devEnvName, stageEnvName, prodEnvName string
	switch os.Getenv("ENV_NAME") {
	case "usdevms", "eudevms":
		devEnvName = "dev"
	case "usstagingms", "eustagingms":
		stageEnvName = "stage"
	case "prodms", "euprodms":
		prodEnvName = "prod"

	}
	fmt.Printf("Environment name received %v %v %v \n", devEnvName, stageEnvName, prodEnvName)

	return devEnvName, stageEnvName, prodEnvName
}*/

func getEnvironmentNameForDb() string {
	os.Setenv("ENV_NAME", "usstagingms")
	fmt.Printf("Value fetched from env %v \n", os.Getenv("ENV_NAME"))
	var envNameForDb string
	switch strings.ToLower(os.Getenv("ENV_NAME")) {
	case "usdevms", "eudevms":
		envNameForDb = "dev"
	case "usstagingms", "eustagingms":
		envNameForDb = "stage"
	case "prodms", "euprodms":
		envNameForDb = "prod"

	}
	fmt.Printf("Environment name received %v \n", envNameForDb)

	return envNameForDb
}
func configViaViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath("src/main/go/Config/") // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var c Configurations
	fmt.Println(viper.Unmarshal(&c))
}
