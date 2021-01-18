package elasticbench

import (
	"../../internal/benchmarkutils"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func UpdateBench() {
	var count int
	var PostDoc []byte
	var UpdateErrors int
	start := time.Now()
	for _, id := range DocumentIds {
		count += 1
		PostAddress := fmt.Sprintf("%s%s%s", address, "/benchmark/_doc/", id)
		PostDoc = GenerateDoc(count)
		DocBytes := bytes.NewBuffer(PostDoc)
		response, err := http.Post(PostAddress, "application/json",DocBytes)
		benchmarkutils.ErrorCheck(err)
		if response.StatusCode != 200 {
			UpdateErrors += 1
		}
	}
	elapsed := time.Since(start)
	log.Printf("Update total time %s, with %d errors", elapsed, UpdateErrors)
}

