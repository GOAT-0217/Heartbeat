package test

import (
	"context"
	"heartbeat/utils"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// TestPublishMessage 测试发布消息到redis
func TestPublishMessage(t *testing.T) {
	msg := "当前时间: " + time.Now().Format("15:04:05")
	err := utils.Publish(ctx, utils.PublishKey, msg)
	if err != nil {
		t.Errorf("发布消息失败: %v", err)
	}
}
