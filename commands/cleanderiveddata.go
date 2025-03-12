package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CleanDerivedDataCommand команда для удаления папки DerivedData
type CleanDerivedDataCommand struct{}

// Execute реализует выполнение команды
func (c CleanDerivedDataCommand) Execute() {
	// Получаем путь к папке DerivedData
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка при получении домашней директории:", err)
		return
	}
	derivedDataPath := filepath.Join(homeDir, "Library", "Developer", "Xcode", "DerivedData")

	// Проверяем, существует ли папка
	if _, err := os.Stat(derivedDataPath); os.IsNotExist(err) {
		fmt.Println("Папка DerivedData не найдена.")
		return
	}

	// Удаляем папку
	cmd := exec.Command("rm", "-rf", derivedDataPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка при удалении DerivedData:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Папка DerivedData успешно удалена!")
}

// Name возвращает название команды
func (c CleanDerivedDataCommand) Name() string {
	return "Удалить папку DerivedData"
}
