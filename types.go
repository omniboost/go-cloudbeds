package cloudbeds

type Status string

var (
	StatusNotConfirmed Status = "not_confirmed"
	StatusCanceled     Status = "canceled"
	StatusCheckedIn    Status = "checked_in"
	StatusCheckedOut   Status = "checked_out"
	StatusNoShow       Status = "no_show"
)

// Toegestane waarden: "not_confirmed",
// "confirmed",
// "canceled",
// "checked_in",
// "checked_out",
// "no_show"
