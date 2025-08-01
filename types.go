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

type CustomTransactionCode struct {
	ID                        string `json:"id"`
	Version                   int    `json:"version"`
	Name                      string `json:"name"`
	Code                      string `json:"code"`
	SKU                       string `json:"sku"`
	ItemGroup                 string `json:"itemGroup"`
	CustomGeneralLedgerCodeID string `json:"customGeneralLedgerCodeId"`
}

type CustomTransactionCodes []CustomTransactionCode

type InternalTransactionCode struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

type InternalTransactionCodes []InternalTransactionCode
