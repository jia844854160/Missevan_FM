package main

import (
	"fmt"

	"Sh_Bot"
	"Sh_Bot/util"
)

func main() {
	fmt.Print("[Sh_Bot] 请输入直播间 ID: ")
	var rid int64
	_, err := fmt.Scanln(&rid)
	if err != nil {
		util.Error("直播间 ID 不正确", nil)
		return
	}

	ch := make(chan Sh_Bot.FmMessage)

	go Sh_Bot.Connect(ch, rid)

	for msg := range ch {
		Sh_Bot.HandleFmMessage(msg)
	}
}
