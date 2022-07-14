package main

import "db2gorm/gen"

//demo
func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local"

	gen.GenerateAll(gen.GenConf{
		Dsn:         dsn,
		WritePath:   "mysql", //生成到指定目录
		Stdout:      false,
		Overwrite:   true,
		PackageName: "test", //生成出来的文件package
		DirName:     "test", //生成的文件夹
	})
}
