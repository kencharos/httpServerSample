package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// args[0] is program name.
	path := os.Args[1]
	port := "4000"
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/ GET[", time.Now(), "]")
		w.Header().Set("Content-Type", "text/html")
		file, _ := ioutil.ReadFile(path + "/welcom.html")

		fmt.Fprint(w, string(file))
	})

	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("/calc ", r.Method, "[", time.Now(), "]")

		w.Header().Set("Content-Type", "text/html")
		if r.Method == "GET" {
			file, _ := ioutil.ReadFile(path + "/calc.html")
			// string is []byte alias
			fmt.Fprintf(w, string(file))
		} else {
			n1 := r.FormValue("n1")
			n2 := r.FormValue("n2")

			res := "Wrong input"
			if isNumber(n1) && isNumber(n2) {
				i1, _ := strconv.Atoi(n1)
				i2, _ := strconv.Atoi(n2)
				res = strconv.Itoa(i1 + i2)
			}

			file, _ := ioutil.ReadFile(path + "/result.html")

			fmt.Fprintf(w, string(file), res)

		}
	})
	fmt.Println("Address = :" + port)
	fmt.Println("path = " + path)
	err := http.ListenAndServe(":"+port, nil)
	fmt.Println(err)
}

func isNumber(input string) bool {
	if input == "" {
		return false
	}
	ok, _ := regexp.MatchString("^[+-]?[0-9]+$", input)
	return ok

}
