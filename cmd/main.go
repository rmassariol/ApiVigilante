//go get -u github.com/gorilla/mux
//go get -u github.com/jinzhu/gorm

//go get -u github.com/gorilla/handlers
//go get -u github.com/go-sql-driver/mysql

// crtl shift p  (Ferramentas)
// install go  tools
// update go tools
// go mod init nome_projeto
// go run .\main.g
// go install .

//go mod init nome do projeto

// git init
// git clone https://github.com/rmassariol/ApiVigilante.git
//git remote add origin https://github.com/rmassariol/ApiVigilante.git
//https://medium.com/@saumya.ranjan/how-to-create-a-rest-api-in-golang-crud-operation-in-golang-a7afd9330a7b

//gopath
//%USERPROFILE%\go

package main

import (
	empresasc "ApiVigilante/controllers/empresasc"
	"ApiVigilante/controllers/usuariosc"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
	//	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
	//	empresasc "../controllers"
	//"github.com/gorilla/mux"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	//clientes
	//	myRouter.HandleFunc("/", hello).Methods("GET")
	//myRouter.HandleFunc("/hello", pkg.HelloWord).Methods("GET")

	//login]
	//myRouter.HandleFunc("/login", login.Logar).Methods("GET")

	//empresas
	myRouter.HandleFunc("/empresas", empresasc.TodasEmpresas).Methods("GET")
	myRouter.HandleFunc("/empresas/{cd_empresa}", empresasc.ListaEmpresa).Methods("GET")
	myRouter.HandleFunc("/empresas/listasql", empresasc.ListaSQLGormNativo).Methods("GET")
	myRouter.HandleFunc("/empresas", empresasc.InserirEmpresa).Methods("POST")
	myRouter.HandleFunc("/empresas", empresasc.AlterarEmpresa).Methods("PUT")
	myRouter.HandleFunc("/empresas", empresasc.ApagarEmpresa).Methods("DELETE")

	//usuarios
	myRouter.HandleFunc("/usuarios", usuariosc.TodasUsuarios).Methods("GET")
	myRouter.HandleFunc("/usuarios/{cd_usuario}", usuariosc.ListaUsuario).Methods("GET")
	//	myRouter.HandleFunc("/listasql", usuariosc.ListaSQLGormNativo).Methods("GET")
	myRouter.HandleFunc("/usuarios", usuariosc.InserirUsuario).Methods("POST")
	myRouter.HandleFunc("/usuarios", usuariosc.AlterarUsuario).Methods("PUT")
	myRouter.HandleFunc("/usuarios", usuariosc.ApagarUsuario).Methods("DELETE")

	//PORTA
	log.Fatal(http.ListenAndServe(":8500", myRouter))
}

func main() {

	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/vigilante")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// insert, err := db.Query("INSERT INTO tipos_ocorrencias ( NM_TIPO_OCORRENCIA, CD_EMPRESA) VALUES ('TESTE GO', 1)")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()

	fmt.Println("SERVIDOR VIGILANTE ONLINE")
	fmt.Println("SIMPLESTI.COM.BR")
	// Handle Subsequent requests
	handleRequests()
}
