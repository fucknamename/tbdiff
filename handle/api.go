package handle

import (
	"errors"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"tbdiff/model"
	"tbdiff/utils"
	"time"
)

// 查看库的数据表
func HandleShowTable(w http.ResponseWriter, r *http.Request, db string, comment bool) {
	if db == "" || db == "/" {
		w.Write([]byte("please set the database name"))
		return
	}
	if db != "s" && db != "l" {
		w.Write([]byte("database only s or l"))
		return
	}

	var builder strings.Builder
	var tables []*databaseTables

	if db == "s" {
		tables, _ = getDB_Tables(true) // 源库数据表
	} else {
		tables, _ = getDB_Tables(false) // 要更新的库数据表
	}

	var count = len(tables)
	builder.WriteString("---------------------------tables---------------------------\n")

	for i := 0; i < count; i++ {
		if comment {
			builder.WriteString("[" + tables[i].TABLE_NAME + "]" + tables[i].TABLE_COMMENT + "\n")
		} else {
			builder.WriteString(tables[i].TABLE_NAME + "\n")
		}
	}

	w.Write([]byte(builder.String()))
}

// 单个数据表备份
func HandleBackUp(w http.ResponseWriter, r *http.Request, table string) {
	if table == "" || table == "/" {
		w.Write([]byte("please set the backup table name"))
		return
	}

	if !checkMySQLDump() {
		w.Write([]byte("mysqldump command not found"))
		return
	}

	sql, err := backupFile(table, DbCfg.DBL) //这里备份的是：【要更新的数据库】中的表
	if err != "" {
		w.Write([]byte(err))
		return
	}

	w.Write([]byte("backup table [" + table + "] success > " + sql))
}

// 表比较
func HandleCompared(w http.ResponseWriter, r *http.Request, table string) {
	var (
		ts, tl   *databaseTables
		tbs, dbs = getDB_Tables(true)  // 源库数据表
		tbl, dbl = getDB_Tables(false) // 要更新的库数据表
		count    = len(tbs)
	)

	if count == 0 {
		w.Write([]byte("The source database no tables"))
		return
	}

	if table == "" || table == "/" || table == "diff" {
		// todo ...
		w.Write([]byte("Full table comparison ... "))
		return
	}

	for i := 0; i < count; i++ {
		if tbs[i].TABLE_NAME == table {
			ts = tbs[i]
			break
		}
	}

	if ts == nil || ts.TABLE_NAME == "" {
		w.Write([]byte("The source database didn't find table: " + table))
		return
	}

	// 表比较
	for _, v := range tbl {
		if v.TABLE_NAME == table {
			tl = v
			break
		}
	}

	sql, err := getDiffColumn(table, dbs, dbl, ts, tl)
	if err != nil {
		w.Write([]byte("update table [" + table + "] faild"))
		return
	}

	// go func(txt string) {
	// 	// todo: write txt to update sql file
	// }(sql)

	w.Write([]byte(sql))
}

func getDiffColumn(table, dbs, dbl string, ts, tl *databaseTables) (string, error) {
	var (
		// sql bytes.Buffer
		sql strings.Builder
		// sCloumns, err = getColumns(true, dbs, []*databaseTables{ts})
		add = tl == nil || tl.TABLE_NAME == "" // 目标库是否要添加表
	)

	if add {
		output, tip := getTableStructAndData(table, DbCfg.DBS)
		if tip != "" {
			return "", errors.New(tip)
		}

		cmd := exec.Command("mysql",
			"-u", DbCfg.DBL.User,
			"-p"+DbCfg.DBL.Password,
			"-h", DbCfg.DBL.Host,
			"-P", strconv.Itoa(DbCfg.DBL.Port),
			DbCfg.DBL.DataBase,
		)
		cmd.Stdin = strings.NewReader(output)
		ot, err := cmd.Output()
		return string(ot), err
	}

	// 1取建表语句
	// 2逐行对比

	// if err != nil {
	// 	return ""
	// }

	// // 遍历
	// for _, v := range sCloumns {

	// }

	return sql.String(), nil
}

func backupFile(table string, db model.MysqlDb) (string, string) {
	output, tip := getTableStructAndData(table, db)
	if tip != "" {
		return "", tip
	}

	// 包含建表语句、初始化数据
	timetag := time.Now().Format("20060102150405")
	sqlfile := `./sql/[bak_` + timetag + `].` + table + `.sql`
	if err := utils.NewFileManager(sqlfile).WriteToFile([]byte(output)); err != nil {
		return "", "failed to write backup sql file"
	}

	return sqlfile, ""
}

func getTableStructAndData(table string, mysqlcfg model.MysqlDb) (string, string) {
	// mysqldump -u root -p123456 -h 127.0.0.1 -P 3306 ff_db kissme > kissme.sql
	cmd := exec.Command("mysqldump",
		"-u", mysqlcfg.User,
		"-p"+mysqlcfg.Password,
		"-h", mysqlcfg.Host,
		"-P", strconv.Itoa(mysqlcfg.Port),
		mysqlcfg.DataBase,
		table)

	output, err := cmd.Output()

	// defer cmd.Process.Release()

	if err == nil {
		return string(output), ""
	}

	if e, ok := err.(*exec.ExitError); ok {
		tip := string(e.Stderr)
		if strings.Contains(tip, "caching_sha2_password") {
			/*
				[mysqld]
				default_authentication_plugin=mysql_native_password
			*/
			return "", "please set mysql config file with :\n-------------------------------\n[mysqld]\ndefault_authentication_plugin=mysql_native_password"
		} else {
			return "", tip
		}
	} else {
		return "", "backup table [" + table + "] faild"
	}
}

func execSqlfile(sqlfile string) error {
	// mysqlcfg := DbCfg.DBL // 要变更的数据库表
	// cmd := exec.Command("mysqldump", "-u", mysqlcfg.User, "-p"+mysqlcfg.Password, "-h", mysqlcfg.Host, "-P", strconv.Itoa(mysqlcfg.Port), mysqlcfg.DataBase, table)

	return nil
}
