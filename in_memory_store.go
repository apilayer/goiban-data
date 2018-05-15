package data

type Query struct {
	Country  string
	BankCode string
}

type InMemoryStore struct {
	records map[Query]*BankInfo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		records: make(map[Query]*BankInfo),
	}
}

func (s *InMemoryStore) Find(country string, bankCode string) (*BankInfo, error) {
	key := Query{country, bankCode}

	return s.records[key], nil
}

func (s *InMemoryStore) Store(data BankInfo) (bool, error) {
	key := Query{data.Country, data.Bankcode}

	s.records[key] = &data

	return true, nil
}

func (s *InMemoryStore) Clear(source string) (int, error) {
	deleted := 0
	for k, v := range s.records {
		if v.Source == source {
			delete(s.records, k)
			deleted++
		}
	}

	return deleted, nil
}
