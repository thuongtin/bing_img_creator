package test

import (
	"context"
	"testing"

	bing "github.com/thuongtin/bing_img_creator"
)

func TestSMS(t *testing.T) {
	bingChat, err := bing.NewChat("new bing cookie ")
	if err != nil {
		t.Error(err)
		return
	}
	ans, err := bingChat.Chat(context.TODO(), "你是chatGPT吗？")
	if err != nil {
		t.Error(err)
		return
	}
	println(ans)

	ans, err = bingChat.Chat(context.TODO(), "你能做什么？")
	if err != nil {
		t.Error(err)
		return
	}
	println(ans)

}
