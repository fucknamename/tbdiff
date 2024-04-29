package model

type CompareCfg struct {
	DatabaseS string // 参考的数据库连接
	DatabaseL string // 修改的数据库连接
	Table     string // 修改的数据库
}

type DatabaseCfg struct {
	DBS MysqlDb `json:"dbs"`
	DBL MysqlDb `json:"dbl"`
}

type MysqlDb struct {
	DataBase string `json:"database"` //数据库名
	Port     int    `json:"port"`     //数据库端口
	Host     string `json:"host"`     //数据库地址
	User     string `json:"user"`     //数据库用户名
	Password string `json:"password"` //数据库连接密码
}
