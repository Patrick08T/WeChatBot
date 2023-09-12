package redisc

import (
	"ZhuRong/backend/src/util/log"
	ssh2 "ZhuRong/backend/src/util/net/ssh"
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

var redisClient redis.UniversalClient

func InitRedisBySSHPrivateKey(privateKeyPath string, remoteAddr string, addr []string, usr string, pw string, db, poolSize int) (redis.UniversalClient, error) {
	cli, err := ssh2.GetSSHClientByPrivateKey(privateKeyPath, remoteAddr)
	if nil != err {
		log.ERRORF("get ssh client by ssh private key failed. %v", err)
		return nil, err
	}
	redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		DB:         db,
		Addrs:      addr,
		MaxRetries: 3,
		PoolFIFO:   false,
		PoolSize:   poolSize,
		//MinIdleConns:       0,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return cli.Dial(network, addr)
		},
		ReadTimeout:  -2,
		WriteTimeout: -2,
	})
TryAgain:
	if err = redisClient.Ping(context.TODO()).Err(); nil != err {
		if strings.HasPrefix(err.Error(), "LOADING") {
			redisClient.Close()
			time.Sleep(time.Second)
			goto TryAgain
		}
		panic(fmt.Sprintf("connect to redis failed. %v", err))
	}
	return nil, nil
}

func InitRedis(addr []string, us, pw string, db, poolSize int) {
	redisClient = initRedisMgr(addr, us, pw, db, poolSize)
}

func initRedisMgr(addr []string, usr string, pw string, db, poolSize int) (client redis.UniversalClient) {
	client = redis.NewUniversalClient(&redis.UniversalOptions{
		Username:   usr,
		Password:   pw,
		DB:         db,
		Addrs:      addr,
		MaxRetries: 3,
		PoolFIFO:   false,
		PoolSize:   poolSize,
		//MinIdleConns:       0,
	})
	return client
}

func Client() redis.UniversalClient {
	return redisClient
}
