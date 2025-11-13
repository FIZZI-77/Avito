package main

import (
	"avito/pkg"
	"avito/src/core/handler"
	"avito/src/core/repository"
	"avito/src/core/service"
	"context"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	_ "runtime/debug"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := pkg.NewPostgresDB(pkg.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handl := handler.NewHandler(services)

	srv := new(Server)

	go func() {
		if err := srv.Run(os.Getenv("PORT_SERVER"), handl.InitRouters()); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	g, ctx := errgroup.WithContext(ctx)

	log.Printf("Server started on port %s", os.Getenv("PORT_SERVER"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err = srv.ShutDown(shutdownCtx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	if err = db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	if err = g.Wait(); err != nil {
		log.Printf("Error in background tasks: %v", err)
	}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
