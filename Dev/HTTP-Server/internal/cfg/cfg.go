package cfg

type ServerConfig struct {
	Host string
	Port string
}

func NewConfig() *ServerConfig {

	return &ServerConfig{
		Host: "127.0.0.1",
		Port: "8000",
	}

}
