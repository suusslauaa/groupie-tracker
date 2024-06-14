package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewApplication(errorLog, infoLog *log.Logger) *Application {
	return &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func NewServer(addr *string, errorLog *log.Logger, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
}

func Web(addr *string) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	// Создаем канал для приема системных сигналов
	sigs := make(chan os.Signal, 1)

	// Перехватываем сигналы типа SIGINT
	signal.Notify(sigs, syscall.SIGINT)

	// Создаем HTTP сервер
	server := NewServer(addr, errorLog, app.Routes())

	// Канал для завершения работы
	done := make(chan bool, 1)

	// Запускаем горутину для обработки сигналов
	go func() {
		sig := <-sigs
		if sig == os.Interrupt {
			infoLog.Printf("Сервер на http://localhost%s выключен", *addr)
		}

		// Контекст с таймаутом для завершения сервера
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Завершаем работу сервера
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("Ошибка при завершении работы сервера:", err)
		}
		done <- true
	}()

	// Запуск HTTP сервера в горутине
	go func() {
		infoLog.Printf("Запуск сервера на http://localhost%s", *addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errorLog.Printf("Ошибка при запуске сервера на http://localhost%s", *addr)
			errorLog.Println(err)
		}
	}()

	// Ожидаем сигнала завершения
	<-done
}
