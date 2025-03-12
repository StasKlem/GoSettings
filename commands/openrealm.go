package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"
)

// OpenRealmCommand команда для поиска и открытия последнего файла Realm
type OpenRealmCommand struct{}

// Execute реализует выполнение команды
func (o OpenRealmCommand) Execute() {
	// Путь к папке iOS Simulator
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка при получении домашней директории:", err)
		return
	}
	simulatorPath := filepath.Join(homeDir, "Library", "Developer", "CoreSimulator", "Devices")

	// Ищем все файлы .realm
	realmFiles, err := findRealmFiles(simulatorPath)
	if err != nil {
		fmt.Println("Ошибка при поиске файлов Realm:", err)
		return
	}

	if len(realmFiles) == 0 {
		fmt.Println("Файлы Realm не найдены.")
		return
	}

	// Сортируем файлы по времени изменения
	sort.Slice(realmFiles, func(i, j int) bool {
		return realmFiles[i].ModTime.After(realmFiles[j].ModTime)
	})

	// Берем последний измененный файл
	latestFile := realmFiles[0].Path
	fmt.Printf("Найден последний файл Realm: %s\n", latestFile)

	// Открываем файл
	cmd := exec.Command("open", latestFile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	fmt.Println("Файл успешно открыт!")
}

// Name возвращает название команды
func (o OpenRealmCommand) Name() string {
	return "Открыть последний файл Realm"
}

// RealmFile представляет информацию о файле Realm
type RealmFile struct {
	Path    string
	ModTime time.Time
}

// findRealmFiles рекурсивно ищет файлы .realm в указанной директории
func findRealmFiles(root string) ([]RealmFile, error) {
	var realmFiles []RealmFile

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ищем файлы с расширением .realm
		if !info.IsDir() && filepath.Ext(path) == ".realm" {
			realmFiles = append(realmFiles, RealmFile{
				Path:    path,
				ModTime: info.ModTime(),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return realmFiles, nil
}
