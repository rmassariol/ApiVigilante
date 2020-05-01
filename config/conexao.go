// go get -u github.com/go-sql-driver/mysql

package conexao

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"

	//so para conectar o mysql
	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"database/sql"
// 	//	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go

// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	//"github.com/jinzhu/gorm"
// )

//ConectarComGorm  abre a conexão com o banco de dados
func ConectarComGorm() (conexao *gorm.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "vigilante"

	bd, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := ""
	// dbName := "vigilante"

	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	// if err != nil {
	// 	panic(err.Error())
	// }

	//conexao, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	//	conexao, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)

	// if err != nil {
	// 	panic(err.Error())
	// }

	return bd
}

//ConectarSemGorm abre a conexão com o banco de dados
func ConectarSemGorm() (conexao *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "vigilante"

	bd, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return bd
}

func main() {

	fmt.Println("CONECTADO AO BANCO DE DADOS")

}
