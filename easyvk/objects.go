package easyvk

// An UserObject contains information about user.
// https://vk.com/dev/objects/user
type UserObject struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Sex        int    `json:"sex"`
	Nickname   string `json:"nickname"`
	MaidenName string `json:"maiden_name"`
	Domain     string `json:"domain"`
	ScreenName string `json:"screen_name"`
	Bdate      string `json:"bdate"`
	City       struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Country struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	Photo50                string `json:"photo_50"`
	Photo100               string `json:"photo_100"`
	Photo200               string `json:"photo_200"`
	PhotoMax               string `json:"photo_max"`
	Photo200Orig           string `json:"photo_200_orig"`
	Photo400Orig           string `json:"photo_400_orig"`
	PhotoMaxOrig           string `json:"photo_max_orig"`
	PhotoID                string `json:"photo_id"`
	HasPhoto               int    `json:"has_photo"`
	HasMobile              int    `json:"has_mobile"`
	IsFriend               int    `json:"is_friend"`
	FriendStatus           int    `json:"friend_status"`
	Online                 int    `json:"online"`
	WallComments           int    `json:"wall_comments"`
	CanPost                int    `json:"can_post"`
	CanSeeAllPosts         int    `json:"can_see_all_posts"`
	CanSeeAudio            int    `json:"can_see_audio"`
	CanWritePrivateMessage int    `json:"can_write_private_message"`
	CanSendFriendRequest   int    `json:"can_send_friend_request"`
	MobilePhone            string `json:"mobile_phone"`
	HomePhone              string `json:"home_phone"`
	Site                   string `json:"site"`
	Status                 string `json:"status"`
	LastSeen               struct {
		Time     int `json:"time"`
		Platform int `json:"platform"`
	} `json:"last_seen"`
	CropPhoto struct {
		Photo struct {
			ID       int    `json:"id"`
			AlbumID  int    `json:"album_id"`
			OwnerID  int    `json:"owner_id"`
			Photo75  string `json:"photo_75"`
			Photo130 string `json:"photo_130"`
			Photo604 string `json:"photo_604"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
			Text     string `json:"text"`
			Date     int    `json:"date"`
			PostID   int    `json:"post_id"`
		} `json:"photo"`
		Crop struct {
			X  float64 `json:"x"`
			Y  float64 `json:"y"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"crop"`
		Rect struct {
			X  int `json:"x"`
			Y  int `json:"y"`
			X2 int `json:"x2"`
			Y2 int `json:"y2"`
		} `json:"rect"`
	} `json:"crop_photo"`
	Verified         int `json:"verified"`
	FollowersCount   int `json:"followers_count"`
	Blacklisted      int `json:"blacklisted"`
	BlacklistedByMe  int `json:"blacklisted_by_me"`
	IsFavorite       int `json:"is_favorite"`
	IsHiddenFromFeed int `json:"is_hidden_from_feed"`
	CommonCount      int `json:"common_count"`
	Career           []struct {
		GroupID   int    `json:"group_id"`
		Company   string `json:"company"`
		CountryId int    `json:"country_id"`
		CityId    int    `json:"city_id"`
		CityName  string `json:"city_name"`
		From      int    `json:"from"`
		Until     int    `json:"until"`
		Position  string `json:"position"`
	} `json:"career"`
	Military       []interface{} `json:"military"`
	University     int           `json:"university"`
	UniversityName string        `json:"university_name"`
	Faculty        int           `json:"faculty"`
	FacultyName    string        `json:"faculty_name"`
	Graduation     int           `json:"graduation"`
	HomeTown       string        `json:"home_town"`
	Relation       int           `json:"relation"`
	Personal       struct {
		Religion   string `json:"religion"`
		InspiredBy string `json:"inspired_by"`
		PeopleMain int    `json:"people_main"`
		LifeMain   int    `json:"life_main"`
		Smoking    int    `json:"smoking"`
		Alcohol    int    `json:"alcohol"`
	} `json:"personal"`
	Interests    string `json:"interests"`
	Music        string `json:"music"`
	Activities   string `json:"activities"`
	Movies       string `json:"movies"`
	Tv           string `json:"tv"`
	Books        string `json:"books"`
	Games        string `json:"games"`
	Universities []struct {
		ID              int    `json:"id"`
		Country         int    `json:"country"`
		City            int    `json:"city"`
		Name            string `json:"name"`
		Faculty         int    `json:"faculty"`
		FacultyName     string `json:"faculty_name"`
		Chair           int    `json:"chair"`
		ChairName       string `json:"chair_name"`
		Graduation      int    `json:"graduation"`
		EducationForm   string `json:"education_form"`
		EducationStatus string `json:"education_status"`
	} `json:"universities"`
	Schools     []interface{} `json:"schools"`
	About       string        `json:"about"`
	Relatives   []interface{} `json:"relatives"`
	Quotes      string        `json:"quotes"`
	Deactivated string        `json:"deactivated"`
}

// A PhotoObject contains information about photo.
// https://vk.com/dev/objects/photo
type PhotoObject struct {
	ID      int `json:"id"`
	AlbumID int `json:"album_id"`
	OwnerID int `json:"owner_id"`
	UserID  int `json:"user_id"`
	Sizes   []struct {
		Url    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Type   string `json:"type"`
	} `json:"sizes"`
	Text      string `json:"text"`
	Date      int    `json:"date"`
	PostID    int    `json:"post_id"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Photo75   string `json:"photo_75"`
	Photo130  string `json:"photo_130"`
	Photo604  string `json:"photo_604"`
	Photo807  string `json:"photo_807"`
	Photo1280 string `json:"photo_1280"`
	Photo2560 string `json:"photo_2560"`
	Likes     struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
	Comments struct {
		Count int `json:"count"`
	} `json:"comments"`
	CanComment int `json:"can_comment"`
	Tags       struct {
		Count int `json:"count"`
	} `json:"tags"`
}

// A VideoObject contains information about video.
// https://vk.com/dev/objects/video
type VideoObject struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Date        int    `json:"date"`
	Comments    int    `json:"comments"`
	Views       int    `json:"views"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Photo130    string `json:"photo_130"`
	Photo320    string `json:"photo_320"`
	Photo800    string `json:"photo_800"`
	AddingDate  int    `json:"adding_date"`
	Files       struct {
		Mp4240 string `json:"mp4_240"`
		Mp4360 string `json:"mp4_360"`
		Mp4480 string `json:"mp4_480"`
		Mp4720 string `json:"mp4_720"`
	} `json:"files"`
	Player     string `json:"player"`
	CanAdd     int    `json:"can_add"`
	CanComment int    `json:"can_comment"`
	CanRepost  int    `json:"can_repost"`
	Likes      struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Repeat int `json:"repeat"`
}

// A GroupObject contains information about group.
// https://vk.com/dev/objects/group
type GroupObject struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Type         string `json:"type"`
	IsAdmin      int    `json:"is_admin"`
	IsMember     int    `json:"is_member"`
	IsAdvertiser int    `json:"is_advertiser"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
}

