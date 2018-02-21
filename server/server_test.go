package server_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	slack "github.com/lestrrat-go/slack"
	"github.com/lestrrat-go/slack/server"
)

func ExampleServer() {
	s := server.New()
	s.Handle("channels.info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// You can do what you want in this handler. Here, we're just
		// going to parse the request, and then print it out
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var c slack.ChannelsInfoCall
		if err := c.FromValues(r.Form); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		json.NewEncoder(os.Stdout).Encode(c)

		w.WriteHeader(http.StatusNoContent)
	}))

	var hs http.Server
	hs.Addr = ":9090"
	hs.Handler = s
	go hs.ListenAndServe()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hs.Shutdown(ctx)
}
