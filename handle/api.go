package handle

import (
	"fmt"
	"net/http"
)

// 表比较
func HandleCompared(w http.ResponseWriter, r *http.Request, table string) {
	var (
		find   = false
		tbs    = getDBS_Tables()
		counts = len(tbs)
	)

	if counts == 0 {
		w.Write([]byte("目标数据库未获取到数据表"))
		return
	}

	if table == "" || table == "/" {
		// 全表对比
		// todo ...
		return
	}

	for i := 0; i < counts; i++ {
		if tbs[i].TABLE_NAME == table {
			find = true
			break
		}
	}

	if !find {
		w.Write([]byte(fmt.Sprintf("目标数据库未找到表：%s ", table)))
		return
	}

	// 表结构比较，并生成sql语句
	// todo ...

	w.Write([]byte("present sql file path ... "))
}
