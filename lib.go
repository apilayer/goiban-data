package data

// BankInfo describes data associated with an IBAN
type BankInfo struct {
	Bankcode string `json:"bankCode"`
	Name     string `json:"name"`
	Zip      string `json:"zip,omitempty"`
	City     string `json:"city,omitempty"`
	Bic      string `json:"bic,omitempty"`

	Country   string `json:"-"`
	CheckAlgo string `json:"-"`
	Source    string `json:"-"`
}

// BankDataRepository provides storage mechanisms
type BankDataRepository interface {
	// Find an entry
	Find(countryCode string, bankCode string) (*BankInfo, error)

	// Store bank info
	Store(data BankInfo) (bool, error)

	// Clear all entries from a specific source
	//
	// Returns the number of removed entries and an optional error
	Clear(source string) (int, error)
}
