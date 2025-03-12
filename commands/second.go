package commands

import "fmt"

// SecondCommand команда для выполнения второй команды
type SecondCommand struct {
	proxyService ProxyService
}

func NewSecondCommand(proxyService ProxyService) *SecondCommand {
	return &SecondCommand{
		proxyService: proxyService,
	}
}

// Execute реализует выполнение команды
func (s SecondCommand) Execute() {
	fmt.Println("Вы выбрали вторую команду!")

	username := s.proxyService.GetUsername()
	password := s.proxyService.GetPassword()
	ip := s.proxyService.GetIP()
	port := s.proxyService.GetPort()

	proxyStr := fmt.Sprintf("http://%s:%s@%s:%s", username, password, ip, port)
	fmt.Printf("http_proxy=%s\n", proxyStr)

	// Здесь можно добавить логику для второй команды
}

// Name возвращает название команды
func (s SecondCommand) Name() string {
	return "Вторая команда"
}
