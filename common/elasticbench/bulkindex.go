package elasticbench

import (
	"../../internal/benchmarkutils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Index()  {
	data := ReadFile()
	start := time.Now()
	res, error := http.Post(BulkAddress, "application/json", strings.NewReader(*data))
	benchmarkutils.ErrorCheck(error)
	elapsed := time.Since(start)
	took := Results(res)
	log.Printf("Total %v, Took %vms", elapsed, took)
}

func Results(response *http.Response) int64 {
	defer response.Body.Close()
	var took TookTime
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		benchmarkutils.ErrorCheck(err)
		error := json.Unmarshal(bodyBytes, &took)
		benchmarkutils.ErrorCheck(error)
	}
	return took.Took
}