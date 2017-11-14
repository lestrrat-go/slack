package mockserver_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/lestrrat/go-slack/server/mockserver"
)

func ExampleMockServer() {
	ts := httptest.NewServer(mockserver.New())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/channels.info")
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	// OUTPUT:
	// {"ok":true,"error":{"code":0,"msg":""},"ts":"","id":"123456789ABCDEFG","created":233431200,"is_open":false,"creator":"yoda","is_archived":false,"is_group":false,"is_mpim":false,"members":["obiwan","lukeskywalker"],"name":"jedis","name_normalized":"jedis","num_members":2,"previous_names":null,"purpose":{"value":"There is no emotion, there is peace.\nThere is no ignorance, there is knowledge.\nThere is no passion, there is serenity.\nThere is no chaos, there is harmony.\nThere is no death, there is the Force.","creator":"yoda","last_set":233431200},"topic":{"value":"Jedi meetup and drinks next Tuesday","creator":"yoda","last_set":233431200},"is_channel":false,"is_general":false,"is_member":false,"is_org_shared":false,"is_shared":false}
}
