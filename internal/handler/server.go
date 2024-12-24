package handler

import (
    "encoding/json"
    "net/http"
    "go-tunnel-panel/internal/model"
    "go-tunnel-panel/internal/service"
)

type ServerHandler struct {
    tunnel  *service.Tunnel
    monitor *service.Monitor
}

func NewServerHandler(t *service.Tunnel, m *service.Monitor) *ServerHandler {
    return &ServerHandler{
        tunnel:  t,
        monitor: m,
    }
}

func (h *ServerHandler) AddServer(w http.ResponseWriter, r *http.Request) {
    var server model.Server
    if err := json.NewDecoder(r.Body).Decode(&server); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "status": "success",
        "id":     server.ID,
    })
}

func (h *ServerHandler) GetServerStats(w http.ResponseWriter, r *http.Request) {
    serverID := r.URL.Query().Get("id")
    stats := h.monitor.GetStats(serverID)
    json.NewEncoder(w).Encode(stats)
} 