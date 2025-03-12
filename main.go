package main

import (
	"GoSettings/commands" // Импортируем наш пакет с командами
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// Создаем экземпляр сервиса
	proxyService, err := NewMyProxyService("config.json")
	if err != nil {
		fmt.Println("Ошибка при создании сервиса:", err)
		return
	}

	// Массив команд, реализующих интерфейс Command
	commands := []commands.Command{
		commands.RunChecksCommand{},
		commands.ProxyCommand{},
		commands.NewSecondCommand(proxyService),
		commands.NewSetProxyCommand(proxyService),
		commands.CleanProxyCommand{},
		commands.CleanDerivedDataCommand{},
		commands.OpenRealmCommand{},
		commands.ExitCommand{},
	}

	for {
		// Выводим меню
		fmt.Println("Выберите команду:")
		for i, cmd := range commands {
			fmt.Printf("%d. %s\n", i+1, cmd.Name())
		}

		// Чтение выбора пользователя
		var choice int
		fmt.Print("Введите номер команды: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
			continue
		}

		// Проверка корректности выбора
		if choice < 1 || choice > len(commands) {
			fmt.Println("Неверный выбор. Пожалуйста, выберите команду из списка.")
			continue
		}

		// Выполнение выбранной команды
		commands[choice-1].Execute()
	}
}

// ProxyConfig структура для хранения настроек прокси
type ProxyConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
}

// MyProxyService реализация интерфейса ProxyService
type MyProxyService struct {
	config ProxyConfig
}

// NewMyProxyService создает новый экземпляр MyProxyService
func NewMyProxyService(configFile string) (*MyProxyService, error) {
	// Открываем файл конфигурации
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл конфигурации: %w", err)
	}
	defer file.Close()

	// Декодируем JSON в структуру ProxyConfig
	var config ProxyConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("не удалось декодировать JSON: %w", err)
	}

	return &MyProxyService{config: config}, nil
}

// GetUsername возвращает логин из конфигурации
func (m *MyProxyService) GetUsername() string {
	return m.config.Username
}

// GetPassword возвращает пароль из конфигурации
func (m *MyProxyService) GetPassword() string {
	return m.config.Password
}

// GetIP возвращает IP-адрес из конфигурации
func (m *MyProxyService) GetIP() string {
	return m.config.IP
}

// GetPort возвращает порт из конфигурации
func (m *MyProxyService) GetPort() string {
	return m.config.Port
}
