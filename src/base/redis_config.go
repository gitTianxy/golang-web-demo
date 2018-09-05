package base

func GetRedisConfig() RedisConfig {
	return RedisConfig{
		"localhost",
		6379,
		"", // no password set
		0,  // use default DB
	}
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}
