package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	apachelog "github.com/lestrrat-go/apache-logformat"
	"github.com/lestrrat-go/slack/server"
	"github.com/lestrrat-go/slack/server/proxyserver"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	var token string
	var listen string
	flag.StringVar(&listen, "listen", ":8080", "listen address")
	flag.StringVar(&token, "token", "", "slack API token")
	flag.Parse()

	if len(token) == 0 {
		return errors.New(`-token is required`)
	}

	h := proxyserver.New(token)
	s := server.New()
	h.InstallHandlers(s)

	log.Printf("Listening on address %s", listen)
	return http.ListenAndServe(listen, apachelog.CombinedLog.Wrap(s, os.Stdout))
}
