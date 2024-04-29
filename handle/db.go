package handle

import "fmt"

const (
	required    = " validate:\"required\""
	sqlDatabase = `SELECT DATABASE() AS dbname`
	sqlTable    = `SELECT TABLE_NAME, TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='%s' AND TABLE_TYPE='BASE TABLE' ORDER BY TABLE_SCHEMA ASC`
	sqlColumn   = `SELECT TABLE_NAME, COLUMN_NAME, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s' ORDER BY ORDINAL_POSITION`
)

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
}

func getDBS_Tables() []*databaseTables {
	var (
		err    error
		tables []*databaseTables
		dbname = DbCfg.DBS.DataBase
	)

	if dbname == "" {
		dbname = getDbNameS()
	}

	if err = GetS().Raw(fmt.Sprintf(sqlTable, dbname)).Find(&tables).Error; err != nil {
		// fmt.Println(err)
		return nil
	}
	return tables
}

func getDBL_Tables() []*databaseTables {
	var (
		err    error
		tables []*databaseTables
		dbname = DbCfg.DBL.DataBase
	)

	if dbname == "" {
		dbname = getDbNameL()
	}

	if err = GetS().Raw(fmt.Sprintf(sqlTable, dbname)).Find(&tables).Error; err != nil {
		// fmt.Println(err)
		return nil
	}
	return tables
}

func getDbNameS() string {
	var (
		dbname string
	)

	_ = GetS().Raw(sqlDatabase).Find(&dbname).Error
	return dbname
}

func getDbNameL() string {
	var (
		dbname string
	)

	_ = GetL().Raw(sqlDatabase).Find(&dbname).Error
	return dbname
}
