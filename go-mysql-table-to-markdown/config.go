package main

// dbName 数据库名称
var dbName = "mysql"

// connString
var connString = "root:123456@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"

// outFilePath 输出文件
var outFilePath = "./mysql_engine_tables.md"

// tablePrefix 输出表前缀 len(tablePrefix) == 0 时输出数据库所有表
// 表前缀:描述
var tablePrefix = []string{
	"engine:引擎",
}
