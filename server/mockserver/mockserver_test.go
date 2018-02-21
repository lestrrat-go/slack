package mockserver_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"

	"github.com/lestrrat-go/slack"
	"github.com/lestrrat-go/slack/server"
	"github.com/lestrrat-go/slack/server/mockserver"
)

const token = "AbCdEfG"

func ExampleMockServer() {
	h := mockserver.New(mockserver.WithToken(token))
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

	buf, _ := json.MarshalIndent(channel, "", "  ")
	os.Stdout.Write(buf)
	// OUTPUT:
	//{
	//   "id": "C1J3D1ZRUL3",
	//   "created": 233431200,
	//   "is_open": false,
	//   "creator": "U0000001",
	//   "is_archived": false,
	//   "is_group": false,
	//   "is_mpim": false,
	//   "members": [
  //     "U0123456",
  //     "U0012345",
  //     "U0000001"
	//   ],
	//   "name": "jedis",
	//   "name_normalized": "jedis",
	//   "num_members": 3,
	//   "purpose": {
	//     "value": "There is no emotion, there is peace.\nThere is no ignorance, there is knowledge.\nThere is no passion, there is serenity.\nThere is no chaos, there is harmony.\nThere is no death, there is the Force.",
	//     "creator": "yoda",
	//     "last_set": 233431200
	//   },
	//   "topic": {
	//     "value": "Jedi meetup and drinks next Tuesday",
	//     "creator": "yoda",
	//     "last_set": 233431200
	//   },
	//   "is_channel": true,
	//   "is_general": false,
	//   "is_member": true,
	//   "is_org_shared": false,
	//   "is_shared": false
	//}
}
