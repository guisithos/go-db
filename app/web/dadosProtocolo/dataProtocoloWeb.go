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
	tmpl := `
    <html>
    <head><title>Teste</title></head>
    <body>
        <h1>Informações do Protocolo</h1>
        <p>Protocolo: {{.NrProtocolo}}</p>
        <p>Seq Protocolo: {{.NrSeqProtocolo}}</p>
    </body>
    </html>`

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		return err
	}

	err = t.Execute(w, dadosProtocolo)
	if err != nil {
		return err
	}

	fmt.Println(dadosProtocolo.NrProtocolo, dadosProtocolo.NrSeqProtocolo)

	return nil
}
