package redisc

import "strconv"

func UserLockKey(userId int64) string {
	return "user_lock_" + strconv.FormatInt(userId, 10)
}
