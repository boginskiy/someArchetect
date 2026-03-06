package router

import (
	"fmt"
	"net/http"
)

// SimpleRouter реализует интерфейс http.Handler
type SimpleRouter struct {
	routes map[string]http.HandlerFunc
}

// Handle добавляет новый маршрут и ассоциированный с ним обработчик
func (sr *SimpleRouter) Handle(path string, handler func(http.ResponseWriter, *http.Request)) {
	if sr.routes == nil {
		sr.routes = make(map[string]http.HandlerFunc)
	}
	sr.routes[path] = handler
}

// ServeHTTP реализует метод интерфейса http.Handler
func (sr *SimpleRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, exists := sr.routes[r.URL.Path]
	if !exists {
		http.NotFound(w, r)
		return
	}
	handler(w, r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать домой!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'.")
}

func main() {
	sRouter := &SimpleRouter{}

	// Регистрация маршрутов
	sRouter.Handle("/", homeHandler)
	sRouter.Handle("/about", aboutHandler)

	// Запуск сервера
	http.ListenAndServe(":8080", sRouter)
}
