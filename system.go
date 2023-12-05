package zenapi

//
// Системные сущности доступные только для чтения
//

// Instrument - валюта
type Instrument struct {
	ID         int     `json:"id"`
	Changed    int     `json:"changed"` // Unix timestamp
	Title      string  `json:"title"`
	ShortTitle string  `json:"shortTitle"` // Трехбуквенный код данной валюты
	Symbol     string  `json:"symbol"`     // Символ валюты
	Rate       float64 `json:"rate"`       // Стоимость единицы валюты в рублях
}

// Company - это банк либо другая платежная организация, в которой могут существовать счета.
type Company struct {
	ID        int    `json:"id"`
	Changed   int    `json:"changed"` // Unix timestamp
	Title     string `json:"title"`
	FullTitle string `json:"fullTitle"`
	Www       string `json:"www"`
	Country   int    `json:"country"` // ID страны
}

// User - пользователь ZenMoney
type User struct {
	ID       int     `json:"id"`
	Changed  int     `json:"changed"` // Unix timestamp
	Login    *string `json:"login,omitempty"`
	Currency int     `json:"currency"`         // Instrument.id. Основная валюта пользователя. В ней система считает балансы и показывает пользователю отчеты
	Parent   *int32  `json:"parent,omitempty"` // User.id. Родительский пользователь семейного учета. Он является администратором и может удалять дочерних пользователей. Для родительского пользователя parent == null
}

// Country - страны
type Country struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Currency int    `json:"currency"`
	Domain   string `json:"domain"` // ru, us, de, etc.
}
