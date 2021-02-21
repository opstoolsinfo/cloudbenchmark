package main

import (
	"./common/elasticbench"
)

func main() {
	elasticbench.Generate()
	elasticbench.Index()
	elasticbench.SearchAll()
	elasticbench.GetBench()
	elasticbench.UpdateBench()
	elasticbench.DeleteBench()
}
