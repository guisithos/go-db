package app

import (
	"github.com/guisithos/go-db/database"
)

type ProtocoloData struct {
	NrProtocolo    string
	NrSeqProtocolo int
}

func GetProtocoloData() (ProtocoloData, error) {
	db, err := database.ConnectToDatabase()
	if err != nil {
		return ProtocoloData{}, err
	}
	defer db.Close()

	query := "select nr_protocolo, nr_seq_protocolo from protocolo_convenio where nr_seq_protocolo = 210"
	var nrProtocolo string
	var nrSeqProtocolo int

	err = db.QueryRow(query).Scan(&nrProtocolo, &nrSeqProtocolo)
	if err != nil {
		return ProtocoloData{}, err
	}

	dadosProtocolo := ProtocoloData{
		NrProtocolo:    nrProtocolo,
		NrSeqProtocolo: nrSeqProtocolo,
	}

	db.Close()

	return dadosProtocolo, nil

}
