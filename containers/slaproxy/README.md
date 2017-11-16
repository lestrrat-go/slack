# Docker image for slaproxy

This directory contains a Docker image for slaproxy, so you can 1) mock
API methods that have side-effects (such as posting a message), while
2) proxy everything else to the real server.

Currently this is intended to be used as a sidecar along side the component
that requires access to the Slack API.

## Building the image

The following command will create a image named `go-slack/slaproxy:latest`

```
make docker
```

## Configuring the container

(Note: this is TODO)
(TODO: do we need certificate management?)


| Environment Variable | Default Value | Description       |
|:---------------------|:--------------|:------------------|
| SLAPROXY_PORT        | 8080          | Port to listen to |
| SLAPROXY_TOKEN       | (null)        | Slack token to expect. Leaving it blank should work for most cases |

