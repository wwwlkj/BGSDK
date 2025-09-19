package BGSDK

import "time"

type ByteGameSDK struct {
	appID       string // 应用ID
	appSecret   string // 应用密钥
	accessToken string // access_token
	msgSecret   string // 消息密钥
}

func NewByteGameSDK(appID string, appSecret string, msgSecret string) *ByteGameSDK {
	return &ByteGameSDK{
		appID:     appID,
		appSecret: appSecret,
		msgSecret: msgSecret,
	}
}

func (sdk *ByteGameSDK) AppID() string {
	return sdk.appID
}

func (sdk *ByteGameSDK) AppSecret() string {
	return sdk.appSecret
}

func (sdk *ByteGameSDK) MsgSecret() string {
	return sdk.msgSecret
}

func (sdk *ByteGameSDK) AccessToken() string {
	return sdk.accessToken
}

func (sdk *ByteGameSDK) SetAccessToken(token string) {
	sdk.accessToken = token
}

// GetAccessTokenStart 获取access_token
func (sdk *ByteGameSDK) GetAccessTokenStart() {
	go func() {
		for {
			token, err := getAccessToken(sdk.appID, sdk.appSecret)
			if err == nil {
				sdk.accessToken = token
				time.Sleep(4 * time.Second)
			}
		}
	}()
}

// GetRoomInfo 获取直播间信息
func (sdk *ByteGameSDK) GetRoomInfo(liveToken string) GetRoomInfoRes {
	return getRoomInfo(liveToken, sdk.accessToken)
}

// StartRoom 启动直播间推送任务
func (sdk *ByteGameSDK) StartRoom(roomId, msgType string) (StartRoomRes, error) {
	return startRoom(sdk.appID, roomId, msgType, sdk.accessToken)
}

// StopRoom 停止直播间推送任务
func (sdk *ByteGameSDK) StopRoom(roomId, msgType string) (StopRoomRes, error) {
	return stopRoom(sdk.appID, roomId, msgType, sdk.accessToken)
}

// GetTask 查询直播间任务状态
func (sdk *ByteGameSDK) GetTask(roomId, msgType string) (GetTaskRes, error) {
	return getTask(sdk.appID, roomId, msgType, sdk.accessToken)
}

// RoundSyncStatus 同步游戏状态
func (sdk *ByteGameSDK) RoundSyncStatus(anchorOpenId, roomid string, status int, startTime, endTime int64) bool {
	return roundSyncStatus(sdk.appID, sdk.appSecret, anchorOpenId, roomid, sdk.accessToken, status, startTime, endTime)
}

// Signature 消息签名
func (sdk *ByteGameSDK) Signature(header map[string]string, bodyStr string) string {
	return signature(header, bodyStr, sdk.msgSecret)
}

// HandlerChatMessage 解析评论消息
func (sdk *ByteGameSDK) HandlerChatMessage(msg []byte) []DmMsg {
	return handlerChatMessage(msg)
}

// HandlerGiftMessage 解析礼物消息
func (sdk *ByteGameSDK) HandlerGiftMessage(msg []byte) []DmMsg {
	return handlerGiftMessage(msg)
}

// HandlerLikeMessage 解析点赞消息
func (sdk *ByteGameSDK) HandlerLikeMessage(msg []byte) []DmMsg {
	return handlerLikeMessage(msg)
}

// HandlerFansMessage 解析粉丝团消息
func (sdk *ByteGameSDK) HandlerFansMessage(msg []byte) []DmMsg {
	return handlerFansMessage(msg)
}
