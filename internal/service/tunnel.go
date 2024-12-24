package service

import (
    "github.com/gorilla/websocket"
    "net/http"
    "go-tunnel-panel/internal/model"
)

type Tunnel struct {
    upgrader websocket.Upgrader
    servers  map[string]*model.Server
}

func NewTunnel() *Tunnel {
    return &Tunnel{
        upgrader: websocket.Upgrader{
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
        },
        servers: make(map[string]*model.Server),
    }
}

func (t *Tunnel) HandleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := t.upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    go t.handleTunnelData(conn)
}

func (t *Tunnel) handleTunnelData(conn *websocket.Conn) {
    for {
        messageType, data, err := conn.ReadMessage()
        if err != nil {
            return
        }

        // 处理数据
        processedData := data // 这里需要实现数据处理逻辑

        err = conn.WriteMessage(messageType, processedData)
        if err != nil {
            return
        }
    }
} 