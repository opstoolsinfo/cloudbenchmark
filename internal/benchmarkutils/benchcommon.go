package benchmarkutils

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
