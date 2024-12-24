package main

import (
    "log"
    "net/http"
    "go-tunnel-panel/internal/handler"
    "go-tunnel-panel/internal/service"
)

func main() {
    tunnel := service.NewTunnel()
    monitor := service.NewMonitor()
    handler := handler.NewServerHandler(tunnel, monitor)

    // 设置路由
    http.HandleFunc("/api/server/add", handler.AddServer)
    http.HandleFunc("/api/server/stats", handler.GetServerStats)
    http.HandleFunc("/ws/tunnel", tunnel.HandleConnection)

    log.Printf("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
} 