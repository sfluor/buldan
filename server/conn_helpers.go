package main

import (
	"net/http"

	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

func errOut(lg *zap.Logger, w http.ResponseWriter, msg string, code int) {
	lg.Error("An error occurred", zap.Int("code", code), zap.String("message", msg))
	http.Error(w, msg, code)
}

func errOutWS(lg *zap.Logger, c *websocket.Conn, msg string, code websocket.StatusCode) {
	lg.Error("A websocket error occurred", zap.Any("code", code), zap.String("message", msg))
	if err := c.Close(code, msg); err != nil {
		lg.Error("Failed to close websocket following error", zap.Error(err), zap.Any("code", code), zap.String("message", msg))
	}
}
