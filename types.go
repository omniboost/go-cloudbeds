package cloudbeds

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cydev/zero"
)

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
	ItemID                    string `json:"itemId"`
	PosItemID                 string `json:"posItemId"`
	TaxID                     string `json:"taxId"`
	FeeID                     string `json:"feeId"`
	PaymentID                 string `json:"paymentId"`
}

type CustomTransactionCodes []CustomTransactionCode

type InternalTransactionCode struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

type InternalTransactionCodes []InternalTransactionCode

type Addon struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProductID   string `json:"productId"`
	Price       struct {
		Amount       Int    `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"price"`
}

type Addons []Addon

type Account struct {
	// ID of account
	ID string `json:"id"`
	// Description of the account
	Description string `json:"description"`
	// Name of the account
	Name string `json:"name"`
	// Account category
	// - DEPOSITS
	Category string `json:"category"`
	// Chart of account type
	// - LIABILITIES
	// - REVENUE
	// - ASSETS
	// - EQUITY
	// - EXPENSES
	ChartOfAccountType string `json:"chartOfAccountType"`
}

func (a Account) IsZero() bool {
	return zero.IsZero(a)
}

type Rooms []Room

type Room struct {
	ReservationRoomID string `json:"reservationRoomID"`
	RoomTypeID        string `json:"roomTypeID"`
	RoomTypeName      string `json:"roomTypeName"` // Room Type Name where guest is assigned
	RoomTypeIsVirtual bool   `json:"roomTypeIsVirtual"`
	RoomID            string `json:"roomID"`           // Room ID where guest is assigned
	RoomName          string `json:"roomName"`         // Room Name where guest is assigned
	SubReservationID  string `json:"subReservationID"` // Sub Reservation ID where guest is assigned
}

type AssignedRooms []AssignedRoom

type AssignedRoom struct {
	ReservationRoomID string `json:"reservationRoomID"`
	RoomTypeID        string `json:"roomTypeID"`
	RoomTypeName      string `json:"roomTypeName"` // Room Type Name where guest is assigned
	RoomTypeIsVirtual bool   `json:"roomTypeIsVirtual"`
	RoomCheckIn       Date   `json:"roomCheckIn"`
	RoomCheckOut      Date   `json:"roomCheckOut"`
	RoomID            string `json:"roomID"`            // Room ID where guest is assigned
	RoomName          string `json:"roomName"`          // Room Name where guest is assigned
	SubReservationID  string `json:"subReservationID"`  // Sub Reservation ID where guest is assigned
	RoomTypeNameShort string `json:"roomTypeNameShort"` // Short name of the assigned room type
	RateID            string `json:"rateID"`
	RatePlaneName     string `json:"ratePlanName"`
	StartDate         Date   `json:"startDate"`
	EndDate           Date   `json:"endDate"`
	Adults            Int    `json:"adults"`
	Children          Int    `json:"children"`
	DailyRates        []struct {
		Date Date    `json:"date"`
		Rate float64 `json:"rate"`
	} `json:"dailyRates"`
	RoomTotal  StringFloat `json:"roomTotal"`
	MarketName string      `json:"marketName"`
	MarketCode string      `json:"marketCode"`
	RoomStatus string      `json:"roomStatus"`
}

type GuestRequirements struct {
}

func (g GuestRequirements) IsZero() bool {
	return zero.IsZero(g)
}

type HourMinute struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func (h HourMinute) IsZero() bool {
	return zero.IsZero(h)
}

func (h HourMinute) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.Itoa(h.Hour) + `:` + strconv.Itoa(h.Minute) + `"`), nil
}

func (h *HourMinute) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid time format: %s", value)
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid hour: %s", parts[0])
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid minute: %s", parts[1])
	}

	h.Hour = hour
	h.Minute = minute
	return nil
}

type BalanceDetailed struct {
	SuggestedDeposit StringFloat `json:"suggestedDeposit"`
	SubTotal         float64     `json:"subTotal"`
	AdditionalItems  int         `json:"additionalItems"`
	TaxesFees        float64     `json:"taxesFees"`
	GrandTotal       float64     `json:"grandTotal"`
	Paid             float64     `json:"paid"`
}

type CardsOnFile []CardOnFile

type CardOnFile struct {
	CardID     string `json:"cardID"`
	CardNumber string `json:"cardNumber"`
	CardType   string `json:"cardType"`
}

type CustomFields []CustomField

type CustomField struct {
	CustomFieldName  string `json:"customFieldName"`  // Custom Field Name
	CustomFieldValue string `json:"customFieldValue"` // Custom Field Value
}

// RecipientType represents the type of fiscal document recipient
type RecipientType string

const (
	RecipientTypeCompany RecipientType = "COMPANY"
	RecipientTypePerson  RecipientType = "PERSON"
	RecipientTypeManual  RecipientType = "MANUAL"
)

