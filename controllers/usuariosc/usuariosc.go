package usuariosc

import (
	usuariosm "ApiVigilante/models/usuariosm"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//validaCampos para validar a entrada de dados
func validaCampos(campo usuariosm.Usuarios) (bool, string) {
	if campo.NmUsuario == "" {
		return false, "Nome da usuario esta vazio!"
	}

	return true, "OK"
}

//TodasUsuarios - lista tudo
func TodasUsuarios(w http.ResponseWriter, r *http.Request) {

	var p []usuariosm.Usuarios
	var err error

	p, err = usuariosm.TodasUsuarios(p)

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

//ListaUsuario - lista tudo
func ListaUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	codigo := vars["cd_usuario"]

	var p usuariosm.Usuarios
	var err error

	p, err = usuariosm.ListaUsuario(codigo)

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
// func ListaSQLGormNativo(w http.ResponseWriter, r *http.Request) {

// 	// ua := r.Header.Get("TOKEN")
// 	// fmt.Println(ua)

// 	//validando o acesso
// 	if utils.ValidaAcesso(r.Header.Get("USUARIO"), r.Header.Get("SENHA"), r.Header.Get("TOKEN")) == false {
// 		fmt.Fprint(w, `{"SITUACAO" : "RESTRICAO", "DS_SITUACAO": "USUARIO SEM ACESSO!"}`)
// 		return
// 	}

// 	var p []usuariosm.ResultadoNativo
// 	var err error

// 	p, err = usuariosm.ListaSQLGormNativo()

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		// w.Header().Set("Content-Type", "application/json")
// 		// json.NewEncoder(w).Encode("erro")
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(p)

// }

//InserirUsuario controle
func InserirUsuario(w http.ResponseWriter, r *http.Request) {

	var p usuariosm.Usuarios
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

	p, err = usuariosm.InserirUsuario(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

//AlterarUsuario controle
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {

	var p usuariosm.Usuarios
	var err error
	err = json.NewDecoder(r.Body).Decode(&p) //recebendo o json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p, err = usuariosm.AlterarUsuario(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

//ApagarUsuario apaga nova usuario
func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	var p usuariosm.Usuarios
	var err error

	err = json.NewDecoder(r.Body).Decode(&p) //recebendo o json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err = usuariosm.ApagarUsuario(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
