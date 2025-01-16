package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "task-cli is a todo app",
	Long:  `task-cli will help you get more done in less time.`,
}

var dataFile string
var cfgFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config",
		"",
		"config file (default is $HOME/task-cli.yaml)")

	rootCmd.PersistentFlags().StringVar(&dataFile,
		"datafile",
		home+string(os.PathSeparator)+"task-cli.json",
		"data file to store todos")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType(".task-cli")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("task-cli")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file", viper.ConfigFileUsed())
	}
}
