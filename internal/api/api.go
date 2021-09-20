package api

import (
	"context"
	"fmt"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/handler"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/server"
	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"time"
)

func Run() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

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

	e := echo.New()
	e.Debug = opts.Debug

	if err := handler.InitRoutes(e); err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(server.Opts{
		Addr:         fmt.Sprintf(":%d", opts.API.Port),
		ReadTimeout:  opts.API.ReadTimeout,
		WriteTimeout: opts.API.WriteTimeout,
	})

	// start server
	go func() {
		if err := e.StartServer(srv); err != nil {
			e.Logger.Infof("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Infof("error occurred while shutting down: %v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds left")
	}
	log.Println("server successfully stopped")
}
