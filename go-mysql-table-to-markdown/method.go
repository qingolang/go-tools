package main

import (
	"fmt"
	"os"
	"strings"
)

// showTableStruct
func showTableStruct(tableName string) string {
	res := make(map[string]interface{})
	getDB().Raw("SHOW CREATE TABLE `" + tableName + "`").Scan(res)
	if res["Create Table"] == nil {
		return ""
	}
	return strings.ReplaceAll(res["Create Table"].(string), "ENGINE=InnoDB", "ENGINE=MyISAM")
}

// ShowColumns
func ShowColumns(tableName, dbName string) []map[string]interface{} {
	var list []map[string]interface{}
	getDB().Table("information_schema.COLUMNS").Where("`TABLE_SCHEMA` = ? ", dbName).Where("`TABLE_NAME` = ?", tableName).Find(&list)
	return list
}

// ShowTables
func showTables(dbName string) (tableName []string) {
	if dbName == "" {
		return nil
	}
	db := getDB()
	err := db.Raw(" select table_name from information_schema.tables where table_schema= ? ", dbName).Scan(&tableName).Error
	if err != nil {
		return nil
	}
	return
}

// createContent
func createContent(content, table string) string {
	content += "\n"
	content += `## 表名` + table + "\n"
	content += "\n"
	content += `### 字段解释` + "\n"
	content += "\n"
	content += "|"
	for _, c := range columns {
		content += ` ` + c + ` |`
	}
	content += "\n"
	content += "|"
	for i := 0; i < len(columns); i++ {
		content += ` :---: |`
	}
	content += "\n"
	// 查询表字段
	tableColumns := ShowColumns(table, dbName)
	for _, tableColumn := range tableColumns {
		content += "| "
		for _, tc := range columns {
			content += fmt.Sprint(tableColumn[tc]) + " |"
		}
		content += "\n"
	}
	content += `### 建表语句` + "\n"
	content += "\n"
	content += "```mysql" + "\n"
	// 查询表结构
	content += showTableStruct(table)
	content += "```" + "\n"
	return content
}

// wirteData
func wirteData(path string, data []byte) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	n, _ := f.Seek(0, os.SEEK_END)
	f.WriteAt(data, n)
}
