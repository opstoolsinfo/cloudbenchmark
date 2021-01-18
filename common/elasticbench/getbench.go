package elasticbench

import (
	"../../internal/benchmarkutils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetBench() {
	start := time.Now()
	var GetErrors int
	for _, id := range DocumentIds {
		GetAddress := fmt.Sprintf("%s%s%s",address,"/benchmark/_doc/", id)
		response, err := http.Get(GetAddress)
		benchmarkutils.ErrorCheck(err)
		if response.StatusCode != 200 {
			GetErrors += 1
		}
	}
	elapsed := time.Since(start)
	log.Printf("Get total time %s, with %d errors", elapsed, GetErrors)
}
