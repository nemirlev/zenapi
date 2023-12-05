package zenapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// Deletion - информация по удаленному объекту. Некоторые объекты, например Transaction, ReminderMarker, Budget могут быть помечены удаленными полем внутри себя, но все пользовательские объекты, у которых есть id, могут быть удалены окончательно через deletion. При получении объекта deletion получившая сторона обязана удалить у себя этот объект.
// Пример:
// Допустим клиент полностью удалил у себя операцию с id '7DE41EB0-3C61-4DB2-BAE8-BDB2A6A46604'. Тогда он в Diff передает следующий объект deletion:
//
//	{
//	    //...
//	    deletion: [
//	        {
//	            id: '7DE41EB0-3C61-4DB2-BAE8-BDB2A6A46604',
//	            object: 'transaction',
//	            user: 123456,
//	            stamp: 1490008039
//	        }
//	    ]
//	    //...
//	}
type Deletion struct {
	ID     string `json:"id"`     // Object.id
	Object string `json:"object"` // Object.class
	Stamp  int    `json:"stamp"`
	User   int    `json:"user"`
}

type Response struct {
	ServerTimestamp int                `json:"serverTimestamp"`      // Unix timestamp
	ForceFetch      *map[string]string `json:"forceFetch,omitempty"` // [String -> Object.class]
	Instrument      []Instrument       `json:"instrument,omitempty"`
	Country         []Country          `json:"country,omitempty"`
	Company         []Company          `json:"company,omitempty"`
	User            []User             `json:"user,omitempty"`
	Account         []Account          `json:"account,omitempty"`
	Tag             []Tag              `json:"tag,omitempty"`
	Merchant        []Merchant         `json:"merchant,omitempty"`
	Budget          []Budget           `json:"budget,omitempty"`
	Reminder        []Reminder         `json:"reminder,omitempty"`
	ReminderMarker  []ReminderMarker   `json:"reminderMarker,omitempty"`
	Transaction     []Transaction      `json:"transaction,omitempty"`
}

type Request struct {
	CurrentClientTimestamp int                `json:"currentClientTimestamp"` // Unix timestamp
	ServerTimestamp        int                `json:"serverTimestamp"`        // Unix timestamp
	ForceFetch             *map[string]string `json:"forceFetch,omitempty"`   // [String -> Object.class]
	Instrument             []Instrument       `json:"instrument,omitempty"`
	Country                []Country          `json:"country,omitempty"`
	Company                []Company          `json:"company,omitempty"`
	User                   []User             `json:"user,omitempty"`
	Account                []Account          `json:"account,omitempty"`
	Tag                    []Tag              `json:"tag,omitempty"`
	Merchant               []Merchant         `json:"merchant,omitempty"`
	Budget                 []Budget           `json:"budget,omitempty"`
	Reminder               []Reminder         `json:"reminder,omitempty"`
	ReminderMarker         []ReminderMarker   `json:"reminderMarker,omitempty"`
	Transaction            []Transaction      `json:"transaction,omitempty"`
	Deletion               []Deletion         `json:"deletion,omitempty"`
}

// Sync отправляет запрос на отправку и получение изменений из ZenMoney
func (c *Client) Sync(body Request) (Response, error) {
	res, err := c.sendRequest("diff/", "POST", body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to send request: %w", err)
	}

	var result Response
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return Response{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return result, nil
}

// FullSync отправляет запрос на получение всех данных из ZenMoney
func (c *Client) FullSync() (Response, error) { // Change return type to Response
	body := Request{
		CurrentClientTimestamp: int(time.Now().Unix()),
		ServerTimestamp:        0,
	}

	return c.Sync(body)
}
