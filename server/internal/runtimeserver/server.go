package runtimeserver

import (
	stream "github.com/hongsam14/boxer-remote-control/server/internal/proto"
)

type server struct {
	stream.UnimplementedStreamerServer
}

func (s *server) StreamFrames(_ *stream.Empty, dataStream stream.Streamer_StreamFramesServer) error {
	return nil
}
