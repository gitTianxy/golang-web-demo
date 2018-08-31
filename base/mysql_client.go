package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlClient() MySQLClient {
	client := MySQLClient{}
	err := client.init(GetDBConfig())
	CheckErr(err)
	return client
}

type MySQLClient struct {
	conf MysqlConfig
	pool *sql.DB
}

func (mc *MySQLClient) init(conf MysqlConfig) (err error) {
	mc.conf = conf
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		conf.User(),
		conf.Password(),
		conf.Host(),
		conf.Port(),
		conf.Db())
	// Open 全局一个实例只需调用一次
	mc.pool, err = sql.Open("mysql", uri)
	if err != nil {
		return err
	}
	//使用前 Ping, 确保 DB 连接正常
	err = mc.pool.Ping()
	if err != nil {
		return err
	}
	// 设置最大连接数，一定要设置 MaxOpen
	mc.pool.SetMaxIdleConns(conf.MaxIdle())
	mc.pool.SetMaxOpenConns(conf.MaxOpen())
	return nil
}

func (client MySQLClient) GetPool() (*sql.DB)  {
	return client.pool
}
