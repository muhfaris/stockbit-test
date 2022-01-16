package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/gateway/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile = ""

var rootCmd = &cobra.Command{
	Use:   "movies",
	Short: "API Moviews",
	Run: func(cmd *cobra.Command, args []string) {
		app := configs.CreateApp()
		handler.InitRouter(app)

	},
}

func init() {
	cobra.OnInitialize(initconfig)
}

func initconfig() {
	viper.SetConfigType("toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// search config in home directory with name "config" (without extension)
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}

	//read env
	viper.AutomaticEnv()

	// if a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Config application:", err)
	}

	log.Println("using config file:", viper.ConfigFileUsed())
}

// Execute is root function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
