package blog

import (
	"context"
	"log"
)

// Server should be exported
type Server struct {
}

// SayHello should be exported
func (s *Server) SayHello(ctx context.Context, blog *Blog) (*Blog, error) {
	log.Println("Message from client", blog)
	return &Blog{Title: "Berita dari server", Body: "Sudah terkoneksi bro, ini pesan dari server"}, nil
}
