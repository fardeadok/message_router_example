package main

import (
	"fmt"
	"regexp"
)

// ReaderBytesInterface интерфейс с методом ReadBytes
type ReaderBytesInterface interface {
	ReadBytes(path string, message []byte)
}

// Route структура для хранения паттернов и обработчиков
type Route struct {
	Pattern *regexp.Regexp
	Handler ReaderBytesInterface
}

// Router структура для хранения маршрутов
type Router struct {
	routes []Route
}

// NewRouter создает новый экземпляр Router
func NewRouter() *Router {
	return &Router{}
}

// Handle для добавления маршрутов
func (r *Router) Handle(pattern string, handler ReaderBytesInterface) {
	regex := regexp.MustCompile(pattern)
	r.routes = append(r.routes, Route{Pattern: regex, Handler: handler})
}

// Push принимает путь и сообщение, и обрабатывает его
func (r *Router) Push(path string, message []byte) {
	for _, route := range r.routes {
		if route.Pattern.MatchString(path) {
			// Отправляем путь и сообщение в экземпляр ReaderBytesInterface
			route.Handler.ReadBytes(path, message)
			return
		}
	}
	fmt.Printf("Not Found: %s\n", path)
}

type HomeSender struct{}

func (h *HomeSender) ReadBytes(path string, message []byte) {
	fmt.Printf("HomeSender - Path: %s, 	Message: %s\n", path, string(message))
}

type AboutSender struct{}

func (a *AboutSender) ReadBytes(path string, message []byte) {
	fmt.Printf("AboutSender - Path: %s, 	Message: %s\n", path, string(message))
}

type UsersSender struct{}

func (u *UsersSender) ReadBytes(path string, message []byte) {
	fmt.Printf("UsersSender - Path: %s, 	Message: %s\n", path, string(message))
}

func main() {
	router := NewRouter()

	router.Handle("^/$", &HomeSender{})
	router.Handle("^/about$", &AboutSender{})
	router.Handle("^/users/.*$", &UsersSender{})

	fmt.Println("Starting server on :8080...")

	// Имитация вызовов метода Push
	router.Push("/", []byte("Request to Home"))
	router.Push("/about", []byte("Request to About"))
	router.Push("/users/123", []byte("Request to Users"))
	router.Push("/nonexistent", []byte("Request to Nonexistent URL"))
}
