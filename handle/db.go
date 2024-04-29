package handle

import (
	"fmt"
	"tbdiff/utils"
)

const (
	required       = " validate:\"required\""
	sqlTableCreate = `SHOW CREATE TABLE %s`
	sqlDatabase    = `SELECT DATABASE() AS dbname`
	sqlTable       = `SELECT TABLE_NAME, TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='%s' AND TABLE_TYPE='BASE TABLE' ORDER BY TABLE_SCHEMA ASC`
	sqlColumn      = `SELECT TABLE_NAME, COLUMN_NAME, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' ORDER BY ORDINAL_POSITION`
)

type tableCreate struct {
	Table  string `gorm:"column:table"`
	Create string `gorm:"column:Create Table"`
}

// 数据库表信息
type databaseTables struct {
	TABLE_NAME    string //表名
	TABLE_COMMENT string //表描述
}

// 表字结构信息
type tableStructInfo struct {
	TABLE_NAME     string //所属表名
	COLUMN_NAME    string //字段名
	IS_NULLABLE    string //是否可空
	COLUMN_TYPE    string //字段类型
	COLUMN_COMMENT string //字段描述
	// 索引
	// 长度
}

// 获取数据表的所有字段
func getColumns(source bool, dbname string, tables []*databaseTables) (map[string][]*tableStructInfo, error) {
	var (
		err    error
		size   = utils.Max(len(tables), 1)
		result = make(map[string][]*tableStructInfo, size)
	)

	if size == 0 {
		return nil, nil
	}

	for _, table := range tables {
		var tableInfo []*tableStructInfo
		if source {
			err = GetS().Raw(fmt.Sprintf(sqlColumn, dbname, table.TABLE_NAME)).Find(&tableInfo).Error
		} else {
			err = GetL().Raw(fmt.Sprintf(sqlColumn, dbname, table.TABLE_NAME)).Find(&tableInfo).Error
		}

		if err != nil {
			return nil, err
		}

		result[table.TABLE_NAME] = tableInfo
	}

	return result, nil
}

// 获取数据库的所有表
func getDB_Tables(source bool) ([]*databaseTables, string) {
	var (
		err    error
		dbname string
		tables []*databaseTables
	)

	if source {
		dbname = DbCfg.DBS.DataBase
	} else {
		dbname = DbCfg.DBL.DataBase
	}

	if dbname == "" {
		dbname = getDbName(source)
	}

	if source {
		err = GetS().Raw(fmt.Sprintf(sqlTable, dbname)).Find(&tables).Error
	} else {
		err = GetL().Raw(fmt.Sprintf(sqlTable, dbname)).Find(&tables).Error
	}

	if err != nil {
		// fmt.Println(err)
		return nil, dbname
	}
	return tables, dbname
}

// 获取当前mysql连接的数据库名
func getDbName(source bool) string {
	var (
		dbname string
	)

	if source {
		_ = GetS().Raw(sqlDatabase).Find(&dbname).Error
	} else {
		_ = GetL().Raw(sqlDatabase).Find(&dbname).Error
	}
	return dbname
}
