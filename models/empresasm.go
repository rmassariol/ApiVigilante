package empresasm

import (
	conexao "ApiVigilante/config"
)

//Empresas estrutura
type Empresas struct {
	ID        int64  `gorm:"primary_key;column:CD_EMPRESA" json:"CD_EMPRESA"`
	NmEmpresa string `gorm:"column:NM_EMPRESA" json:"NM_EMPRESA"`
	CdUsuario int64  `gorm:"column:CD_USUARIO" json:"CD_USUARIO"`
}

//TodasEmpresas lista todas
func TodasEmpresas(p []Empresas) ([]Empresas, error) {
	var err error
	db := conexao.ConectarComGorm()
	db.Find(&p)
	return p, err
}

//ListaEmpresa lista todas
func ListaEmpresa(codigo string) (Empresas, error) {
	var p Empresas
	var err error

	db := conexao.ConectarComGorm()
	db.Where("CD_EMPRESA = ?", codigo).First(&p)
	return p, err
}

//InserirEmpresa insere empresa
func InserirEmpresa(p Empresas) (Empresas, error) {
	db := conexao.ConectarComGorm()
	err := db.Create(&p).Error
	return p, err
}

//AlterarEmpresa altera empresa
func AlterarEmpresa(p Empresas) (Empresas, error) {
	var err error
	db := conexao.ConectarComGorm()
	//	db.First(&p)
	//	p.NmEmpresa = "jinzhu 2"
	db.Save(&p)
	return p, err
}

//ApagarEmpresa apaga nova empresa
func ApagarEmpresa(p Empresas) (Empresas, error) {
	var err error
	db := conexao.ConectarComGorm()
	db.Delete(&p)
	return p, err
}

//RodaSQL exemplo para um sql qualquer
// func RodaSQL(w http.ResponseWriter, r *http.Request) {
// 	type empresa struct {
// 		NmEmpresa  string `gorm:"column:NM_EMPRESA" json:"NM_EMPRESA"`
// 		CdUsuario  int    `gorm:"column:CD_USUARIO" json:"CD_USUARIO"`
// 		OutroCampo int    `gorm:"column:OUTRO_CAMPO" json:"OUTRO_CAMPO"`
// 	}

//-----------------------------------------------
//	gorm
//--------------------------------------------
// conexao := conexao.Conectar()
// var empresas []empresa
// conexao.Select("NM_EMPRESA, CD_USUARIO, (select count(*) from tipos_ocorrencias) AS OUTRO_CAMPO ").Find(&empresas)
// w.Header().Set("Content-Type", "application/json")
// json.NewEncoder(w).Encode(empresas)
// defer conexao.Close()
//--------------------------------------------

//------------------
//gorm - nativo
//-----------------------
// rows, err := conexao.Raw("SELECT NM_EMPRESA, CD_USUARIO from empresas").Rows() // (*sql.Rows, error)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }
// var registros []resultado
// for rows.Next() {

// 	var linha resultado
// 	rows.Scan(&linha.NmEmpresa, &linha.CdUsuario)
// 	registros = append(registros, linha)
// }
// defer rows.Close()
// defer conexao.Close()

// json, _ := json.Marshal(registros)
// w.Header().Set("Content-Type", "application/json")
// fmt.Fprint(w, string(json))
//-----------------------

//}

//Todasempresas - todas as empresas
// func Todasempresas(w http.ResponseWriter, r *http.Request) {
// 	db := conexao.Conectar()
// 	var emp []Empresas
// 	db.Find(&emp)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(emp)
// 	defer db.Close()

// type resultado struct {
// 	CdEmpresa int
// 	NmEmpresa string
// }

// conexao := conexao.Conectar()

// rows, err := conexao.Raw("SELECT CD_EMPRESA, NM_EMPRESA from empresas").Rows() // (*sql.Rows, error)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }

// var t []resultado
// for rows.Next() {

// 	var tipoocorrencia resultado
// 	rows.Scan(&tipoocorrencia.CdEmpresa, &tipoocorrencia.NmEmpresa)
// 	t = append(t, tipoocorrencia)
// }

// json, _ := json.Marshal(t)

// w.Header().Set("Contenct-Type", "application/json")
// fmt.Fprint(w, string(json))

// defer rows.Close()
// defer conexao.Close()

//}

//Alterarempresa altera empresa
// func Alterarempresa(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, r.Header.Get("USUARIO"))

// var p Empresas
// //recebendo o json
// err := json.NewDecoder(r.Body).Decode(&p)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }
// conexao := conexao.Conectar()
// conexao.Model(&p).Updates(map[string]interface{}{"NM_EMPRESA": p.NmEmpresa, "CD_USUARIO": p.CdUsuario})
// json.NewEncoder(w).Encode(p)
// defer conexao.Close()
//}

//Apagarempresa apaga nova empresa
// func Apagarempresa(w http.ResponseWriter, r *http.Request) {
// 	var p Empresas
// 	err := json.NewDecoder(r.Body).Decode(&p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	db := conexao.Conectar()
// 	db.Delete(&p)
// 	json.NewEncoder(w).Encode(p)
// 	defer db.Close()
// }
