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
