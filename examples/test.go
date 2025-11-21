package main

import (
	"fmt"
	BGSDK "github.com/wwwlkj/bgsdk"
	"strconv"
	"time"
)

// SKkwUvTSSUkVuvnk5q9m/Z+KBJZJFKq9adc1IIZuSI+E9xNSIYgEb1jP/cc+4nJJ3r8DM6vUKH2bnU6DrgRKxe6GXQSnWv6VuP5r2DEckCmR63iG+vDPZQrFER5QhH1K0F1vbVJ89ZhH8e9w
// "tt8c44faf9134c0bd010", "8220b08f73cf49b8380d098f8fda2b284ba914e3"
func main() {

	accessToken, err := BGSDK.GetAccessToken("", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accessToken)
	roomInfo := BGSDK.GetRoomInfo("", accessToken)
	fmt.Println(roomInfo)
	roomRes := BGSDK.RoundSyncStatus("",
		"",
		"",
		strconv.Itoa(7551611853076777769),
		accessToken,
		1,
		time.Now().Unix(),
		time.Now().Unix()+10000,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RoundSyncStatus:", roomRes)
}
