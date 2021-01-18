package elasticbench

import (
	"../../internal/benchmarkutils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func DeleteBench()  {
	start := time.Now()
	var DeleteErrors int
	client := &http.Client{}
	for _, id := range DocumentIds {
		GetAddress := fmt.Sprintf("%s%s%s",address,"/benchmark/_doc/", id)
		request, err := http.NewRequest("DELETE", GetAddress, nil)
		benchmarkutils.ErrorCheck(err)
		response, err := client.Do(request)
		if response.StatusCode != 200 {
			DeleteErrors += 1
		}
	}
	elapsed := time.Since(start)
	log.Printf("Delete total time %s, with %d errors", elapsed, DeleteErrors)
}