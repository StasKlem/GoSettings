package checks

import (
	"os/exec"
)

// GitConfigCheck проверка значения git config
type GitConfigCheck struct{}

// Name возвращает название проверки
func (g GitConfigCheck) Name() string {
	return "Проверка значения git config"
}

// Run выполняет проверку
func (g GitConfigCheck) Run() (string, bool) {
	cmd := exec.Command("git", "config", "--list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Ошибка при выполнении git config: " + err.Error(), false
	}

	// Преобразуем вывод в строку
	config := string(output)
	if config == "" {
		return "Git config не настроен.", false
	}

	return "Git config:\n" + config, true
}
