package orderentity

type PaymentOrder struct {
	IsPaid    bool       `bun:"is_paid"`
	TotalPaid *float64   `bun:"total_paid"`
	Change    *float64   `bun:"change"`
	Method    *PayMethod `bun:"method"`
}

type PayMethod string

// Tipos de cartão
const (
	Dinheiro        PayMethod = "Dinheiro"
	Visa            PayMethod = "Visa"
	MasterCard      PayMethod = "MasterCard"
	Ticket          PayMethod = "Ticket"
	VR              PayMethod = "VR"
	AmericanExpress PayMethod = "American Express"
	Elo             PayMethod = "Elo"
	DinersClub      PayMethod = "Diners Club"
	Hipercard       PayMethod = "Hipercard"
	VisaElectron    PayMethod = "Visa Electron"
	Maestro         PayMethod = "Maestro"
	Alelo           PayMethod = "Alelo"
	PayPal          PayMethod = "PayPal"
	Outros          PayMethod = "Outros"
)

func GetAll() []PayMethod {
	return []PayMethod{
		Visa,
		MasterCard,
		Ticket,
		VR,
		AmericanExpress,
		Elo,
		DinersClub,
		Hipercard,
		VisaElectron,
		Maestro,
		Alelo,
		PayPal,
		Outros,
	}
}
