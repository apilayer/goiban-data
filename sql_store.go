package data

import (
	"database/sql"
	"log"
)

var (
	INSERT_BANK_DATA             *sql.Stmt
	SELECT_SOURCE_ID             *sql.Stmt
	SELECT_BANK_INFORMATION      = "SELECT bankcode, name, zip, city, bic FROM BANK_DATA WHERE bankcode = ? AND country = ?;"
	SELECT_BANK_INFORMATION_STMT *sql.Stmt
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(dbType string, url string) *SQLStore {
	db, err := sql.Open(dbType, url)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
		return nil
	}

	err = prepareStatements(db)

	if err != nil {
		log.Fatalf("DB Prepare Statement error: %v", err)
		return nil
	}

	return &SQLStore{
		db,
	}
}

func (s *SQLStore) Find(countryCode string, bankCode string) (*BankInfo, error) {
	var dbBankcode, dbName, dbZip, dbCity, dbBic string

	err := SELECT_BANK_INFORMATION_STMT.QueryRow(bankCode, countryCode).Scan(&dbBankcode, &dbName, &dbZip, &dbCity, &dbBic)

	if err != nil {
		return nil, err
	}

	return &BankInfo{
		Bankcode: dbBankcode,
		Name:     dbName,
		Zip:      dbZip,
		City:     dbCity,
		Bic:      dbBic,
	}, nil
}

func (s *SQLStore) Store(data BankInfo) (bool, error) {
	_, err := INSERT_BANK_DATA.Exec(
		data.Source,
		data.Bankcode,
		data.Name,
		data.Zip,
		data.City,
		data.Bic,
		data.Country,
		data.CheckAlgo)
	if err != nil {
		log.Fatalf("Failed to insert %v: %v", err, data)
	}

	return true, nil
}

func (s *SQLStore) Clear(source string) (int, error) {
	sourceID, err := getDataSourceId(source)

	log.Printf("Removing entries for source '%v' (%v)", source, sourceID)
	result, err := s.db.Exec("DELETE FROM BANK_DATA WHERE source = ?;", source)

	if err != nil {
		return -1, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(n), nil
}

func getDataSourceId(sourceName string) (int, error) {
	var id int
	result := SELECT_SOURCE_ID.QueryRow(sourceName)

	err := result.Scan(&id)

	if err != nil {
		log.Fatalf("Data source %v not found: %v", sourceName, err)
		return -1, err
	}

	return id, nil
}

func prepareStatements(db *sql.DB) error {
	var err error

	INSERT_BANK_DATA, err = db.Prepare(`INSERT INTO BANK_DATA
		(id, source, bankcode, name, zip, city, bic, country, algorithm, created, last_update)
		VALUES
		(NULL, ?, ?, ?, ?, ?, ?, ?, ?, NULL, NULL);`)

	if err != nil {
		log.Fatalf("Error while preparing statement: %v", err)
		return err
	}

	SELECT_SOURCE_ID, err = db.Prepare(`SELECT id FROM DATA_SOURCE where name = ?`)

	if err != nil {
		log.Fatalf("Error while preparing statement: %v", err)
		return err
	}

	SELECT_BANK_INFORMATION_STMT, err = db.Prepare(SELECT_BANK_INFORMATION)

	if err != nil {
		log.Fatalf("Error while preparing statement: %v", err)
		return err
	}

	return nil
}
