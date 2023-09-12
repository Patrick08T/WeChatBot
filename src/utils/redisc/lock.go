package redisc

import (
	"ZhuRong/backend/src/util/log"
	"ZhuRong/backend/src/util/zuuid"
	"context"
	"time"
)

type Lock struct {
	key  string
	uuid int64
}

func NewLock(key string) (Lock, bool) {
	lock := Lock{
		key:  key,
		uuid: zuuid.NewUUID(),
	}
	ok, err := lock.TryLock(3 * time.Minute)
	if err != nil {
		return lock, false
	}
	return lock, ok
}

func (lock Lock) TryLock(expire time.Duration) (bool, error) {
	if expire <= 0 || expire > 5*time.Minute {
		panic("expire must be in (0, 5min]")
	}
	cmd := Client().SetNX(context.TODO(), lock.key, lock.uuid, expire)
	if cmd.Err() != nil {
		log.ERRORF("try lock failed. %v", cmd.Err().Error())
		return false, cmd.Err()
	}
	return cmd.Val(), nil
}

const unlockLua = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end
`

// UnLock 解分布式锁
func (lock Lock) UnLock() error {
	cmd := Client().Eval(context.TODO(), unlockLua, []string{lock.key}, lock.uuid)
	if cmd.Err() != nil {
		log.ERRORF("unlock failed. %v", cmd.Err().Error())
		return cmd.Err()
	}
	return nil
}
