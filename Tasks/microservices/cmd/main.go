package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os/exec"
	s "user/pkg/Server"
)

// as proxy server need to push the request directly to the target server

func main() {
	var (
		proxyAddr = "0.0.0.0:60490"
		addr1     = "0.0.0.0:60494"
		addr2     = "0.0.0.0:60495"
		counter   = 0
	)

	// run command: go run cmd/main.go
	go func() {
		err := exec.Command("go", "run", "first.go").Run()
		if err != nil {
			log.Error(err.Error())
		}
	}()
	go func() {
		err := exec.Command("go", "run", "second.go").Run()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	server := s.NewServer(proxyAddr)
	log.Info("The server is up and running at ", server.Addr, "\n")

	// handle all requests to the proxy server
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if counter == 0 {
			http.Redirect(w, r, addr1, http.StatusFound)
			counter++
		} else {
			http.Redirect(w, r, addr2, http.StatusFound)
			counter--
		}
	})

	log.Info(server.ListenAndServe())
}
