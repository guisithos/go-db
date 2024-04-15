package internal

import (
	"database/sql"
	"fmt"
)

func SearchDatabase(db *sql.DB, searchTerm string) ([]string, error) {

	query := fmt.Sprintf("SELECT nr_protocolo FROM protocolo_convenio WHERE nr_seq_protocolo =  '%%%s%%'", searchTerm)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var columnName string
		err := rows.Scan(&columnName)
		if err != nil {
			return nil, err
		}
		results = append(results, columnName)
	}
	return results, nil
}
