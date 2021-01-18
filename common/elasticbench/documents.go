package elasticbench

import (
	"../../internal/benchmarkutils"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"io/ioutil"
	"strconv"
	"strings"
)

type document struct {
	ID int `json:"id"`
	TITLE string `json:"title"`
	GAME string `json:"game"`
	TRACK string `json:"track"`
	CASE string `json:"case"`
	ROLE string `json:"role"`
	DANCE string `json:"dance"`
	ROMAN string `json:"roman"`
	MEME string `json:"meme"`
}

func GenerateDoc(i int) []byte {
	doc := &document{
		ID: i,
		TITLE: strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
		GAME: strings.Join([]string{"G", strconv.Itoa(i)}, " "),
		TRACK: strings.Join([]string{"T", strconv.Itoa(i)}, " "),
		CASE: strings.Join([]string{"C", strconv.Itoa(i)}, " "),
		ROLE: strings.Join([]string{"Role", strconv.Itoa(i)}, " "),
		DANCE: strings.Join([]string{"Dance", strconv.Itoa(i)}, " "),
		ROMAN: strings.Join([]string{"Roman", strconv.Itoa(i)}, " "),
		MEME: strings.Join([]string{"Meme", strconv.Itoa(i)}, " "),
	}
	empJSON, err := json.Marshal(doc)
	benchmarkutils.ErrorCheck(err)
	return empJSON
}

type TookTime struct {
	Took int64 `json:"took"`
}

func ReadFile() *string {
	data, err := ioutil.ReadFile(BenchFile)
	benchmarkutils.ErrorCheck(err)
	dat := string(data)
	return &(dat)
}

func Connect() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, _ := elasticsearch.NewClient(cfg)
	return es
}