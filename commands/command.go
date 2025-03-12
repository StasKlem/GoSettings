package commands

// Command интерфейс, который должна реализовывать каждая команда
type Command interface {
	Execute()     // Метод для выполнения команды
	Name() string // Метод для получения названия команды
}
