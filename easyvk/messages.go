package easyvk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Methods for send messages.
// https://vk.com/dev/messages
type Messages struct {
	vk *VK
}

// MessagesSendParams provides structure for
// parameters for get method.
// Returns only ids.
// https://vk.com/dev/messages.send
type MessagesSendParams struct {
	UserID          uint
	RandomID        int64
	PeerID          int
	Domain          string
	ChatId          uint
	UserIDs         string
	Message         string
	Lat             float32
	Long            float32
	Attachment      string
	ReplyTo         int
	ForwardMessages string
	StickerId       uint
	GroupId         uint
	Keyboard        *Keyboard
	Payload         string
}

// Returns a list of user IDs or detailed
// information about a user's friends.
// https://vk.com/dev/messages.send
func (m *Messages) Send(par MessagesSendParams) (uint, error) {
	params := make(map[string]string)
	if par.UserID != 0 {
		params["user_id"] = fmt.Sprint(par.UserID)
	}
	if par.RandomID != 0 {
		params["random_id"] = fmt.Sprint(par.RandomID)
	}
	if par.PeerID != 0 {
		params["peer_id"] = fmt.Sprint(par.PeerID)
	}
	if par.Domain != "" {
		params["domain"] = par.Domain
	}
	if par.ChatId != 0 {
		params["chat_id"] = fmt.Sprint(par.ChatId)
	}
	if par.UserIDs != "" {
		params["user_ids"] = par.UserIDs
	}
	if par.Message != "" {
		params["message"] = par.Message
	}
	if par.Lat != 0 {
		params["lat"] = fmt.Sprint(par.Lat)
	}
	if par.Long != 0 {
		params["long"] = fmt.Sprint(par.Long)
	}
	if par.Attachment != "" {
		params["attachment"] = par.Attachment
	}
	if par.ReplyTo != 0 {
		params["reply_to"] = fmt.Sprint(par.ReplyTo)
	}
	if par.ForwardMessages != "" {
		params["forward_messages"] = par.ForwardMessages
	}
	if par.StickerId != 0 {
		params["sticker_id"] = fmt.Sprint(par.StickerId)
	}
	if par.GroupId != 0 {
		params["group_id"] = fmt.Sprint(par.GroupId)
	}
	if par.Keyboard != nil {
		if data, err := json.Marshal(par.Keyboard); err == nil {
			params["keyboard"] = string(data)
		}
	}
	if par.Payload != "" {
		params["payload"] = par.Payload
	}

	resp, err := m.vk.Request("messages.send", params)
	if err != nil {
		return 0, err
	}

	messageId, err := strconv.Atoi(string(resp))

	return uint(messageId), err
}

// MessagesSetActivityParams provides structure for
// parameters for get method.
// https://vk.com/dev/messages.setActivity
type MessagesSetActivityParams struct {
	UserID  uint
	Type    TypeActivity
	PeerID  int
	GroupId uint
}

// Changes the status of a user as
// typing in a conversation.
// https://vk.com/dev/messages.setActivity
func (m *Messages) SetActivity(par MessagesSetActivityParams) (uint, error) {
	params := make(map[string]string)
	if par.UserID != 0 {
		params["user_id"] = fmt.Sprint(par.UserID)
	}

	params["type"] = fmt.Sprint(par.Type)

	if par.PeerID != 0 {
		params["peer_id"] = fmt.Sprint(par.PeerID)
	}

	if par.GroupId != 0 {
		params["group_id"] = fmt.Sprint(par.GroupId)
	}

	resp, err := m.vk.Request("messages.setActivity", params)
	if err != nil {
		return 0, err
	}

	messageId, err := strconv.Atoi(string(resp))

	return uint(messageId), err
}

// GetByIdParams provides structure for
// parameters for get method.
// https://vk.com/dev/messages.getById
type GetByIdParams struct {
	MessageIDs    string
	PreviewLength int
	Extended      bool
	Fields        string
	GroupId       uint
}

// GetByIdResponse describes
// https://vk.com/dev/messages.getById
type GetByIdResponse struct {
	Count int
	Items []MessageObject
}

// Returns messages by their IDs.
// https://vk.com/dev/messages.getById
func (m *Messages) GetById(par GetByIdParams) (GetByIdResponse, error) {
	params := make(map[string]string)

	params["message_ids"] = par.MessageIDs
	params["preview_length"] = fmt.Sprint(par.PreviewLength)
	params["extended"] = boolConverter(par.Extended)
	params["fields"] = par.Fields
	params["group_id"] = fmt.Sprint(par.GroupId)

	resp, err := m.vk.Request("messages.getById", params)
	if err != nil {
		return GetByIdResponse{}, err
	}

	var getByIdResp GetByIdResponse
	err = json.Unmarshal(resp, &getByIdResp)
	if err != nil {
		return GetByIdResponse{}, err
	}
	return getByIdResp, nil
}

