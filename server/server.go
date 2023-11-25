package server

import (
	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

type GeisterServer struct {
	geisterpb.UnimplementedGeisterServer
}

func NewGeisterServer() *GeisterServer {
	return &GeisterServer{}
}
