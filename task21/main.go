package main

import "fmt"

func main() {

	client := &client{}
	xmlS := &xmlService{}

	client.receiveXMLToClient(xmlS)

	jsonS := &jsonService{}
	jsAdapter := &jsonAdapter{
		jService: jsonS,
	}

	client.receiveXMLToClient(jsAdapter)
}

// Структура нашего клиента,
type client struct{}

// Наш клиент запрашивает данные, в XML формате
func (c *client) receiveXMLToClient(s service) {
	fmt.Println("Клиент запросил данные от сервиса ")
	s.receiveData()
}

// Клиентский интерфейс
type service interface {
	receiveData()
}

// Сервис который отправляет данные через XML
type xmlService struct {
}

// метод реализующий клиентский интерфейс
func (m *xmlService) receiveData() {
	fmt.Println("Получены данные в формате xml")
}
// структура сервиса, от которого получаем данные в формате json
type jsonService struct{}
// запрашиваем json данные
func (j *jsonService) receiveJsonData() {
	fmt.Println("Получены данные в формате json")
}
// структура адаптера
type jsonAdapter struct {
	jService *jsonService
}
// Преобразование полученные данных в формат xml
func (j *jsonAdapter) receiveData() {
	j.jService.receiveJsonData()
	fmt.Println("Адаптер конвертирует данные из json в xml")
}
