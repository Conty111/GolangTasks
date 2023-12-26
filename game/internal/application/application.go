//package application
//
//import (
//	"context"
//	"example.com/m/internal/config"
//	"example.com/m/pkg/life"
//	"fmt"
//	"time"
//)
//
//type Application struct {
//	Cfg *config.Config
//}
//
//func New(config *config.Config) *Application {
//	return &Application{
//		Cfg: config,
//	}
//}
//
//func (a *Application) Run(ctx context.Context) error {
//	// Объект для хранения текущего состояния сетки
//	currentWorld, _ := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
//	// Объект для хранения очередного состояния сетки
//	nextWorld, _ := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
//	// Заполняем сетку на проценты
//	currentWorld.RandInit(a.Cfg.Percents)
//	for {
//		// Здесь мы можем записывать текущее состояние  — например, в очередь сообщений. Для нашего примера просто выводим на экран
//		fmt.Println(currentWorld.String())
//		life.NextState(currentWorld, nextWorld)
//		currentWorld = nextWorld
//		// Проверяем контекст
//		select {
//		case <-ctx.Done():
//			return ctx.Err() // Возвращаем причину завершения
//		default: // По умолчанию делаем паузу
//			time.Sleep(300 * time.Millisecond)
//			break
//		}
//		// Очищаем экран
//		fmt.Print("\033[H\033[2J")
//	}
//}

package application

import (
	"context"
	"example.com/m/http/server"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) int {
	// Создание логгера с настройками для production
	logger := setupLogger()

	shutDownFunc, err := server.Run(ctx, logger, a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		logger.Error(err.Error())

		return 1 // вернем код для регистрация ошибки системой
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-c
	cancel()
	//  завершим работу сервера
	shutDownFunc(ctx)

	return 0

}

// настройки логгера
func setupLogger() *zap.Logger {
	// Настройка конфигурации логгера
	config := zap.NewProductionConfig()

	// Уровень логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// Настройка логгера с конфигом
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
	}

	return logger
}
