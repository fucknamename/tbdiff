package handle

import (
	"bytes"
	"fmt"
	"net/http"
)

// 表比较
func HandleCompared(w http.ResponseWriter, r *http.Request, table string) {
	var (
		ts, tl   *databaseTables
		tbs, dbs = getDB_Tables(true)  // 源库数据表
		tbl, dbl = getDB_Tables(false) // 要更新的库数据表
		count    = len(tbs)
	)

	if count == 0 {
		w.Write([]byte("源数据库未获取到数据表"))
		return
	}

	if table == "" || table == "/" {
		// 全表对比
		// todo ...
		return
	}

	for i := 0; i < count; i++ {
		if tbs[i].TABLE_NAME == table {
			ts = tbs[i]
			break
		}
	}

	if ts == nil || ts.TABLE_NAME == "" {
		w.Write([]byte(fmt.Sprintf("源数据库未找到表：%s ", table)))
		return
	}

	// 表结构比较，并生成sql语句
	for _, v := range tbl {
		if v.TABLE_NAME == table {
			tl = v
			break
		}
	}

	sql := getDiffColumn(table, dbs, dbl, ts, tl)
	// go func(txt string) {
	// 	// todo: write txt to sql file
	// }(sql)

	w.Write([]byte(sql))
}

func getDiffColumn(table, dbs, dbl string, ts, tl *databaseTables) string {
	var (
		sql bytes.Buffer
		// sCloumns, err = getColumns(true, dbs, []*databaseTables{ts})
		add = tl == nil || tl.TABLE_NAME == "" // 目标库是否要添加表
	)

	if add {
		// if err != nil {
		// 	return ""
		// }

		// // 遍历
		// for _, v := range sCloumns {

		// }
	}

	// 1取建表语句
	// 2逐行对比

	return sql.String()
}
