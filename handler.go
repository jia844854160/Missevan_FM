package Sh_Bot

import (
	"fmt"

	"Sh_Bot/util"
	"github.com/fatih/color"
)

func HandleFmMessage(msg FmMessage) {
	if msg.Type == TypeMember { // 只处理成员相关的消息
		handleMember(msg)
	}
}

// handleMember 处理类型为 Member 的消息
func handleMember(msg FmMessage) {
	if msg.Event == EventJoinQueue {
		handleMemberJoinQueue(msg)
	}
}

func handleMemberJoinQueue(msg FmMessage) {
	for _, v := range msg.Queue {
		name := v.Username
		if name == "" {
			util.Print("匿名用户进入直播间", color.FgHiCyan)
		} else {
			util.Print(fmt.Sprintf("用户 @%s 进入直播间", name), color.FgHiCyan)
		}
	}
}
