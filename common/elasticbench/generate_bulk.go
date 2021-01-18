package elasticbench

import (
	"../../configs"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var address string
var BulkAddress string
var BenchFile string
var BenchSize int
var IndexName string

func init() {
	var conf = configs.TestConfReader
	host := conf.Get("elasticsearch.host")
	port := conf.Get("elasticsearch.port")
	file := conf.Get("elasticsearch.bulk_file")
	size := conf.Get("elasticsearch.bench_size")
	index := conf.Get("elasticsearch.index_name")
	BenchSize, _ = strconv.Atoi(size.(string))
	BenchFile = fmt.Sprintf("%s",file)
	IndexName = fmt.Sprintf("%s",index)
	address = fmt.Sprintf("%s%s%s%s","http://",host,":",port)
	BulkAddress = fmt.Sprintf("%s%s", address, "/_bulk?pretty")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Generate() {
	var index = "{\"index\":{\"_index\":\"" + IndexName + "\"}}"
	file, err := os.Create(BenchFile)
	check(err)
	start := time.Now()
	fmt.Println("Bench Size is: ", BenchSize)
	for i := 1; i < BenchSize; i++ {
		doc := GenerateDoc(i)
		_, err = file.WriteString(index + "\n" + string(doc) + "\n")
		check(err)
	}
	elapsed := time.Since(start)
	log.Printf("Generation time %s", elapsed)
}
