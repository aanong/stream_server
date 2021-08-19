package server

type P2PServerConfig struct {
	Host          string
	Port          int
	CertFile      string
	KeyFile       string
	HTMLRoot      string
	WebSocketPath string
}

func DefaultConfig() P2PServerConfig {
	return P2PServerConfig{
		Host:          "0.0.0.0",
		Port:          8000,
		HTMLRoot:      "html",
		WebSocketPath: "/ws",
	}
}
