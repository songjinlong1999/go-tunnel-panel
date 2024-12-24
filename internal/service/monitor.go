package service

import (
    "time"
    "go-tunnel-panel/internal/model"
)

type Monitor struct {
    stats map[string]*model.ServerStats
}

func NewMonitor() *Monitor {
    return &Monitor{
        stats: make(map[string]*model.ServerStats),
    }
}

func (m *Monitor) CollectStats(serverID string) {
    stats := &model.ServerStats{
        ServerID:    serverID,
        Timestamp:   time.Now().Unix(),
    }
    
    // TODO: 实现实际的状态收集逻辑
    
    m.stats[serverID] = stats
}

func (m *Monitor) GetStats(serverID string) *model.ServerStats {
    return m.stats[serverID]
} 