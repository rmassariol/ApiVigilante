package usuariosm

import (
	conexao "ApiVigilante/config"
)

//Usuarios estrutura
type Usuarios struct {
	ID        int64  `gorm:"primary_key;column:CD_USUARIO" json:"CD_USUARIO"`
	NmLogin   string `gorm:"not null; column:NM_LOGIN; " json:"NM_LOGIN"`
	NmUsuario string `gorm:"not null; column:NM_USUARIO; " json:"NM_USUARIO"`
	DsSenha   string `gorm:"not null; column:DS_SENHA; " json:"DS_SENHA"`
	NrCelular string `gorm:"not null; column:NR_CELULAR; " json:"NR_CELULAR"`
	DsEmail   string `gorm:"not null; column:DS_EMAIL; " json:"DS_EMAIL"`
	FlAtivo   string `gorm:"not null; column:FL_ATIVO; " json:"FL_ATIVO"`
	TpUsuario string `gorm:"not null; column:TP_USUARIO; " json:"TP_USUARIO"`
}

//Exibe teste
func Exibe() string {
	return "teste"
}

//TodasUsuarios lista todas
func TodasUsuarios(p []Usuarios) ([]Usuarios, error) {
	db := conexao.ConectarComGorm()
	err := db.Find(&p).Error
	return p, err
}

//ListaUsuario lista todas
func ListaUsuario(codigo string) (Usuarios, error) {
	var p Usuarios

	db := conexao.ConectarComGorm()
	err := db.Where("CD_USUARIO = ?", codigo).First(&p).Error
	return p, err
}

//InserirUsuario insere usuario
func InserirUsuario(p Usuarios) (Usuarios, error) {
	db := conexao.ConectarComGorm()
	err := db.Create(&p).Error
	return p, err
}

//AlterarUsuario altera usuario
func AlterarUsuario(p Usuarios) (Usuarios, error) {
	db := conexao.ConectarComGorm()
	err := db.Save(&p).Error
	return p, err
}

//ApagarUsuario apaga nova usuario
func ApagarUsuario(p Usuarios) (Usuarios, error) {
	db := conexao.ConectarComGorm()
	err := db.Delete(&p).Error
	return p, err
}

// //adicionaOcorrencias as ocorrencias em um campo do tipo ocorrencias
// func adicionaOcorrencias(obj []Ocorrencias, PcdUsuario int64) ([]Ocorrencias, error) {

// 	var ltOcorrencias []Ocorrencias

// 	db := conexao.ConectarComGorm()

// 	rows, err := db.Raw("SELECT CD_OCORRENCIA, DS_OCORRENCIA , CD_USUARIO from ocorrencias where CD_USUARIO=? ORDER BY 1", PcdUsuario).Rows() // (*sql.Rows, error)
// 	if err != nil {
// 		return ltOcorrencias, err
// 	}

// 	for rows.Next() {
// 		var linha Ocorrencias
// 		rows.Scan(&linha.ID, &linha.DsOcorrencia, &linha.CdUsuario)
// 		//tentando adicionar o oco
// 		ltOcorrencias = append(ltOcorrencias, Ocorrencias{ID: linha.ID, DsOcorrencia: linha.DsOcorrencia, CdUsuario: linha.CdUsuario})
// 	}

// 	return ltOcorrencias, err
// }

// //ListaSQLGormNativo funcao em gorm para sql sem padrao (neste exemplo ele lista as Usuarios com as ocorrencias da mesma)
// func ListaSQLGormNativo() ([]ResultadoNativo, error) {

// 	var registros []ResultadoNativo

// 	// 	------------------
// 	// gorm - nativo
// 	// -----------------------
// 	db := conexao.ConectarComGorm()
// 	rows, err := db.Raw("SELECT CD_USUARIO, NM_USUARIO from Usuarios ORDER BY 1").Rows() // (*sql.Rows, error)
// 	if err != nil {
// 		return registros, err
// 	}

// 	for rows.Next() {
// 		var linha ResultadoNativo
// 		rows.Scan(&linha.CdUsuario, &linha.NmUsuario)

// 		//adicionando as ocorrencias
// 		gg, err := adicionaOcorrencias(linha.Oco, linha.CdUsuario)
// 		if err != nil {
// 			return registros, err
// 		}
// 		for i := 0; i < len(gg); i++ {
// 			linha.Oco = append(linha.Oco, gg[i])
// 		}

// 		registros = append(registros, linha)

// 	}
// 	defer rows.Close()
// 	defer db.Close()

// 	return registros, err

// }
