// Разработайте микросервис, который позволяет создавать, получать, обновлять и удалять задачи в списке дел (todo list):
//
// Используйте язык Golang и фреймворк или библиотеку по вашему выбору.
//
// Примените принципы чистой архитектуры (Clean Architecture).
//
// Микросервис должен иметь RESTful API для взаимодействия с клиентом.
//
// Задачи в списке дел должны иметь следующие атрибуты: идентификатор, название, описание, статус (выполнена/не выполнена), при необходимости можете добавить дополнительные поля.
//
// Реализуйте функциональность для создания новой задачи, получения списка всех задач, обновления существующей задачи и удаления задачи.
//
// Для хранения данных можно использовать простую базу данных или хранить данные в памяти (в памяти будет достаточно).
//
// Не требуется разработка фронтенд-части, только бэкенд-часть (REST API).
//
// Ограничения:
// Время выполнения задания ограничено 4 часами.
// Не требуется добавлять аутентификацию или авторизацию в API.
// Достаточно простой и минималистичный дизайн API.
package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/vgekko/go-tasks-graphql/config"
	v1 "github.com/vgekko/go-tasks-graphql/internal/controller/http/v1"
	"github.com/vgekko/go-tasks-graphql/internal/repository"
	"github.com/vgekko/go-tasks-graphql/internal/usecase"
	"github.com/vgekko/go-tasks-graphql/pkg/httpserver"
	"github.com/vgekko/go-tasks-graphql/pkg/logger"
	"github.com/vgekko/go-tasks-graphql/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	//config
	cfg := config.Load()

	// logger
	log := logger.New(cfg.Logger.Level)

	// postgres
	db, err := postgres.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Error("failed to init postgresql: ", err.Error())
	}

	// repository
	repository := repository.NewRepository(db)

	// use case
	useCase := usecase.NewUseCase(repository)

	// handler
	engine := gin.New()
	v1.NewHandler(engine, useCase, log)

	// http
	httpServer := httpserver.New(engine, cfg.HTTP)

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app.Run: signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error("app.Run: notify: ", err.Error())
	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app.Run: shutdown: ", err.Error())
	}
}
