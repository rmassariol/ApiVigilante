package utils

import (
	conexao "ApiVigilante/config"
	"ApiVigilante/models/usuariosm"
)

//RetornaJson retorna um json
// func RetornaJson(p string) string {

// 	texto := p
// 	return texto
// }

//ValidaAcesso funcao que valida o acesso do usuario
func ValidaAcesso(User string, password string, token string) (bool, string) {

	if token == "" {
		var p usuariosm.Usuarios

		db := conexao.ConectarComGorm()
		err := db.Where("nm_login = ?", User).First(&p).Error
		if err != nil {
			return false, ""
		}

		println(p.DsSenha)

		if p.DsSenha == "123" {
			return true, GeraToken()
		}
	} else {
		return true, "123455789"
	}

	return false, ""

}

//GeraToken gera o token de acesso
func GeraToken() string {

	db := conexao.ConectarComGorm()
	db.Exec("INSERT INTO token (DS_TOKEN, DT_TOKEN) VALUES (?, CURRENT_DATE)", "123456789")
	return "1234567890"
}
