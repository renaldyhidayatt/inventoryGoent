package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/renaldyhidayatt/inventorygoent/internal/handler"
	"github.com/renaldyhidayatt/inventorygoent/internal/mapper"
	"github.com/renaldyhidayatt/inventorygoent/internal/repository"
	"github.com/renaldyhidayatt/inventorygoent/internal/service"
	"github.com/renaldyhidayatt/inventorygoent/pkg/auth"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres"
	"github.com/renaldyhidayatt/inventorygoent/pkg/dotenv"
	"github.com/renaldyhidayatt/inventorygoent/pkg/hash"
	"github.com/renaldyhidayatt/inventorygoent/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Error("Failed to initialize logger:", zap.Error(err))
	}

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	}

	err = dotenv.Viper()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	db, err := postgres.NewClient(ctx)

	if err != nil {
		log.Fatal("Failed to connect to database:", zap.Error(err))
		panic(err)
	}

	hashing := hash.NewHashingPassword()

	repository := repository.NewRepositories(db, ctx)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Fatal("Failed to initialize JWT manager:", zap.Error(err))

		panic(err)
	}

	mapper := mapper.NewMapper()

	service := service.NewServices(service.Deps{
		Repository: repository,
		Logger:     log,
		Hash:       hashing,
		Token:      token,
		Mapper:     mapper,
	})

	myhandler := handler.NewHandler(service)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		WriteTimeout: time.Duration(viper.GetInt("WRITE_TIME_OUT")) * time.Second * 10,
		ReadTimeout:  time.Duration(viper.GetInt("READ_TIME_OUT")) * time.Second * 10,

		IdleTimeout: time.Second * 60,
		Handler:     myhandler.Init(),
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Info("Connected to port: " + viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)

}
