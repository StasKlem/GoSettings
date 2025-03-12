package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

// ProxyCommand команда для отображения настроек прокси
type ProxyCommand struct{}

// Execute реализует выполнение команды
func (p ProxyCommand) Execute() {
	fmt.Println("Настройки прокси для macOS:")

	// Получаем список сетевых интерфейсов
	interfaces, err := getNetworkInterfaces()
	if err != nil {
		fmt.Println("Ошибка при получении списка интерфейсов:", err)
		return
	}

	// Для каждого интерфейса выводим настройки прокси
	for _, iface := range interfaces {
		fmt.Printf("\nИнтерфейс: %s\n", iface)
		cmd := exec.Command("networksetup", "-getwebproxy", iface)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Ошибка при получении настроек прокси для интерфейса %s: %v\n", iface, err)
			continue
		}
		fmt.Println(string(output))
	}
}

// Name возвращает название команды
func (p ProxyCommand) Name() string {
	return "Показать настройки прокси для macOS"
}

// getNetworkInterfaces возвращает список сетевых интерфейсов
func getNetworkInterfaces() ([]string, error) {
	cmd := exec.Command("networksetup", "-listallnetworkservices")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Разделяем вывод на строки и убираем первый элемент (заголовок)
	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		lines = lines[1:] // Пропускаем первую строку (заголовок)
	}

	// Убираем пустые строки
	var interfaces []string
	for _, line := range lines {
		if line != "" {
			interfaces = append(interfaces, line)
		}
	}

	return interfaces, nil
}
