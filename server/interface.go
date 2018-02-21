package server

import (
	"net/http"
	"sync"

	"github.com/lestrrat-go/slack/internal/option"
)

const (
	optkeyPrefix = "prefix"
)

type Option interface {
	option.Interface
}

type Server struct {
	handlers   map[string]http.Handler
	muHandlers sync.RWMutex
	prefix     string
}
