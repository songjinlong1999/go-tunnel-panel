package model

type ServerType string

const (
    EntryServer  ServerType = "entry"
    ExitServer   ServerType = "exit"
)

type Server struct {
    ID          string     `json:"id"`
    Name        string     `json:"name"`
    Type        ServerType `json:"type"`
    IP          string     `json:"ip"`
    Port        int        `json:"port"`
    WSPort      int        `json:"ws_port"`
    Status      string     `json:"status"`
    UploadSpeed int64      `json:"upload_speed"`
    DownSpeed   int64      `json:"down_speed"`
    LastPing    int64      `json:"last_ping"`
}

type ServerStats struct {
    ServerID    string    `json:"server_id"`
    CPU         float64   `json:"cpu_usage"`
    Memory      float64   `json:"memory_usage"`
    NetworkIn   int64     `json:"network_in"`
    NetworkOut  int64     `json:"network_out"`
    Connections int       `json:"connections"`
    Timestamp   int64     `json:"timestamp"`
} 