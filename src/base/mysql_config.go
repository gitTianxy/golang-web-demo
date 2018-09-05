package base

func GetDBConfig() MysqlConfig {
	return MysqlConfig{
		"localhost",
		3306,
		"golang_web_demo",
		"kevin",
		"1234",
		20,
		50,
	}
}

type MysqlConfig struct {
	host    string
	port    int
	db      string
	user    string
	pwd     string
	maxIdle int
	maxOpen int
}

func (config *MysqlConfig) Host() string  {
	return config.host
}

func (config *MysqlConfig) Port() int  {
	return config.port
}

func (config *MysqlConfig) User() string  {
	return config.user
}

func (config *MysqlConfig) Password() string  {
	return config.pwd
}

func (config *MysqlConfig) Db() string  {
	return config.db
}

func (config *MysqlConfig) MaxIdle() int  {
	return config.maxIdle
}

func (config *MysqlConfig) MaxOpen() int  {
	return config.maxOpen
}
