package mockserver_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"

	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/server"
	"github.com/lestrrat/go-slack/server/mockserver"
)

const token = "AbCdEfG"

func ExampleMockServer() {
	h := mockserver.New(token)
	s := server.New()
	h.InstallHandlers(s)
	ts := httptest.NewServer(s)
	defer ts.Close()

	cl := slack.New(token, slack.WithAPIEndpoint(ts.URL))

	channel, err := cl.Channels().Info("jedi").Do(context.Background())
	if err != nil {
		log.Printf(`expected channels.info to succeed: %s`, err)
		return
	}

	json.NewEncoder(os.Stdout).Encode(channel)
	// OUTPUT:
	// {"id":"123456789ABCDEFG","created":233431200,"is_open":false,"creator":"yoda","is_archived":false,"is_group":false,"is_mpim":false,"members":["obiwan","lukeskywalker"],"name":"jedis","name_normalized":"jedis","num_members":2,"previous_names":null,"purpose":{"value":"There is no emotion, there is peace.\nThere is no ignorance, there is knowledge.\nThere is no passion, there is serenity.\nThere is no chaos, there is harmony.\nThere is no death, there is the Force.","creator":"yoda","last_set":233431200},"topic":{"value":"Jedi meetup and drinks next Tuesday","creator":"yoda","last_set":233431200},"is_channel":false,"is_general":false,"is_member":false,"is_org_shared":false,"is_shared":false}
}
