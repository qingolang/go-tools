package main

import (
	"fmt"
	"strings"
)

func main() {
	// 初始化db
	database(connString)
	// 查询表
	tableList := showTables(dbName)
	content := ``
	if len(tablePrefix) == 0 {
		content += `# ` + dbName + "\n"
		for _, table := range tableList {
			content = createContent(content, table)
		}
	} else {
		for _, prefix := range tablePrefix {
			preRes := strings.Split(prefix, ":")
			content += `# 模块` + preRes[1] + "\n"
			for _, table := range tableList {
				if strings.Contains(table, preRes[0]) {
					content = createContent(content, table)
				}
			}
			content += "\n"
		}
	}

	wirteData(outFilePath, []byte(content))
	fmt.Println("----------------成功！")
}
