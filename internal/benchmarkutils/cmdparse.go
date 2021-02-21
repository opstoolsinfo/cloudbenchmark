package benchmarkutils

import (
	"flag"
	"fmt"
	"os"
	//"github.com/spf13/viper"
)

func CmdParse() {
	var provider string
	var region string
	//var az []string
	var database string
	//fmt.Println(os.Args)
	flag.StringVar(&provider, "provider", os.Args[1], "The cloud provider for the benchmark(aws|gcp|generic)")
	flag.StringVar(&region, "region", os.Args[2], "The region to run this benchmark")
	flag.StringVar(&database, "database", os.Args[3], "The database type for the benchmark")
	flag.Parse()
	fmt.Println(provider)
}
