package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Pair struct {
	n1, n2 string
}

var pairs = map[int]Pair{
	0: Pair{"12", "2222"},
	1: Pair{"123", "-12"},
	2: Pair{"+90", "-10000000"},
	3: Pair{"AAAA", "CCCCC"},
	4: Pair{"10000000000", "2000000000"},
	5: Pair{"1", "AAAAAAAAAAAAAAA"},
}

func main() {

	host := os.Args[1]
	port := os.Args[2]
	para, _ := strconv.Atoi(os.Args[3])
	cycle, _ := strconv.Atoi(os.Args[4])

	baseUrl := "http://" + host + ":" + port + "/"

	start := time.Now()

	fmt.Println("Base Url=", baseUrl, "paralell=", para, " cycle par round=", cycle)

	ch := make(chan int, para)
	for i := 0; i < para; i++ {
		go func() {
			round(baseUrl, cycle)
			ch <- 0
		}()
	}

	// block until all round end.
	for i := 0; i < para; i++ {
		<-ch
	}

	end := time.Now()

	fmt.Println("end by ", (end.Sub(start)))

}

func round(base string, cycle int) {
	for i := 0; i < cycle; i++ {

		p := pairs[i%len(pairs)]
		Cycle(base, p.n1, p.n2)
	}
}

func Cycle(base string, n1 string, n2 string) {

	ok := get(base, "")
	if ok {
		ok = get(base, "calc")
	}
	if ok {
		post(base, "calc", n1, n2)
	}

}

func get(base string, content string) bool {

	r, err := http.Get(base + content)

	if err != nil {
		fmt.Println(content, " error", err)
		return false
	}

	defer r.Body.Close()
	ioutil.ReadAll(r.Body)
	return true
}

func post(base string, content string, n1 string, n2 string) bool {
	r, err := http.PostForm(base+content, url.Values{"n1": {n1}, "n2": {n2}})

	//defer r.Body.Close()

	if err != nil {
		fmt.Println(content, " error", err)
		return false
	}
	defer r.Body.Close()
	ioutil.ReadAll(r.Body)

	return true
}
