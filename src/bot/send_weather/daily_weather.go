package send_weather

import (
	"WeChatBot/src/sdk/weather"
	"WeChatBot/src/utils/tick"
	"github.com/eatmoreapple/openwechat"
)

func DailyWeatherToFriend(self *openwechat.Self, friends openwechat.Friends, name string) {
	fs := friends.Search(1, func(friend *openwechat.Friend) bool { return friend.NickName == name })
	tick.DailyTicker(22, 36, 0, func() {
		resp := weather.GetTodayWeather()
		if resp != nil {
			for _, friend := range fs {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	})
}
