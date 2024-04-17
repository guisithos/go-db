package app

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleWebpage(w http.ResponseWriter, r *http.Request) {
	dadosProtocolo, err := GetProtocoloData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = renderWebpage(w, dadosProtocolo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderWebpage(w http.ResponseWriter, dadosProtocolo ProtocoloData) error {
	tmpl, err := template.ParseFiles("/home/gui/go-db/app/web/dadosProtocolo/dadosProtocolo.html")
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, dadosProtocolo)
	if err != nil {
		return err
	}

	fmt.Println(dadosProtocolo.NrProtocolo, dadosProtocolo.NrSeqProtocolo)

	return nil
}
