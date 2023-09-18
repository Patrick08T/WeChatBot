package send_weather

import (
	"bot/src/sdk/weather"
	"bot/src/utils/tick"
	"github.com/eatmoreapple/openwechat"
)

func DailyWeatherToFriend(self *openwechat.Self, friends openwechat.Friends, name string) {
	fs := friends.Search(1, func(friend *openwechat.Friend) bool { return friend.NickName == name })
	tick.DailyTicker(9, 0, 0, func() {
		resp := weather.GetTodayWeather()
		if resp != nil {
			for _, friend := range fs {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	})
}
