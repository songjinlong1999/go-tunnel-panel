let ws;
let chart;

function initWebSocket() {
    ws = new WebSocket('ws://' + window.location.host + '/ws/tunnel');
    ws.onmessage = function(event) {
        updateStats(JSON.parse(event.data));
    };
}

function addServer() {
    // TODO: 实现添加服务器的逻辑
}

function updateStats(data) {
    // TODO: 实现更新统计数据的逻辑
}

// 初始化
document.addEventListener('DOMContentLoaded', function() {
    initWebSocket();
}); 