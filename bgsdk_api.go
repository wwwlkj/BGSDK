package BGSDK

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	credential "github.com/bytedance/douyin-openapi-credential-go/client"
	openApiSdkClient "github.com/bytedance/douyin-openapi-sdk-go/client"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"sort"
	"strings"
)

// getAccessToken 获取access_token
func getAccessToken(appid, appSecret string) (string, error) {
	url := "https://developer.toutiao.com/api/apps/v2/token"
	requestData := map[string]string{
		"appid":      appid,
		"secret":     appSecret,
		"grant_type": "client_credential",
	}

	requestBody, _ := sonic.Marshal(requestData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result AccessTokenRes
	err = sonic.Unmarshal(data, &result)
	if err != nil && result.ErrNo != 0 {
		return "", err
	}
	if result.ErrNo != 0 {
		return "", errors.New(result.ErrTips)
	}
	return result.Data.AccessToken, nil
}

// GetRoomInfo 获取直播间信息 也就是根据Token获取直播ID
func getRoomInfo(liveToken, accessToken string) GetRoomInfoRes {
	client := &http.Client{}
	var data = strings.NewReader(`{"token":"` + liveToken + `"}`)
	req, err := http.NewRequest("POST", "https://webcast.bytedance.com/api/webcastmate/info", data)
	if err != nil {
		return GetRoomInfoRes{}
	}
	req.Header.Set("X-Token", accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return GetRoomInfoRes{}
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetRoomInfoRes{}
	}
	var roomInfo GetRoomInfoRes
	err = sonic.Unmarshal(bodyText, &roomInfo)
	if err != nil {
		return GetRoomInfoRes{}
	}
	return roomInfo
}

// StartRoom 启动直播间推送任务
func startRoom(appid, roomId, msgType, accessToken string) (StartRoomRes, error) {
	url := "https://webcast.bytedance.com/api/live_data/task/start"
	requestData := map[string]string{
		"roomid":   roomId,
		"appid":    appid,
		"msg_type": msgType,
	}
	payload, _ := sonic.Marshal(requestData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StartRoomRes{}, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var result StartRoomRes
	err = sonic.Unmarshal(data, &result)
	if err != nil {
		return StartRoomRes{}, err
	}
	fmt.Println(string(data))
	return result, nil
}

// StopRoom 停止直播间推送任务
func stopRoom(appid, roomId, msgType, accessToken string) (StopRoomRes, error) {
	url := "https://webcast.bytedance.com/api/live_data/task/stop"
	requestData := map[string]string{
		"roomid":   roomId,
		"appid":    appid,
		"msg_type": msgType,
	}
	payload, _ := sonic.Marshal(requestData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StopRoomRes{}, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var result StopRoomRes
	err = sonic.Unmarshal(data, &result)
	if err != nil {
		return StopRoomRes{}, err
	}
	return result, nil
}

// GetTask 查询直播间任务状态
func getTask(appid, roomId, msgType, accessToken string) (GetTaskRes, error) {
	url := "https://webcast.bytedance.com/api/live_data/task/get?roomid=" + roomId + "&appid=" + appid + "&msg_type=" + msgType
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetTaskRes{}, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var result GetTaskRes
	err = sonic.Unmarshal(data, &result)
	if err != nil {
		return GetTaskRes{}, err
	}
	return result, nil
}

// Signature 签名
func signature(header map[string]string, bodyStr, msgSecret string) string {
	keyList := make([]string, 0, 4)
	for key, _ := range header {
		keyList = append(keyList, key)
	}
	sort.Slice(keyList, func(i, j int) bool {
		return keyList[i] < keyList[j]
	})
	kvList := make([]string, 0, 4)
	for _, key := range keyList {
		kvList = append(kvList, key+"="+header[key])
	}
	urlParams := strings.Join(kvList, "&")
	rawData := urlParams + bodyStr + msgSecret
	md5Result := md5.Sum([]byte(rawData))
	return base64.StdEncoding.EncodeToString(md5Result[:])
}

func roundSyncStatus(appid, appSecret, anchorOpenId, roomid, accessToken string, status int, startTime, endTime int64) bool {
	// 初始化SDK client
	opt := new(credential.Config).
		SetClientKey(appid).       // 改成自己的app_id
		SetClientSecret(appSecret) // 改成自己的secret
	sdkClient, err := openApiSdkClient.NewClient(opt)
	if err != nil {
		return false
	}
	sdkRequest := &openApiSdkClient.RoundSyncStatusRequest{}
	sdkRequest.SetAnchorOpenId(anchorOpenId)
	sdkRequest.SetAppId(appid)
	sdkRequest.SetEndTime(endTime)
	sdkRequest.SetRoomId(roomid)
	sdkRequest.SetRoundId(19516516)
	sdkRequest.SetStartTime(startTime)
	sdkRequest.SetStatus(status)
	sdkRequest.SetAccessToken(accessToken)
	// sdk调用
	_, err = sdkClient.RoundSyncStatus(sdkRequest)
	if err != nil {
		return false
	}
	return true
}

// HandlerChatMessage 解析评论消息
func handlerChatMessage(msg []byte) []DmMsg {
	var chatMessage []ChatMessage
	var res []DmMsg
	err := sonic.Unmarshal(msg, &chatMessage)
	if err != nil {
		return []DmMsg{}
	}
	for _, v := range chatMessage {
		msg2 := DmMsg{
			Type:  2,
			Uid:   v.SecOpenid,
			Name:  v.Nickname,
			Url:   v.AvatarUrl,
			Msg:   v.Content,
			Gift:  "",
			Num:   0,
			Price: 0,
		}
		res = append(res, msg2)
	}
	return res
}

// HandlerGiftMessage 解析礼物消息
func handlerGiftMessage(msg []byte) []DmMsg {
	var giftMessage []GiftMessage
	var res []DmMsg
	err := sonic.Unmarshal(msg, &giftMessage)
	if err != nil {
		return []DmMsg{}
	}
	for _, v := range giftMessage {

		msg2 := DmMsg{
			Type:  4,
			Uid:   v.SecOpenid,
			Name:  v.Nickname,
			Url:   v.AvatarUrl,
			Msg:   "送礼",
			Gift:  v.SecGiftId,
			Num:   int(v.GiftNum),
			Price: int(v.GiftValue),
		}
		res = append(res, msg2)
	}

	return res
}

// HandlerLikeMessage 解析点赞消息
func handlerLikeMessage(msg []byte) []DmMsg {
	var likeMessage []LikeMessage
	var res []DmMsg
	err := sonic.Unmarshal(msg, &likeMessage)
	if err != nil {
		return []DmMsg{}
	}
	for _, v := range likeMessage {
		msg2 := DmMsg{
			Type:  3,
			Uid:   v.SecOpenid,
			Name:  v.Nickname,
			Url:   v.AvatarUrl,
			Msg:   "点赞",
			Gift:  "",
			Num:   int(v.LikeNum),
			Price: 0,
		}
		res = append(res, msg2)
	}
	return res
}

// HandlerFansMessage 解析粉丝团消息
func handlerFansMessage(msg []byte) []DmMsg {
	var fansMessage []FansMessage
	var res []DmMsg
	err := sonic.Unmarshal(msg, &fansMessage)
	if err != nil {
		return []DmMsg{}
	}
	for _, v := range fansMessage {
		msg2 := DmMsg{
			Type:  5,
			Uid:   v.SecOpenid,
			Name:  v.Nickname,
			Url:   v.AvatarUrl,
			Msg:   "粉丝团",
			Gift:  "",
			Num:   0,
			Price: 0,
		}
		res = append(res, msg2)
	}
	return res
}
