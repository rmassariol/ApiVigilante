package empresasc

import (
	empresasm "ApiVigilante/models/empresasm"
	"ApiVigilante/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//importacao

//Teste e para teste
// func Teste(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "Server running in port:")

// }

//validaCampos para validar a entrada de dados
func validaCampos(campo empresasm.Empresas) (bool, string) {
	if campo.NmEmpresa == "" {
		return false, "Nome da empresa esta vazio!"
	}

	if (strconv.FormatInt(campo.CdUsuario, 10) == "") || (campo.CdUsuario == 0) {
		return false, "Codigo do usuario invalido!"
	}

	return true, "OK"
}

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
	codigo := vars["cd_empresa"]

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

//ListaSQLGormNativo retorno de teste
func ListaSQLGormNativo(w http.ResponseWriter, r *http.Request) {

	// ua := r.Header.Get("TOKEN")
	// fmt.Println(ua)

	//validando o acesso
	passou, token := utils.ValidaAcesso(r.Header.Get("USUARIO"), r.Header.Get("SENHA"), r.Header.Get("TOKEN"))
	if passou == false {
		fmt.Fprint(w, `{"SITUACAO" : "RESTRICAO", "DS_SITUACAO": "USUARIO SEM ACESSO!"}`)
		return
	}

	println(token)

	var p []empresasm.ResultadoNativo
	var err error

	p, err = empresasm.ListaSQLGormNativo()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode("erro")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

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
	situacao, resposta := validaCampos(p)
	if situacao == false {
		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(b)
		fmt.Fprint(w, `{"SITUACAO" : "RESTRICAO", "DS_SITUACAO": "`+resposta+`"}`)
		return
	}

	p, err = empresasm.InserirEmpresa(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

//AlterarEmpresa controle
func AlterarEmpresa(w http.ResponseWriter, r *http.Request) {

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
