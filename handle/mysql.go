package handle

import (
	"encoding/json"
	"fmt"
	"strings"
	"tbdiff/model"
	"tbdiff/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysql_dbs *gorm.DB
	mysql_dbl *gorm.DB
	DbCfg     model.DatabaseCfg
)

const (
	// link = "root:123456@tcp(172.0.0.1:3306)/ty"
	dsnTmp    = "%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsnTmpCfg = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func InitDB(db_s, db_l string) error {
	var (
		ss = strings.Trim(db_s, "")
		ll = strings.Trim(db_l, "")
		// connection
		cnns = fmt.Sprintf(dsnTmp, ss)
		cnnl = fmt.Sprintf(dsnTmp, ll)
	)

	if ss == "" || ll == "" {
		// 从配置文件读取数据库连接
		if e := getDbConfig(); e == nil {
			cnns = fmt.Sprintf(dsnTmpCfg,
				DbCfg.DBS.User,
				DbCfg.DBS.Password,
				DbCfg.DBS.Host,
				DbCfg.DBS.Port,
				DbCfg.DBS.DataBase)

			cnnl = fmt.Sprintf(dsnTmpCfg,
				DbCfg.DBL.User,
				DbCfg.DBL.Password,
				DbCfg.DBL.Host,
				DbCfg.DBL.Port,
				DbCfg.DBL.DataBase)
		}
	}

	dbs, err := gorm.Open(mysql.Open(cnns))
	if err != nil {
		fmt.Println("目标数据库初始化链接失败")
		return err
	} else {
		mysql_dbs = dbs
	}

	dbl, err := gorm.Open(mysql.Open(cnnl))
	if err != nil {
		fmt.Println("变更数据库初始化链接失败")
		return err
	} else {
		mysql_dbl = dbl
	}

	return nil
}

func getDbConfig() error {
	var (
		err   error
		lines []string
	)

	if lines, err = utils.ReadTxtLines("./mysql.json"); err != nil {
		return err
	}

	content := strings.Join(lines, "")

	if err = json.Unmarshal([]byte(content), &DbCfg); err != nil {
		return err
	}

	return nil
}

func GetS() *gorm.DB {
	return mysql_dbs
}

func GetL() *gorm.DB {
	return mysql_dbl
}
