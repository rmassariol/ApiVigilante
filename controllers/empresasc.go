package empresasc

import (
	empresasm "ApiVigilante/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//importacao

//Teste e para teste
// func Teste(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "Server running in port:")

// }

//var NovaEmpresa models.empresasm

//TodasEmpresas - lista tudo
func TodasEmpresas(w http.ResponseWriter, r *http.Request) {

	var p []empresasm.Empresas
	var err error

	p, err = empresasm.TodasEmpresas(p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode("erro")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

	// 	db.Find(&emp)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(emp)
	// 	defer db.Close()
}

//ListaEmpresa - lista tudo
func ListaEmpresa(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	codigo := vars["cd_tipo_ocorrencia"]

	var p empresasm.Empresas
	var err error

	p, err = empresasm.ListaEmpresa(codigo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

	// 	db.Find(&emp)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(emp)
	// 	defer db.Close()
}

//InserirEmpresa controle
func InserirEmpresa(w http.ResponseWriter, r *http.Request) {

	var p empresasm.Empresas
	var err error

	err = json.NewDecoder(r.Body).Decode(&p) //recebendo o json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//tratamento

	p, err = empresasm.InserirEmpresa(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode("erro")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

//AletarEmpresa controle
func AletarEmpresa(w http.ResponseWriter, r *http.Request) {

	var p empresasm.Empresas
	var err error

	err = json.NewDecoder(r.Body).Decode(&p) //recebendo o json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err = empresasm.AlterarEmpresa(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

//ApagarEmpresa apaga nova empresa
func ApagarEmpresa(w http.ResponseWriter, r *http.Request) {
	var p empresasm.Empresas
	var err error

	err = json.NewDecoder(r.Body).Decode(&p) //recebendo o json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err = empresasm.ApagarEmpresa(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
