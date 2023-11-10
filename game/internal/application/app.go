package application

import (
	"context"
	"example.com/m/internal/config"
	"example.com/m/pkg/life"
	"fmt"
	"time"
)

type Application struct {
	Cfg *config.Config
}

func New(config *config.Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) error {
	// Объект для хранения текущего состояния сетки
	currentWorld, _ := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	// Объект для хранения очередного состояния сетки
	nextWorld, _ := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	// Заполняем сетку на проценты
	currentWorld.RandInit(a.Cfg.Percents)
	for {
		// Здесь мы можем записывать текущее состояние  — например, в очередь сообщений. Для нашего примера просто выводим на экран
		fmt.Println(currentWorld.String())
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		// Проверяем контекст
		select {
		case <-ctx.Done():
			return ctx.Err() // Возвращаем причину завершения
		default: // По умолчанию делаем паузу
			time.Sleep(300 * time.Millisecond)
			break
		}
		// Очищаем экран
		fmt.Print("\033[H\033[2J")
	}
}
