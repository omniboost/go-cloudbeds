package cloudbeds

type Status string
type TransactionFilter string

var (
	StatusNotConfirmed Status = "not_confirmed"
	StatusCanceled     Status = "canceled"
	StatusCheckedIn    Status = "checked_in"
	StatusCheckedOut   Status = "checked_out"
	StatusNoShow       Status = "no_show"

	TransactionFilterSimpleTransactions TransactionFilter = "simple_transactions"
	TransactionFilterAdjustments        TransactionFilter = "adjustments"
	TransactionFilterAdjustmentsVoids   TransactionFilter = "adjustments_voids"
	TransactionFilterVoids              TransactionFilter = "voids"
	TransactionFilterRefunds            TransactionFilter = "refunds"
)

// Toegestane waarden: "not_confirmed",
// "confirmed",
// "canceled",
// "checked_in",
// "checked_out",
// "no_show"

type Filters struct {
	And []And `json:"and,omitempty"`
	Or  []Or  `json:"or,omitempty"`
}
type Sort struct {
	Field     string `json:"field,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type And struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Field    string `json:"field"`
}

type Or struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Field    string `json:"field"`
}
