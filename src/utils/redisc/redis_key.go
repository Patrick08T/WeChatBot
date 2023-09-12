package redisc

import "fmt"

func RegisterKey(userName string) string {
	return fmt.Sprintf("user:%s", userName)
}

func UserIdKey(userId int64) string {
	return fmt.Sprintf("user_id:%v", userId)
}
