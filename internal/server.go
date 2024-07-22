package internal

import "github.com/heywinit/Minnegota/internal/config"

type Server struct {
	properties map[string]string
}

func NewServer(config config.ServerNetworkConfig) Server {
	return Server{}
}
