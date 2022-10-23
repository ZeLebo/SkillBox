// the file suppossed to run lots of servers and proxy
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"user/cmd/handlers"
	"user/internal/database"
	"user/internal/domain"
	"user/internal/service"
)

// the same as main, but you can decide how many instances of server you want
// run this file with int CLI
func spawnServers(amount int, shutdown chan os.Signal) {
	for i := 0; i < amount; i++ {
		i := i
		go func() {
			port := 60491 + i
			address := "localhost:" + strconv.Itoa(port)
			fmt.Println(address)
			logger := log.Default()
			cfg := domain.GetDatabaseConfig()

			client, err := database.NewClient(cfg, logger)
			if err != nil {
				log.Fatalf("error initializing database: %s", err.Error())
			}

			handlerService := service.NewService(client, logger)

			mainRouter := mux.NewRouter()
			handler := handlers.NewRequestHandler(handlerService, logger)

			handler.Routes(mainRouter)

			go func() {
				err := http.ListenAndServe(address, mainRouter)
				if err != nil {
					logger.Fatal("Can't start server" + err.Error())
				}
			}()
			<-shutdown
			defer os.Exit(0)
		}()
	}
}

var counter int
var amount int

func main() {
	var err error
	// check if the user entered the amount of servers
	if len(os.Args) < 2 {
		log.Fatalln("Please enter the amount of servers")
	}
	amount, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Cannot parse the command line arguments")
	}

	shutdown := make(chan os.Signal, amount)
	spawnServers(amount, shutdown)

	proxyIP, err := domain.GetIp("proxy")
	if err != nil {
		log.Fatalln("Cannot start the proxy server")

	}
	http.HandleFunc("/", handle)
	fmt.Println("Proxy server has started")
	go func() {
		err := http.ListenAndServe(proxyIP, nil)
		if err != nil {
			log.Fatal("Cannot start the proxy server")
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		os.Interrupt,
		os.Kill,
	)

	<-signalChan
	for i := 0; i < amount; i++ {
		shutdown <- os.Interrupt
	}
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
