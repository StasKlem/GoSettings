package checks

// Check интерфейс для всех проверок
type Check interface {
	Name() string        // Название проверки
	Run() (string, bool) // Результат выполнения проверки и признак успеха
}
