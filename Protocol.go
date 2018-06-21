package go_push

import "encoding/json"

// websocket的Message对象
type WSMessage struct {
	msgType int
	msgData []byte
}

// 业务消息的固定格式(type+data)
type BizMessage struct {
	Type string `json:"type"`	// type消息类型: PING, PONG, JOIN, LEAVE, PUSH
	Data json.RawMessage `json:"data"`	// data数据字段
}

// Data数据类型

// PUSH
type BizPushData struct {
	Items []json.RawMessage	`json:"items"`
}

// PING
type BizPingData struct {}

// PONG
type BizPongData struct {}

// JOIN
type BizJoinData struct {
	Room string `json:"room"`
}

// LEAVE
type BizLeaveData struct {
	Room string `json:"room"`
}

func BuildWSMessage(msgType int, msgData []byte) (wsMessage *WSMessage) {
	return &WSMessage{
		msgType: msgType,
		msgData: msgData,
	}
}

// 解析{"type": "PING", "data": {...}}的包
func DecodeBizMessage(buf []byte) (bizMessage *BizMessage, err error) {
	var (
		bizMsgObj BizMessage
	)

	if err = json.Unmarshal(buf, &bizMsgObj); err != nil {
		return
	}

	bizMessage = &bizMsgObj
	return
}