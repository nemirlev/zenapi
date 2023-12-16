package zenapi

//
// Пользовательские сущности доступные для чтения, создания, изменения и удаления
//

// Account - счет пользователя
type Account struct {
	ID                    string   `json:"id"`         // UUID
	Changed               int      `json:"changed"`    // Unix timestamp
	User                  int      `json:"user"`       // User.id
	Role                  *int32   `json:"role"`       // User.id
	Instrument            *int32   `json:"instrument"` // Instrument.id
	Company               *int32   `json:"company"`    // Company.id
	Type                  string   `json:"type"`       // Тип счета: ccard - банковская карта, checking - банковский счет, loan - кредит, deposit - депозит, cash - наличные, debt - долги. Счет с типом debt является системным и единственным с таким типом. При добавлении счетов типа loan или deposit нужно указать дополнительные параметры.
	Title                 string   `json:"title"`
	SyncID                []string `json:"syncID"`                // Массив банковских номеров счета. Обычно берутся последние 4 цифры номера счета и последние 4 цифры номеров банковских карт, привязанных к счету.
	Balance               *float64 `json:"balance"`               // Текущий баланс счета
	StartBalance          *float64 `json:"startBalance"`          // Баланс счета в момент открытия. Если тип счета - кредит, то указывается тело кредита.
	CreditLimit           *float64 `json:"creditLimit"`           // >= 0. Кредитный лимит в случае, если тип счета - банковская карта или банковский счет.
	InBalance             bool     `json:"inBalance"`             // Является ли счёт балансовым. Если является, то его баланс учитывается в общем балансе и в отчётах учитываются расходы и доходы по нему.
	Savings               *bool    `json:"savings"`               // Является ли счёт накопительным.
	EnableCorrection      bool     `json:"enableCorrection"`      // Если true, то при распознавании SMS приложение Дзен-мани будет корректировать баланс счёта до его значения в SMS.
	EnableSMS             bool     `json:"enableSMS"`             // Включено ли распознавание SMS по счёту.
	Archive               bool     `json:"archive"`               // Является ли счёт архивным.
	Capitalization        *bool    `json:"capitalization"`        // Для депозита - есть ли капитализация процентов. Для кредита - является ли кредит аннуитетным.
	Percent               *float64 `json:"percent"`               // >= 0 && < 100. Процентная ставка по счету (в процентах).
	StartDate             string   `json:"startDate"`             // 'yyyy-MM-dd'. Дата открытия депозита / кредита.
	EndDateOffset         *int32   `json:"endDateOffset"`         // Срок действия кредита / депозита в промежутках endDateOffsetInterval начиная с даты открытия.
	EndDateOffsetInterval string   `json:"endDateOffsetInterval"` // ('day' | 'week' | 'month' | 'year'). Промежуток между выплатами. Может быть null, тогда считается, что выплата процентов или погашение кредита происходит в конце срока.
	PayoffStep            *int32   `json:"payoffStep"`            // Hаз в сколько payoffInterval происходят выплаты, начиная со следующей даты startDate + payoffInterval. Если payoffInterval == null, то значение должно быть равно 0.
	PayoffInterval        *string  `json:"payoffInterval"`        // ('month' | 'year')?
}

// Tag - категория
type Tag struct {
	ID            string  `json:"id"`      // UUID
	Changed       int     `json:"changed"` // Unix timestamp
	User          int     `json:"user"`    // User.id
	Title         string  `json:"title"`
	Parent        *string `json:"parent"`        // Tag.id. Родительская категория. Допускается степень вложенности не больше 1, т.е. у категории может быть родительская категория, а у родительской категории уже не может быть своего родителя.
	Icon          *string `json:"icon"`          // Id иконки категории
	Picture       *string `json:"picture"`       // Ссылка на картинку для данной категории
	Color         *int64  `json:"color"`         // цвет иконки категории в в виде числа. Рассчитывается по alpha, red, green, blue 0 <= 255. unsigned long color = (a << 24) + (r << 16) + (g << 8) + (b << 0).
	ShowIncome    bool    `json:"showIncome"`    // Является ли категория доходной
	ShowOutcome   bool    `json:"showOutcome"`   // Является ли категория расходной
	BudgetIncome  bool    `json:"budgetIncome"`  // Включена ли категория в расчёт дохода в бюджете
	BudgetOutcome bool    `json:"budgetOutcome"` // Включена ли категория в расчёт расхода в бюджете
	Required      *bool   `json:"required"`      // Являются ли расходы по данной категории обязательными. Если null, то тоже считаются обязательными
}

