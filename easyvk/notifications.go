package easyvk

// Methods for working with notifications.
// https://vk.com/dev/notifications
type Notifications struct {
	vk *VK
}

// Returns a list of notifications about
// other users' feedback to the current
// user's wall posts.
// https://vk.com/dev/notifications.sendMessage
func (n *Notifications) SendMessage(userIDs, message, fragment string) error {
	params := make(map[string]string)
	params["user_ids"] = userIDs
	params["message"] = message
	params["fragment"] = fragment

	_, err := n.vk.Request("notifications.sendMessage", params)
	return err
}
