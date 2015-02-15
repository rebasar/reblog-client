package main

import (
	"github.com/ogier/pflag"
	"github.com/spf13/viper"
)

var Server string = "localhost"
var Database string = "reblog"
var InputFile string = "-"
var Help bool = false

func initConfig() {
	viper.SetConfigName("reblog")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath("$HOME/.config")
	viper.ReadInConfig()
	viper.SetDefault("server", "localhost")
	viper.SetDefault("db", "reblog")
	pflag.StringVarP(&Server, "server", "s", viper.GetString("server"), "MongoDB Server address")
	pflag.StringVarP(&Database, "db", "d", viper.GetString("db"), "Database name on the MongoDB Server")
	pflag.StringVarP(&InputFile, "file", "f", "-", "File name to read the entry from. Omit or use \"-\" for reading from stdin")
	pflag.BoolVarP(&Help, "help", "h", false, "Get help about usage")
	pflag.Parse()
}
