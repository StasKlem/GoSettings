package commands

import (
	"GoSettings/checks" // Импортируем пакет с проверками
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter" // Импортируем библиотеку для таблиц
)

// RunChecksCommand команда для запуска проверок
type RunChecksCommand struct{}

// Execute реализует выполнение команды
func (r RunChecksCommand) Execute() {
	// Список всех проверок
	checks := []checks.Check{
		checks.HTTPSProxyCheck{},
		checks.GitConfigCheck{},
	}

	// Создаем таблицу
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Проверка", "Результат", "Успех"}) // Заголовки таблицы
	table.SetBorder(false)                                      // Убираем рамку вокруг таблицы
	table.SetRowLine(true)                                      // Добавляем линии между строками

	// Запускаем каждую проверку и добавляем результат в таблицу
	for _, check := range checks {
		result, success := check.Run()
		successText := "✅ Успех"
		if !success {
			successText = "❌ Ошибка"
		}
		table.Append([]string{check.Name(), result, successText})
	}

	// Выводим таблицу
	fmt.Println("Результаты проверок:")
	table.Render()
}

// Name возвращает название команды
func (r RunChecksCommand) Name() string {
	return "Запустить проверки"
}
