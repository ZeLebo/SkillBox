package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"user/configs"
	s "user/pkg/server"
)

// the same as main, but you can decide how many instances of server you want
// run this file with int CLI
func spawnServers(amount int) {
	for i := 0; i < amount; i++ {
		i := i
		port := 60491 + i
		address := "localhost:" + strconv.Itoa(port)
		fmt.Println(address)
		go func() {
			server := s.NewServer(address)
			done := make(chan bool)
			err := server.ListenAndServe()
			if err != nil {
				log.Println(err.Error())
			}
			done <- true

			server.WaitShutdown()

			<-done
		}()
	}
}

var counter int
var amount int

func main() {
	var err error
	amount, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Cannot parse the command line arguments")
	}
	spawnServers(amount)

	proxyIP, err := configs.GetIp("proxy")
	if err != nil {
		log.Fatalln("Cannot start the proxy server")

	}
	http.HandleFunc("/", handle)
	fmt.Println("Proxy server has started")
	log.Fatalln(http.ListenAndServe(proxyIP, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	var (
		resp *http.Response
		req  *http.Request
		err  error
	)
	if counter < 0 || counter >= amount {
		counter = 0
	}
	port := 60491 + counter

	dest := "http://localhost:" + strconv.Itoa(port) + r.URL.Path
	req, err = http.NewRequest(r.Method, dest, r.Body)
	if err != nil {
		return
	}
	counter++

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(dest)
		log.Fatalln("Error: ", err.Error())
	}

	//goland:noinspection GoDeprecation
	content, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
