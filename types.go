package BGSDK

// AccessTokenRes 获取access_token响应
type AccessTokenRes struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}

// GetRoomInfoRes 获取直播间信息响应
type GetRoomInfoRes struct {
	Data struct {
		Info struct {
			RoomId       int64  `json:"room_id"`
			AnchorOpenId string `json:"anchor_open_id"`
			AvatarUrl    string `json:"avatar_url"`
			NickName     string `json:"nick_name"`
		} `json:"info"`
	} `json:"data"`
}

// StartRoomRes 启动直播间推送任务响应
type StartRoomRes struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Logid  string `json:"logid"`
	Data   struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

// StopRoomRes 停止直播间推送任务响应
type StopRoomRes struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Logid  string `json:"logid"`
	Data   struct {
	} `json:"data"`
}

// GetTaskRes 查询任务状态响应
type GetTaskRes struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Logid  string `json:"logid"`
	Data   struct {
		Status int `json:"status"`
	} `json:"data"`
}
type DmMsg struct {
	//1来了 2发言 3点赞 4送礼 5关注主播
	Type  int    `json:"type"`
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Msg   string `json:"msg"`
	Gift  string `json:"gift"`
	Num   int    `json:"num"`
	Price int    `json:"price"`
}

// ChatMessage 评论消息
type ChatMessage struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	Content   string `json:"content"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}

// GiftMessage 礼物消息
//
//	type GiftMessage struct {
//		MsgId     string `json:"msg_id"`
//		SecOpenid string `json:"sec_openid"`
//		SecGiftId string `json:"sec_gift_id"`
//		GiftNum   int64  `json:"gift_num"`
//		GiftValue int64  `json:"gift_value"`
//		AvatarUrl string `json:"avatar_url"`
//		Nickname  string `json:"nickname"`
//		Timestamp int64  `json:"timestamp"`
//		Test      bool   `json:"test"`
//	}
//
// GiftMessage 礼物消息
type GiftMessage struct {
	MsgId             string `json:"msg_id"`
	SecOpenid         string `json:"sec_openid"`
	SecGiftId         string `json:"sec_gift_id"`
	GiftNum           int    `json:"gift_num"`
	GiftValue         int    `json:"gift_value"`
	AvatarUrl         string `json:"avatar_url"`
	Nickname          string `json:"nickname"`
	Timestamp         int    `json:"timestamp"`
	Test              bool   `json:"test"`
	AudienceSecOpenId string `json:"audience_sec_open_id"`
	SecMagicGiftId    string `json:"sec_magic_gift_id"` // 加密的幸运魔方礼物id（有该字段，代表的是由幸运魔方抽取到的礼物。sec_gift_id是幸运魔方抽取到的礼物）
}

// LikeMessage 点赞消息
type LikeMessage struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	LikeNum   int64  `json:"like_num"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}

// FansMessage 粉丝团消息
type FansMessage struct {
	MsgId              string `json:"msg_id"`
	SecOpenid          string `json:"sec_openid"`
	AvatarUrl          string `json:"avatar_url"`
	Nickname           string `json:"nickname"`
	Timestamp          int64  `json:"timestamp"`
	FansclubReasonType int64  `json:"fansclub_reason_type"`
	FansclubLevel      int64  `json:"fansclub_level"`
}
