package news

import (
	"bot/src/sdk/news"
	"bot/src/utils/tick"
	"github.com/eatmoreapple/openwechat"
)

func DailyNewsToFriend(self *openwechat.Self, friends openwechat.Friends, name string) {
	fs := friends.Search(1, func(friend *openwechat.Friend) bool { return friend.NickName == name })
	tick.DailyTicker(8, 30, 0, func() {
		resp := news.GetNews()
		if resp != nil {
			for _, friend := range fs {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	})
}
func DailyNewsToFriends(self *openwechat.Self, friends openwechat.Friends) func() {
	return func() {
		resp := news.GetNews()
		if resp != nil {
			for _, friend := range friends {
				self.SendTextToFriend(friend, resp.ToString())
			}
		}
	}
}
