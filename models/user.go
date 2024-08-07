package models

import "github.com/gorilla/websocket"

// User 用户信息
type User struct {
	ID       string
	Conn     *websocket.Conn
	Send     chan []byte
	Backpack []string
}

// Health 用户健康信息
type Health struct {
	ID       string
	IsHealth bool
}

// Message 通用信息
type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}
