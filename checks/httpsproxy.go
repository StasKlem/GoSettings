package checks

import (
	"os"
)

// HTTPSProxyCheck проверка значения переменной окружения https_proxy
type HTTPSProxyCheck struct{}

// Name возвращает название проверки
func (h HTTPSProxyCheck) Name() string {
	return "Проверка значения https_proxy"
}

// Run выполняет проверку
func (h HTTPSProxyCheck) Run() (string, bool) {
	httpsProxy := os.Getenv("https_proxy")
	if httpsProxy == "" {
		return "Переменная https_proxy не установлена.", false
	}
	return "Значение https_proxy: " + httpsProxy, true
}
