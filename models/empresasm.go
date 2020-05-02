package empresasm

import (
	conexao "ApiVigilante/config"
)

//Empresas estrutura
type Empresas struct {
	ID        int64  `gorm:"primary_key;column:CD_EMPRESA" json:"CD_EMPRESA"`
	NmEmpresa string `gorm:"not null; column:NM_EMPRESA; " json:"NM_EMPRESA"`
	CdUsuario int64  `gorm:"not null; column:CD_USUARIO;" json:"CD_USUARIO"`
}

//Ocorrencias mostrando ocorrencias
type Ocorrencias struct {
	ID           int64  `gorm:"primary_key;column:CD_OCORRENCIA" json:"CD_OCORRENCIA"`
	DsOcorrencia string `gorm:"not null; column:DS_OCORRENCIA; " json:"DS_OCORRENCIA"`
	CdEmpresa    int64  `gorm:"not null; column:CD_EMPRESA; " json:"CD_EMPRESA"`
}

//ResultadoNativo utilizado para retorno do sql nativo
type ResultadoNativo struct {
	CdEmpresa int64         `gorm:"primary_key;column:CD_EMPRESA" json:"CD_EMPRESA"`
	NmEmpresa string        `gorm:"column:NM_EMPRESA" json:"NM_EMPRESA"`
	Oco       []Ocorrencias `gorm:"column:OCO;" json:"OCO"`
}

//TodasEmpresas lista todas
func TodasEmpresas(p []Empresas) ([]Empresas, error) {

	db := conexao.ConectarComGorm()
	err := db.Find(&p).Error
	return p, err
}

//ListaEmpresa lista todas
func ListaEmpresa(codigo string) (Empresas, error) {
	var p Empresas

	db := conexao.ConectarComGorm()
	err := db.Where("CD_EMPRESA = ?", codigo).First(&p).Error
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

	db := conexao.ConectarComGorm()

	err := db.Save(&p).Error
	return p, err
}

//ApagarEmpresa apaga nova empresa
func ApagarEmpresa(p Empresas) (Empresas, error) {

	db := conexao.ConectarComGorm()
	err := db.Delete(&p).Error
	return p, err
}

//adicionaOcorrencias as ocorrencias em um campo do tipo ocorrencias
func adicionaOcorrencias(obj []Ocorrencias, PcdEmpresa int64) ([]Ocorrencias, error) {

	var ltOcorrencias []Ocorrencias

	db := conexao.ConectarComGorm()

	rows, err := db.Raw("SELECT CD_OCORRENCIA, DS_OCORRENCIA , CD_EMPRESA from ocorrencias where CD_EMPRESA=? ORDER BY 1", PcdEmpresa).Rows() // (*sql.Rows, error)
	if err != nil {
		return ltOcorrencias, err
	}

	for rows.Next() {
		var linha Ocorrencias
		rows.Scan(&linha.ID, &linha.DsOcorrencia, &linha.CdEmpresa)
		//tentando adicionar o oco
		ltOcorrencias = append(ltOcorrencias, Ocorrencias{ID: linha.ID, DsOcorrencia: linha.DsOcorrencia, CdEmpresa: linha.CdEmpresa})
	}

	return ltOcorrencias, err
}

//ListaSQLGormNativo funcao em gorm para sql sem padrao
func ListaSQLGormNativo() ([]ResultadoNativo, error) {

	var registros []ResultadoNativo

	// 	------------------
	// gorm - nativo
	// -----------------------
	db := conexao.ConectarComGorm()
	rows, err := db.Raw("SELECT CD_EMPRESA, NM_EMPRESA from empresas ORDER BY 1").Rows() // (*sql.Rows, error)
	if err != nil {
		return registros, err
	}

	for rows.Next() {
		var linha ResultadoNativo
		rows.Scan(&linha.CdEmpresa, &linha.NmEmpresa)

		//adicionando as ocorrencias
		gg, err := adicionaOcorrencias(linha.Oco, linha.CdEmpresa)
		if err != nil {
			return registros, err
		}
		for i := 0; i < len(gg); i++ {
			linha.Oco = append(linha.Oco, gg[i])
		}

		registros = append(registros, linha)

	}
	defer rows.Close()
	defer db.Close()

	return registros, err

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
