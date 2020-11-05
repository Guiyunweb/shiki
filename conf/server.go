package conf

import "fmt"

// Server配置
type ServerConf struct {
	HttpPost string
}

func LoadServer() *ServerConf {
	HttpPost := Conf.GetInt("HTTP_PORT")
	if HttpPost == 0 {
		HttpPost = 8080
	}
	return &ServerConf{
		HttpPost: fmt.Sprintf(":%d", HttpPost),
	}
}
