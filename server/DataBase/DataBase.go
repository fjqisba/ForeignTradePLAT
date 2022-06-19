package DataBase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"server/Model"
	"strings"
)

var(
	Instance DataBase
)

const sqlCreateTable_UserROLE = `CREATE TABLE TABLE_USERROLE (
  username VARCHAR (32) PRIMARY KEY,
  password VARCHAR (32),
  role int	COMMENT '1表示管理员,2表示普通用户'
);`

type DataBase struct {
	sql sqlx.DB
}

func init()  {
	err := Instance.initTable()
	if err != nil{
		log.Panicln(err)
	}
}

func (this *DataBase)initTable()error  {
	var err error
	this.sql.DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/Trade")
	if err != nil{
		return err
	}

	err = this.sql.DB.Ping()
	if err != nil{
		return err
	}

	//创建表TABLE_USER
	_,err = this.sql.Exec(sqlCreateTable_UserROLE)
	if err != nil{
		if strings.Contains(err.Error(),"already exists") == false{
			return err
		}
	}
	return nil
}

func (this *DataBase)GetUserRole(userName string,passWord string)(retRole Model.LoginRet)  {
	err := this.sql.Get(&retRole,"select role from TABLE_USERROLE where username=? and password=?",userName,passWord)
	if err != nil{
		return Model.LOGIN_INVALID
	}
	return retRole
}