// RecipientAddress represents the address of a recipient
type RecipientAddress struct {
	Address1 string `json:"address1,omitempty"` // Primary street address line
	Address2 string `json:"address2,omitempty"` // Secondary street address line (e.g. apartment, suite)
	City     string `json:"city,omitempty"`     // City name
	State    string `json:"state,omitempty"`    // State or province
	ZipCode  string `json:"zipCode,omitempty"`  // Postal or ZIP code
	Country  string `json:"country,omitempty"`  // Country code
}

// RecipientCompany represents company information for a recipient
type RecipientCompany struct {
	Name      string `json:"name,omitempty"`      // Legal name of the company
	TaxID     string `json:"taxId,omitempty"`     // Company tax identification number
	TaxIDType string `json:"taxIdType,omitempty"` // Type/classification of the tax ID
	Address1  string `json:"address1,omitempty"`  // Primary street address line
	Address2  string `json:"address2,omitempty"`  // Secondary street address line (e.g. suite, floor)
	City      string `json:"city,omitempty"`      // City name
	State     string `json:"state,omitempty"`     // State or province
	ZipCode   string `json:"zipCode,omitempty"`   // Postal or ZIP code
	Country   string `json:"country,omitempty"`   // Country code
}

// RecipientTaxInfo represents tax information for a recipient
type RecipientTaxInfo struct {
	ID          string `json:"id,omitempty"`          // Tax identification number of the recipient
	CompanyName string `json:"companyName,omitempty"` // Company name associated with the tax record
}

// RecipientContactDetails represents contact details for a recipient
type RecipientContactDetails struct {
	Phone     string `json:"phone,omitempty"`     // Primary phone number
	Gender    string `json:"gender,omitempty"`    // Gender of the recipient
	CellPhone string `json:"cellPhone,omitempty"` // Mobile/cell phone number
	Birthday  string `json:"birthday,omitempty"`  // Date of birth in date-time format (RFC3339)
}

// RecipientDocument represents an identity document for a recipient
type RecipientDocument struct {
	Type           string `json:"type,omitempty"`           // Type of identity document (e.g. passport, national ID)
	Number         string `json:"number,omitempty"`         // Document identification number
	IssuingCountry string `json:"issuingCountry,omitempty"` // Country that issued the document
	IssueDate      string `json:"issueDate,omitempty"`      // Date the document was issued in date-time format (RFC3339)
	ExpirationDate string `json:"expirationDate,omitempty"` // Date the document expires in date-time format (RFC3339)
}

// FiscalDocumentRecipient represents a single recipient associated with a fiscal document
type FiscalDocumentRecipient struct {
	ID             string                   `json:"id,omitempty"`             // Unique identifier of the recipient
	FirstName      string                   `json:"firstName,omitempty"`      // First name of the recipient
	LastName       string                   `json:"lastName,omitempty"`       // Last name of the recipient
	Email          string                   `json:"email,omitempty"`          // Email address of the recipient
	Type           RecipientType            `json:"type,omitempty"`           // Recipient type: COMPANY, PERSON, or MANUAL
	Address        *RecipientAddress        `json:"address,omitempty"`        // Physical address of the recipient
	Company        *RecipientCompany        `json:"company,omitempty"`        // Company information, populated when Type is COMPANY
	Tax            *RecipientTaxInfo        `json:"tax,omitempty"`            // Tax information associated with the recipient
	ContactDetails *RecipientContactDetails `json:"contactDetails,omitempty"` // Additional contact details for the recipient
	Document       *RecipientDocument       `json:"document,omitempty"`       // Identity document information for the recipient
	CountryData    map[string]interface{}   `json:"countryData,omitempty"`    // Arbitrary country-specific fields from guest requirements
}

type Reservations []Reservation

