package commands

import (
	"fmt"
	"os"
)

// CleanProxyCommand команда для установки прокси
type CleanProxyCommand struct {
}

// Execute реализует выполнение команды
func (s CleanProxyCommand) Execute() {
	os.Setenv("http_proxy", "")
	os.Setenv("https_proxy", "")

	fmt.Println("Настройки прокси успешно удалены!")
}

// Name возвращает название команды
func (s CleanProxyCommand) Name() string {
	return "Удалить настройки прокси"
}
