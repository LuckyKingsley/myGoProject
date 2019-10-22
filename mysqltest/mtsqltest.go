package mysqltest

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func MysqlTest() {
	md := mysqlTestInner()
	//md.mysql_insert("INSERT INTO t_apply_check_img set apply_record_id = ?, img_url = ?")
	md.mysql_select("select apply_record_id, img_url from t_apply_check_img")
}

type myselfDb struct {
	db *sql.DB
}

func mysqlTestInner() myselfDb {
	config := mysql.Config{
		User:                 "usr_sheq",
		Passwd:               "Developdb!!)!",
		Net:                  "tcp",
		Addr:                 "39.104.172.208:3306",
		DBName:               "db_qingtongcs",
		AllowNativePasswords: true,
	}
	fmt.Println(config)
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("链接失败", db)
	} else {
		fmt.Println("链接成功", db)
	}

	return myselfDb{
		db: db,
	}
}

func (f *myselfDb) mysql_insert(sql string, args ...interface{}) { //insert  添加数据
	fmt.Println("开始插入")
	stmt, err := f.db.Prepare(sql)
	//defer stmt.Close()
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	stmt.Exec("666", "www.baidu.com")
	fmt.Println("插入成功")
}

func (f *myselfDb) mysql_select(sql_data string) { //select 查询数据
	fmt.Println("sql:", sql_data)
	rows, err := f.db.Query(sql_data)
	if err != nil {
		fmt.Println("查询失败")
	}
	for rows.Next() {
		var applyRecordId int
		var imgUrl string
		err = rows.Scan(&applyRecordId, &imgUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("applyRecordId:", applyRecordId)
		fmt.Println("imgUrl:", imgUrl)
	}

}
