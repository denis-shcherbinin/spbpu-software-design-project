package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jessevdk/go-flags"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/handler"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/postgres"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/server"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

func Run() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var opts Opts
	parser := flags.NewParser(&opts, flags.HelpFlag)
	if _, err := parser.Parse(); err != nil {
		fmt.Println(err)
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	postgresDB, err := postgres.NewPostgresDB(postgres.Opts{
		Host:     opts.DB.Host,
		User:     opts.DB.User,
		Password: opts.DB.Password,
		Name:     opts.DB.Name,
		Port:     opts.DB.Port,
	})
	if err != nil {
		log.Fatalf("error occured while connecting to PostgresDB: %v\n", err)
	}
	log.Println("PostgresDB connected!")

	repo := repository.NewRepository(postgresDB)

	passwordHasher := hasher.NewSHA1Hasher(opts.Auth.PasswordSalt)

	services := service.NewService(repo, passwordHasher)

	handlers := handler.NewHandler(services)
	e := handlers.Init(handler.InitOpts{
		Debug: opts.Debug,
	})

	srv := server.NewServer(server.Opts{
		Addr:         fmt.Sprintf(":%d", opts.API.Port),
		ReadTimeout:  opts.API.ReadTimeout,
		WriteTimeout: opts.API.WriteTimeout,
	})

	// start server
	go func() {
		if err := e.StartServer(srv); err != nil {
			e.Logger.Infof("shutting down the server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		err = postgresDB.Close()
		if err != nil {
			log.Printf("error occurred while disconnecting postgresDB: %v\n", err)
		} else {
			log.Println("PostgresDB disconnected")
		}

		cancel()
	}()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Infof("error occurred while shutting down: %v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds left")
	}
	log.Println("server successfully stopped")
}