// A GroupObject contains information about city.
// https://vk.com/dev/objects/group
type CityObject struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Important bool   `json:"important"`
}

// A AttachmentMObject contains information about attachments_m.
// https://vk.com/dev/objects/attachments_m
type AttachmentMObject struct {
	Type  TypeAttachmentM `json:"type"`
	Photo PhotoObject     `json:"photo"`
}

type TypeAttachmentM string

const TypeAttachmentMPhoto TypeAttachmentM = "photo"
const TypeAttachmentMVideo TypeAttachmentM = "video"
const TypeAttachmentMAudio TypeAttachmentM = "audio"
const TypeAttachmentMDoc TypeAttachmentM = "doc"
const TypeAttachmentMLink TypeAttachmentM = "link"
const TypeAttachmentMMarket TypeAttachmentM = "market"
const TypeAttachmentMMarketAlbum TypeAttachmentM = "market_album"
const TypeAttachmentMWall TypeAttachmentM = "wall"
const TypeAttachmentMWallReply TypeAttachmentM = "wall_reply"
const TypeAttachmentMSticker TypeAttachmentM = "sticker"
const TypeAttachmentMGift TypeAttachmentM = "gift"

// A MessageObject contains information about message.
// https://vk.com/dev/objects/message
type MessageObject struct {
	ID                    int                 `json:"id"`
	Date                  int                 `json:"date"`
	PeerID                int                 `json:"peer_id"`
	FromID                int                 `json:"from_id"`
	Text                  string              `json:"text"`
	RandomID              int64               `json:"random_id"`
	Ref                   string              `json:"ref"`
	RefSource             string              `json:"ref_source"`
	Attachments           []AttachmentMObject `json:"attachments"`
	Important             bool                `json:"important"`
	Payload               string              `json:"payload"`
	Out                   int                 `json:"out"`
	ConversationMessageID int                 `json:"conversation_message_id"`
	FwdMessages           []interface{}       `json:"fwd_messages"`
	IsHidden              bool                `json:"is_hidden"`
}

// A ConversationObject contains information about conversation.
// https://vk.com/dev/objects/conversation
type ConversationObject struct {
	Peer struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		LocalID int    `json:"local_id"`
	} `json:"peer"`
	InRead      int  `json:"in_read"`
	OutRead     int  `json:"out_read"`
	UnreadCount int  `json:"unread_count"`
	Important   bool `json:"important"`
	Unanswered  bool `json:"unanswered"`
	CanWrite    struct {
		Allowed bool `json:"allowed"`
		Reason  int  `json:"reason"`
	} `json:"can_write"`
	ChatSettings struct {
		MembersCount  int         `json:"members_count"`
		Title         int         `json:"title"`
		PinnedMessage interface{} `json:"pinned_message"`
		State         string      `json:"state"`
		Photo         struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
		ActiveIDs      []int `json:"active_ids"`
		IsGroupChannel bool  `json:"is_group_channel"`
	} `json:"chat_settings"`
}
