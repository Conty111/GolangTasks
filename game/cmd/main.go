// package main
//
// import (
//
//	"context"
//	"errors"
//	"example.com/m/internal/application"
//	"example.com/m/internal/config"
//	"fmt"
//	"log"
//	"os"
//	"strconv"
//
// )
//
//	func main() {
//		ctx := context.Background()
//		// Exit завершает программу с заданным кодом
//		os.Exit(mainWithExitCode(ctx))
//	}
//
//	func mainWithExitCode(ctx context.Context) int {
//		cfg, err := setupConfig()
//		if err != nil {
//			log.Println("Get config error", err)
//		}
//		app := application.New(cfg)
//		// Запускаем приложение
//		if err := app.Run(ctx); err != nil {
//			switch {
//			case errors.Is(err, context.Canceled):
//				log.Println("Processing cancelled.")
//			default:
//				log.Println("Application run error", err)
//			}
//			// Возвращаем значение, не равное нулю, чтобы обозначить ошибку
//			return 1
//		}
//		// Выход без ошибок
//		return 0
//	}
//
//	func setupConfig() (*config.Config, error) {
//		cfg := config.New()
//		if os.Args != nil {
//			args := make([]int, len(os.Args))
//			var err error
//			for i := 1; i < 4; i++ {
//				args[i-1], err = strconv.Atoi(os.Args[i])
//				if err != nil {
//					return cfg, err
//				}
//				if args[i-1] < 1 {
//					return cfg, fmt.Errorf("Not positive argument")
//				}
//			}
//			cfg.Width = args[0]
//			cfg.Height = args[1]
//			cfg.Percents = args[2]
//		}
//		file, err := os.OpenFile("config.txt", os.O_CREATE|os.O_WRONLY, 0644)
//		if err != nil {
//			return cfg, err
//		}
//		_, err = file.WriteString(fmt.Sprintf("%dx%d %d%", cfg.Width, cfg.Height, cfg.Percents))
//		if err != nil {
//			return cfg, err
//		}
//		return cfg, nil
//	}
package main

//import (
//	"context"
//	"example.com/m/internal/application"
//	"os"
//)
//
//func main() {
//	ctx := context.Background()
//	// Exit приводит к завершению программы с заданным кодом.
//	os.Exit(mainWithExitCode(ctx))
//}
//
//func mainWithExitCode(ctx context.Context) int {
//	cfg := application.Config{
//		Width:  100,
//		Height: 100,
//	}
//	app := application.New(cfg)
//	return app.Run(ctx)
//}
