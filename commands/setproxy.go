// package commands

// import (
// 	"fmt"
// 	"os"
// )

// // SetProxyCommand команда для установки прокси
// type SetProxyCommand struct{}

// // Execute реализует выполнение команды
// func (s SetProxyCommand) Execute() {
// 	var username, password, ip, port string

// 	// Запрашиваем данные у пользователя
// 	fmt.Print("Введите логин: ")
// 	fmt.Scan(&username)

// 	fmt.Print("Введите пароль: ")
// 	fmt.Scan(&password)

// 	fmt.Print("Введите IP-адрес прокси: ")
// 	fmt.Scan(&ip)

// 	fmt.Print("Введите порт прокси: ")
// 	fmt.Scan(&port)

// 	// Формируем строку прокси
// 	proxyStr := fmt.Sprintf("http://%s:%s@%s:%s", username, password, ip, port)

// 	// Устанавливаем переменную окружения
// 	os.Setenv("http_proxy", proxyStr)
// 	os.Setenv("https_proxy", proxyStr) // Также устанавливаем для HTTPS

// 	fmt.Println("Настройки прокси успешно установлены!")
// 	fmt.Printf("http_proxy=%s\n", proxyStr)
// }

// // Name возвращает название команды
// func (s SetProxyCommand) Name() string {
// 	return "Установить настройки прокси"
// }

package commands

import (
	"fmt"
	"os"
)

// SetProxyCommand команда для установки прокси
type SetProxyCommand struct {
	proxyService ProxyService
}

// NewSetProxyCommand конструктор для SetProxyCommand
func NewSetProxyCommand(proxyService ProxyService) *SetProxyCommand {
	return &SetProxyCommand{
		proxyService: proxyService,
	}
}

// Execute реализует выполнение команды
func (s *SetProxyCommand) Execute() {
	// Получаем данные из сервиса
	username := s.proxyService.GetUsername()
	password := s.proxyService.GetPassword()
	ip := s.proxyService.GetIP()
	port := s.proxyService.GetPort()

	// Формируем строку прокси
	proxyStr := fmt.Sprintf("http://%s:%s@%s:%s", username, password, ip, port)

	// Устанавливаем переменную окружения
	os.Setenv("http_proxy", proxyStr)
	os.Setenv("https_proxy", proxyStr) // Также устанавливаем для HTTPS

	fmt.Println("Настройки прокси успешно установлены!")
	fmt.Printf("http_proxy=%s\n", proxyStr)
}

// Name возвращает название команды
func (s *SetProxyCommand) Name() string {
	return "Установить настройки прокси"
}
