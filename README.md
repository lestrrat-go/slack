# go-slack

[![Build Status](https://travis-ci.org/lestrrat/go-slack.png?branch=master)](https://travis-ci.org/lestrrat/go-slack)

[![GoDoc](https://godoc.org/github.com/lestrrat/go-slack?status.svg)](https://godoc.org/github.com/lestrrat/go-slack)

# Synopsis

Simple REST Client:

```go
package slack_test

import (
  "context"
  "fmt"
  "os"

  "github.com/lestrrat/go-slack"
)

func ExampleClient() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  token := os.Getenv("SLACK_TOKEN")
  cl := slack.New(token)

  // check if we are connected
  authres, err := cl.Auth().Test().Do(ctx)
  if err != nil {
    fmt.Printf("failed to test authentication: %s\n", err)
    return
  }
  fmt.Printf("%#v\n", authres)

  // simplest possible message
  chatres, err := cl.Chat().PostMessage("@username").
    Text("Hello, World!").
    Do(ctx)
  if err != nil {
    fmt.Printf("failed to post messsage: %s\n", err)
    return
  }
  fmt.Printf("%#v\n", chatres)
}
```

# RTM

See the [README in the rtm/ directory](./rtm/README.md).

# Acknowledgements

Based on github.com/nlopes/slack.