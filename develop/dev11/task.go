package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*

Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов


Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры
передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного
выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных
(невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен
возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

*/

const layout = "2006-01-02"

type customDate struct {
	time.Time
}

func (c *customDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c customDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format(layout))), nil
}

// Event Model
type Event struct {
	User        int        `json:"user_id"`
	Date        customDate `json:"date"`
	Description string     `json:"description"`
}

// EventStore хранит Евенты в мапе, где ключ - id юзера
type EventStore struct {
	Events map[int][]Event
}

func newEventStore() *EventStore {
	var s EventStore
	s.Events = make(map[int][]Event)
	return &s
}

func (s *EventStore) create(event Event) {
	s.Events[event.User] = append(s.Events[event.User], event)
}

func (s *EventStore) update(lastEvent, newEvent Event) error {
	if s.Events[lastEvent.User] == nil {
		err := errors.New("event dont exists")
		return err
	}
	for i, val := range s.Events[lastEvent.User] {
		if val.Date == lastEvent.Date && val.Description == lastEvent.Description {
			s.Events[lastEvent.User] = deleteFromEvent(s.Events[lastEvent.User], i)
			s.Events[newEvent.User] = append(s.Events[newEvent.User], newEvent)
		}
	}
	return nil
}

func (s *EventStore) delete(event Event) error {
	if s.Events[event.User] == nil {
		err := errors.New("event dont exists")
		return err
	}
	for i, val := range s.Events[event.User] {
		if val.Date == event.Date && val.Description == event.Description {
			s.Events[event.User] = deleteFromEvent(s.Events[event.User], i)
		}
	}
	return nil
}

func deleteFromEvent(slice []Event, s int) []Event {
	return append(slice[:s], slice[s+1:]...)
}

func (s *EventStore) getEventsForDay() ([]Event, error) {
	today := time.Now()

	var res []Event
	for _, v := range s.Events {
		for _, val := range v {
			if val.Date.Day() == today.Day() && val.Date.Month() == today.Month() && val.Date.Year() == today.Year() {
				res = append(res, val)
			}
		}
	}
	return res, nil
}

func (s *EventStore) getEventForWeek() ([]Event, error) {
	today := time.Now()
	var res []Event
	for _, v := range s.Events {
		for _, val := range v {
			valYear, valWeek := val.Date.ISOWeek()
			todayYear, todayWeek := today.ISOWeek()
			if valWeek == todayWeek && valYear == todayYear {
				res = append(res, val)
			}
		}
	}
	return res, nil
}

func (s *EventStore) getEventForMonth() ([]Event, error) {
	today := time.Now()
	var res []Event
	for _, v := range s.Events {
		for _, val := range v {
			if val.Date.Month() == today.Month() && val.Date.Year() == today.Year() {
				res = append(res, val)
			}
		}
	}
	return res, nil
}

func sendJSONRequest(w http.ResponseWriter, status int, message string, types string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if types == "Error" {
		data := struct {
			Message string `json:"Error"`
		}{Message: message}
		json.NewEncoder(w).Encode(data)
		return
	} else if types == "Changes" {
		data := struct {
			Message string `json:"Result"`
		}{Message: message}
		json.NewEncoder(w).Encode(data)
	}
}

func sendJSONEvents(w http.ResponseWriter, status int, events []Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data := struct {
		Message []Event `json:"Result"`
	}{Message: events}
	json.NewEncoder(w).Encode(data)
}

// Service обработчик запросов
type Service struct {
	Events EventStore
}

func newService(store EventStore) *Service {
	return &Service{Events: store}
}

func (s *Service) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	var event Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		sendJSONRequest(w, 400, "error json input", "Error")
		return
	}

	s.Events.create(event)
	sendJSONRequest(w, 200, "event add", "Changes")

}

func (s *Service) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	var event []Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		sendJSONRequest(w, 400, "error json input", "Error")
		return
	}
	err = s.Events.update(event[0], event[1])
	if err != nil {
		return
	}
	sendJSONRequest(w, 200, "event update", "Changes")

}

func (s *Service) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendJSONRequest(w, 500, "error methods", "Error")
		return
	}

	var event Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		sendJSONRequest(w, 400, "error json input", "Error")
		return
	}

	err = s.Events.delete(event)
	if err != nil {
		return
	}
	sendJSONRequest(w, 200, "event delete", "Changes")

}

func (s *Service) eventsForDay(w http.ResponseWriter, r *http.Request) {
	day, err := s.Events.getEventsForDay()
	if err != nil {
		return
	}
	sendJSONEvents(w, 200, day)
}

func (s *Service) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	week, err := s.Events.getEventForWeek()
	if err != nil {
		return
	}
	sendJSONEvents(w, 200, week)
}

func (s *Service) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	month, err := s.Events.getEventForMonth()
	if err != nil {
		return
	}
	sendJSONEvents(w, 200, month)
}

// Handler функция переадресующая http запросы
func Handler(port string) {
	eventStore := newEventStore()
	service := newService(*eventStore)

	http.HandleFunc("/create_event", middleware(service.createEvent))
	http.HandleFunc("/update_event", middleware(service.updateEvent))
	http.HandleFunc("/delete_event", middleware(service.deleteEvent))
	http.HandleFunc("/events_for_day", middleware(service.eventsForDay))
	http.HandleFunc("/events_for_week", middleware(service.eventsForWeek))
	http.HandleFunc("/events_for_month", middleware(service.eventsForMonth))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}

type config struct {
	Port string
}

func main() {
	var conf config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		log.Fatal(err)
	}
	Handler(conf.Port)
}

// ЛОггер запросов
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}
