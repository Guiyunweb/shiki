package conf

// Redis配置
type RedisConf struct {
	RedisDefaultHost      string
	RedisDefaultInitConns int
	RedisDefaultMaxConns  int
}

func LoadRedis() *RedisConf {
	RedisDefaultHost := Conf.GetString("REDIS_DEFAULT_HOST")
	if RedisDefaultHost == "" {
		RedisDefaultHost = "127.0.0.1:6379"
	}
	RedisDefaultInitConns := Conf.GetInt("REDIS_DEFAULT_INIT_CONNS")
	if RedisDefaultInitConns == 0 {
		RedisDefaultInitConns = 1
	}
	RedisDefaultMaxConns := Conf.GetInt("REDIS_DEFAULT_MAX_CONNS")
	if RedisDefaultMaxConns == 0 {
		RedisDefaultMaxConns = 1
	}
	return &RedisConf{
		RedisDefaultHost,
		RedisDefaultInitConns,
		RedisDefaultMaxConns,
	}
}