// Merchant - контрагент операции (продавец, поставщик, плательщик)
type Merchant struct {
	ID      string `json:"id"`      // UUID
	Changed int    `json:"changed"` // Unix timestamp
	User    int    `json:"user"`    // User.id
	Title   string `json:"title"`
}

// Reminder - объект описывающий принцип создания планируемых операций.
// Пример:
// Reminder с параметрами interval: 'day', step: 7, points: [0, 2, 4], startDate: '2017-03-08', endDate: null
// означает, что нужно повторять операции каждую неделю, начиная с 2017-03-08 по средам, пятницам и воскресеньям.
// Потому что 2017-03-08 - среда, значит точка 0 - среда, точка 2 - пятница, точка 4 - воскресенье.
// Каждую неделю - потому как шаг 7 дней.
type Reminder struct {
	ID                string   `json:"id"`                // UUID
	Changed           int      `json:"changed"`           // Unix timestamp
	User              int      `json:"user"`              // User.id
	IncomeInstrument  int      `json:"incomeInstrument"`  // Instrument.id
	IncomeAccount     string   `json:"incomeAccount"`     // Account.id
	Income            float64  `json:"income"`            // >= 0
	OutcomeInstrument int      `json:"outcomeInstrument"` // Instrument.id
	OutcomeAccount    string   `json:"outcomeAccount"`    // Account.id
	Outcome           float64  `json:"outcome"`           // >= 0
	Tag               []string `json:"tag"`               // Tag.id
	Merchant          *string  `json:"merchant"`          // Merchant.id
	Payee             string   `json:"payee"`
	Comment           string   `json:"comment"`
	Interval          *string  `json:"interval"`  // ('day' | 'week' | 'month' | 'year'). Интервал шага. Если null, значит, планируемая без повторения.
	Step              *int32   `json:"step"`      // >= 0. Шаг с которым создаются планируемые операции
	Points            []int    `json:"points"`    // >= 0 && < step. Точки внутри шага, в которых создаются планируемые.
	StartDate         string   `json:"startDate"` // 'yyyy-MM-dd'. с какой даты создавать планируемые.
	EndDate           *string  `json:"endDate"`   // 'yyyy-MM-dd'. До какой даты включительно создавать планируемые. Если null, то бессрочно.
	Notify            bool     `json:"notify"`    // Уведомлять ли о данных операциях.
}

// ReminderMarker - планируемая операция. Поля те же, что и в Reminder, только есть еще дата операции и ее состояние.
type ReminderMarker struct {
	ID                string   `json:"id"`                // UUID
	Changed           int      `json:"changed"`           // Unix timestamp
	User              int      `json:"user"`              // User.id
	IncomeInstrument  int      `json:"incomeInstrument"`  // Instrument.id
	IncomeAccount     string   `json:"incomeAccount"`     // Account.id
	Income            float64  `json:"income"`            // >= 0
	OutcomeInstrument int      `json:"outcomeInstrument"` // Instrument.id
	OutcomeAccount    string   `json:"outcomeAccount"`    // Account.id
	Outcome           float64  `json:"outcome"`           // >= 0
	Tag               []string `json:"tag"`               // Tag.id
	Merchant          *string  `json:"merchant"`          // Merchant.id
	Payee             string   `json:"payee"`
	Comment           string   `json:"comment"`
	Date              string   `json:"date"`     // 'yyyy-MM-dd'. Дата операции
	Reminder          string   `json:"reminder"` // Reminder.id
	State             string   `json:"state"`    // Состояние операции: planned - планируемая, processed - обработанная (внесенная, по ней была создана обычная операция Transaction),  deleted - удаленная
	Notify            bool     `json:"notify"`   // Уведомлять ли о данных операциях
}