// GetHistoryParams provides structure for
// parameters for get method.
// https://vk.com/dev/messages.getHistory
type GetHistoryParams struct {
	Offset         int
	Count          int
	UserId         int
	PeerId         int
	StartMessageId int
	Rev            bool
	Extended       bool
	Fields         string
	GroupId        uint
}

// GetHistoryResponse describes
// https://vk.com/dev/messages.getHistory
type GetHistoryResponse struct {
	Count int
	Items []MessageObject
}

// Returns message history for the
// specified user or group chat.
// https://vk.com/dev/messages.getHistory
func (m *Messages) GetHistory(par GetHistoryParams) (GetHistoryResponse, error) {
	params := make(map[string]string)

	params["offset"] = fmt.Sprint(par.Offset)
	params["count"] = fmt.Sprint(par.Count)
	params["user_id"] = fmt.Sprint(par.UserId)
	params["peer_id"] = fmt.Sprint(par.PeerId)
	//params["start_message_id"] = fmt.Sprint(par.StartMessageId)
	params["rev"] = boolConverter(par.Rev)
	params["extended"] = boolConverter(par.Extended)
	params["fields"] = par.Fields
	params["group_id"] = fmt.Sprint(par.GroupId)

	resp, err := m.vk.Request("messages.getHistory", params)
	if err != nil {
		return GetHistoryResponse{}, err
	}

	var getHistoryResp GetHistoryResponse
	err = json.Unmarshal(resp, &getHistoryResp)
	if err != nil {
		return GetHistoryResponse{}, err
	}
	return getHistoryResp, nil
}

// https://vk.com/dev/messages.getConversations
type GetConversationsParams struct {
	Offset         int
	Count          int
	Filter         TypeFilter
	Extended       bool
	StartMessageId int
	Fields         string
	GroupId        uint
	MajorSortId    uint
}

// https://vk.com/dev/messages.getConversations
type GetConversationsResponse struct {
	Count int `json:"count"`
	Items []struct {
		MessageObject      MessageObject      `json:"last_message"`
		ConversationObject ConversationObject `json:"conversation"`
	}
	UnreadCount int `json:"unread_count"`
	//Profiles    []UserObject  `json:"profiles"`
	//Groups      []GroupObject `json:"groups"`
}

// https://vk.com/dev/messages.getConversations
func (m *Messages) GetConversations(par GetConversationsParams) (GetConversationsResponse, error) {
	params := make(map[string]string)

	params["offset"] = fmt.Sprint(par.Offset)
	params["count"] = fmt.Sprint(par.Count)
	params["filter"] = fmt.Sprint(par.Filter)
	params["extended"] = boolConverter(par.Extended)
	params["fields"] = par.Fields
	params["group_id"] = fmt.Sprint(par.GroupId)
	params["major_sort_id"] = fmt.Sprint(par.MajorSortId)

	resp, err := m.vk.Request("messages.getConversations", params)
	if err != nil {
		return GetConversationsResponse{}, err
	}

	var getConversationsResp GetConversationsResponse
	err = json.Unmarshal(resp, &getConversationsResp)
	if err != nil {
		return GetConversationsResponse{}, err
	}
	return getConversationsResp, nil
}

type Keyboard struct {
	OneTime bool               `json:"one_time"`
	Buttons [][]KeyboardButton `json:"buttons"`
}

type KeyboardButton struct {
	Action KeyboardButtonAction `json:"action"`
	Color  KeyboardButtonColor  `json:"color"`
}

type KeyboardButtonAction struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}

type KeyboardButtonColor string

const KeyboardButtonColorPrimary KeyboardButtonColor = "primary"
const KeyboardButtonColorDefault KeyboardButtonColor = "default"
const KeyboardButtonColorNegative KeyboardButtonColor = "negative"
const KeyboardButtonColorPositive KeyboardButtonColor = "positive"

type keyboardBuilder struct {
	oneTime bool
	buttons [][]KeyboardButton
}

func NewKeyboardBuilder() *keyboardBuilder {
	return &keyboardBuilder{buttons: make([][]KeyboardButton, 0)}
}

func (builder *keyboardBuilder) NewRow(buttons ...KeyboardButton) {
	builder.buttons = append(builder.buttons, buttons)
}

func (builder *keyboardBuilder) SetOneTime(oneTime bool) {
	builder.oneTime = oneTime
}

func (builder *keyboardBuilder) Build() *Keyboard {
	return &Keyboard{
		OneTime: builder.oneTime,
		Buttons: builder.buttons,
	}
}

type TypeActivity string

const TypeActivityTyping TypeActivity = "typing"
const TypeActivityAudio TypeActivity = "audiomessage"

type TypeFilter string

const (
	TypeFilterAll        TypeFilter = "all"
	TypeFilterUnread     TypeFilter = "unread"
	TypeFilterImportant  TypeFilter = "important"
	TypeFilterUnanswered TypeFilter = "unanswered"
)
