# go-slack

[![Build Status](https://travis-ci.org/lestrrat/go-slack.png?branch=master)](https://travis-ci.org/lestrrat/go-slack)

[![GoDoc](https://godoc.org/github.com/lestrrat/go-slack?status.svg)](https://godoc.org/github.com/lestrrat/go-slack)

# Status

* Many APIs are still unimplemented (please file an issue!)
* RTM events are not covered entirely yet (please file an issue!)

Missing parts are missing only because the author does not have immediate need for them. With proper prodding, I will gladly take on implementing them.

Please see [#4](https://github.com/lestrrat/go-slack/issues/4) for a list of currently known unimplemented methods.

# Features

* Google API style library
* Full support for context.Context

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

# Calling Conventions

The API is constructed along the same lines of Google API Go libraries (https://google.golang.org/api), so if you are familiar with that style you should have no problem. But if you have not used their API, I'm sure you're confused - I know I was when I first saw their API.

This section will walk you through how the APIs in this REST client generally work. The basic idea is that you have a central client, which you can create via `New`:

```go
  client := slack.New(token)
```

The client is nothing but a placeholder for other "services". For example, to
use the slack APIs for `chat.*` endpoints, lookup the `ChatService` object
which can be obtained by calling the `Chat()` method:

```go
  service := client.Chat()
```

The `ChatService` object can construct an intermediate object, which are called
the `Call` objects. These objects exist to accumulate parameters that you want to
ultimately call the API endpoint with. In the initial call to construct the objects, you will have to enter the mandatory parameters. For example, to start a
`PostMessage` call (which is supposed to access `chat.postMessage` endpoint), you do:

```go
  call := service.PostMessage(channel)
```

The `channel` parameter is required, so you must pass that to the `PostMessage` method (you would think that `text` is always required, but you could either provide `text`, or an _attachment_ with text, so it's not always the case)

Once you have the `call` object, you can specify additional parameters which are
optional. The following would specify a `text` attribute for this call.

```go
  call.Text(yourMessage)
```

Finally, when you are done tweaking your call object, you should call `Do`, which will fire the actual request.

```go
  res, err := call.Do(ctx)
```

And you are done. Don't for get to pass a `context.Context` to all `Do` calls.

Also, for your convenience these `call` object methods can all be chained. So the above example could look like the following:

```go
  res, err := client.Users().PostMessage(channel).Text(yourMessage).Do(ctx)
```

# Acknowledgements

Based on github.com/nlopes/slack.