package commands

import (
	"fmt"
	"os"
)

// ExitCommand команда для выхода из программы
type ExitCommand struct{}

// Execute реализует выполнение команды
func (e ExitCommand) Execute() {
	fmt.Println("Выход из программы.")
	os.Exit(0)
}

// Name возвращает название команды
func (e ExitCommand) Name() string {
	return "Выход"
}
