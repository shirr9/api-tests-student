package tests

import (
	"log"
	"testing"

	"api-tests-template/internal/utils"
)

// RunSuite загружает окружение и запускает тесты пакета.
func RunSuite(m *testing.M) {
	SetupSuite()
	defer TearDownSuite()
	m.Run()
}

// SetupSuite действия перед запуском сьюта с тестами
func SetupSuite() {
	log.Println("Init environment variables")
	utils.LoadEnv()
}

// TearDownSuite действия после запуска сьюта с тестами
func TearDownSuite() {
	log.Println("Tear down suite")
}

// Precondition выводит текст в консоль с меткой [Precondition]
func Precondition(text string) {
	utils.LogWithLabelAndTimeStamp("Precondition", text)
}