type Reservation struct {
	PropertyID           string          `json:"propertyID"`    // Properties identifier
	ReservationID        string          `json:"reservationID"` // Reservation's unique identifier
	DateCreated          DateTime        `json:"dateCreated"`
	DateModified         DateTime        `json:"dateModified"`
	EstimatedArrivalTime HourMinute      `json:"estimatedArrivalTime"`
	Source               string          `json:"source"`
	SourceID             string          `json:"sourceID"`
	Status               Status          `json:"status"`
	TotalRevenue         float64         `json:"totalRevenue"`
	GuestID              Int             `json:"guestID"`
	ProfileID            Int             `json:"profileID"`
	GuestName            string          `json:"guestName"`
	GuestEmail           string          `json:"guestEmail"`
	StartDate            Date            `json:"startDate"`
	EndDate              Date            `json:"endDate"`
	OrderID              string          `json:"orderID"`
	Origin               string          `json:"origin"`
	Cancellation         string          `json:"cancellation"`
	AllotmentBlockCode   string          `json:"allotmentBlockCode"`
	GroupCode            *string         `json:"groupCode"`
	Adults               Int             `json:"adults"`
	Children             Int             `json:"children"`
	Total                float64         `json:"total"`
	Balance              float64         `json:"balance"`
	BalanceDetailed      BalanceDetailed `json:"balanceDetailed"`
	Assigned             AssignedRooms   `json:"assigned"`
	Unassigned           AssignedRooms   `json:"unassigned"`
	CardsOnFile          CardsOnFile     `json:"cardsOnFile"`
	CustomFields         CustomFields    `json:"customFields"`
	SourceName           string          `json:"sourceName"` // Source of reservation
	ThirdPartyIdentifier string          `json:"thirdPartyIdentifier"`
	GuestList            map[string]struct {
		GuestID string `json:"guestID"` // ID of the guest

		GuestName                   string            `json:"guestName"`
		GuestFirstName              string            `json:"guestFirstName"`
		GuestLastName               string            `json:"guestLastName"`
		Guestgender                 string            `json:"guestGender"` // Toegestane waarden: "M", "F", "N/A"
		GuestEmail                  string            `json:"guestEmail"`
		GuestPhone                  string            `json:"guestPhone"`
		GuestCellPhone              string            `json:"guestCellPhone"`
		GuestAddress                string            `json:"guestAddress"`
		GuestAddress2               string            `json:"guestAddress2"`
		GuestCity                   string            `json:"guestCity"`
		GuestState                  string            `json:"guestState"`
		GuestStatus                 string            `json:"guestStatus"` // Toegestane waarden: "active", "inactive"
		GuestCountry                string            `json:"guestCountry"`
		GuestZip                    string            `json:"guestZip"`
		GuestBirthdate              Date              `json:"guestBirthdate"`
		GuestDocumentType           string            `json:"guestDocumentType"`
		GuestDocumentNumber         string            `json:"guestDocumentNumber"`
		GuestDocumentIssueDate      Date              `json:"guestDocumentIssueDate"`
		GuestDocumentIsseingCountry string            `json:"guestDocumentIssuingCountry"`
		GuestDocumentExpirationDate Date              `json:"guestDocumentExpirationDate"`
		TaxID                       string            `json:"taxID"`        //  Guest's tax ID
		CompanyTaxID                string            `json:"companyTaxID"` // Guest's company tax ID
		CompanyName                 string            `json:"companyName"`  // Guest's company name
		SubReservationID            string            `json:"subReservationID"`
		StartDate                   Date              `json:"startDate"`
		EndDate                     Date              `json:"endDate"`
		AssignedRoom                bool              `json:"assignedRoom"` // Returns true if guest has roomed assigned, false if not
		RoomID                      string            `json:"roomID"`       // Room ID where guest is assigned
		RoomName                    string            `json:"roomName"`     // Room Name where guest is assigned
		RoomTypeName                string            `json:"roomTypeName"` // Room Name where guest is assigned
		IsMainGuest                 bool              `json:"isMainGuest"`
		UnassignedRooms             AssignedRooms     `json:"unassignedRooms"` // List of all rooms that guest is assigned to but not yet checked in
		Rooms                       AssignedRooms     `json:"rooms"`           // List of all rooms that guest is assigned to
		GuestRequirements           GuestRequirements `json:"guestRequirements,omitzero"`
		CustomFields                CustomFields      `json:"customFields"`
		IsAnonymized                bool              `json:"isAnonymized"` //  Flag indicating the guest data was removed upon request
		ReservationRoomID           string            `json:"reservationRoomID"`
	} `json:"guestList"`
	IsAnonymized bool `json:"isAnonymized"` //  Flag indicating the guest data was removed upon request
}

type FiscalDocumentTransactions []FiscalDocumentTransaction

type FiscalDocumentTransaction struct {
		ID                       string   `json:"id"`
		PropertyID               string   `json:"propertyId"`
		SourceID                 string   `json:"sourceId"`
		SourceIdentifier         string   `json:"sourceIdentifier"`
		SourceKind               string   `json:"sourceKind"` // Kind of the source entity (GROUP_PROFILE RESERVATION HOUSE_ACCOUNT ACCOUNTS_RECEIVABLE_LEDGER)
		TransactionDate          string   `json:"transactionDate"`
		GuestName                string   `json:"guestName"`
		Description              string   `json:"description"`
		InternalCode             string   `json:"internalCode"`
		Amount                   float64  `json:"amount"`
		AvailableAmount          float64  `json:"availableAmount"`
		DocumentFiscalizedAmount *float64 `json:"documentFiscalizedAmount"`
		FolioID                  string   `json:"folioId"`
		Status                   string   `json:"status"` // Status of the transaction - PENDING for unpaid transactions, POSTED for paid transactions

		PaidAmount float64 `json:"paidAmount"`

		Allocations []struct {
			ReceiptNumber string `json:"receiptNumber"`
		} `json:"allocations"`
}
