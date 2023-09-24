package send_weather

import (
	"bot/src/sdk/weather"
	"bot/src/utils/tick"
	"github.com/eatmoreapple/openwechat"
)

func DailyWeatherToFriend(self *openwechat.Self, friends openwechat.Friends, name string) {
	fs := friends.Search(1, func(friend *openwechat.Friend) bool { return friend.NickName == name })
	tick.NineAMTasks(func() {
		resp := weather.GetTodayWeather()
		if resp != nil {
			for _, friend := range fs {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	})
}
func DailyWeatherToFriends(self *openwechat.Self, friends openwechat.Friends) {
	tick.NineAMTasks(func() {
		resp := weather.GetTodayWeather()
		if resp != nil {
			for _, friend := range friends {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	})
}