// Transaction - денежная операция
//
//	income: incomeAccount=outcomeAccount && income > 0
//	outcome: incomeAccount=outcomeAccount && outcome > 0
//	debt income: type(incomeAccount) == "debt"
//	debt outcome: type(outcomeAccount) == "debt"
//	transfer: other
//
// Примеры: https://github.com/zenmoney/ZenPlugins/wiki/ZenMoney-API#transaction
type Transaction struct {
	ID                  string   `json:"id"`      // UUID
	Changed             int      `json:"changed"` // Unix timestamp
	Created             int      `json:"created"` // Unix timestamp
	User                int      `json:"user"`    // User.id
	Deleted             bool     `json:"deleted"`
	Hold                *bool    `json:"hold"`
	IncomeInstrument    int      `json:"incomeInstrument"`  // Instrument.id. То же самое, что и incomeAccount.instrument
	IncomeAccount       string   `json:"incomeAccount"`     // Account.id
	Income              float64  `json:"income"`            // >= 0. зачислено на счёт incomeAccount.
	OutcomeInstrument   int      `json:"outcomeInstrument"` // Instrument.id. То же самое, что и outcomeAccount.instrument за исключением случая, когда этот счёт - долговой. В случае долговой операции сумма операции всегда пишется в валюте недолгового счёта, а в поле instrument стоит значение instrument недолгового счёта. Валюта же долгового счёта всегда равна user.currency - основной валюте пользователя.
	OutcomeAccount      string   `json:"outcomeAccount"`    // Account.id
	Outcome             float64  `json:"outcome"`           // >= 0. снято со счета outcomeAccount.
	Tag                 []string `json:"tag"`               // Tag.id
	Merchant            *string  `json:"merchant"`          // Merchant.id
	Payee               string   `json:"payee"`
	OriginalPayee       string   `json:"originalPayee"`
	Comment             string   `json:"comment"`
	Date                string   `json:"date"` // 'yyyy-MM-dd'
	Mcc                 *int32   `json:"mcc"`
	ReminderMarker      *string  `json:"reminderMarker"`      // ReminderMarker.id
	OpIncome            *float64 `json:"opIncome"`            // >= 0
	OpIncomeInstrument  *int32   `json:"opIncomeInstrument"`  // Instrument.id. Валюта операции. Допустим была операция снятия долларов с рублевого счёта. Тогда в outcome будет сумма в рублях. А действительную сумму в долларах нужно записать в opOutcome. Данное поле следует использовать только, когда валюта операции отличается от валюты счета.
	OpOutcome           *float64 `json:"opOutcome"`           // >= 0
	OpOutcomeInstrument *int32   `json:"opOutcomeInstrument"` // Instrument.id. Валюта операции
	Latitude            *float64 `json:"latitude"`            // >= -90  && <= 90
	Longitude           *float64 `json:"longitude"`           // >= -180 && <= 180
	QRCode              *string  `json:"qrCode"`
	Source              *string  `json:"source"`
	IncomeBankID        *string  `json:"incomeBankID"`
	OutcomeBankID       *string  `json:"outcomeBankID"`
}

// Budget - бюджет пользователя
type Budget struct {
	Changed     int     `json:"changed"`     // Unix timestamp
	User        int     `json:"user"`        // User.id
	Tag         *string `json:"tag"`         // Tag.id | '00000000-0000-0000-0000-000000000000'. Категория бюджета. Если null, то это бюджет по операциям без категории. Если '00000000-0000-0000-0000-000000000000', то это бюджет совокупный за месяц.
	Date        string  `json:"date"`        // 'yyyy-MM-dd'. Дата начала месяца
	Income      float64 `json:"income"`      // Доходный бюджет
	IncomeLock  bool    `json:"incomeLock"`  // Если true, то сумма income задает точный доходный бюджет по данной категории. Если false, то в качестве бюджета по данной категории берется сумма income и всех доходов по планируемым операциям в этом месяце по данной категории.
	Outcome     float64 `json:"outcome"`     // Расходный бюджет
	OutcomeLock bool    `json:"outcomeLock"` // То же самое, что для incomeLock
}
