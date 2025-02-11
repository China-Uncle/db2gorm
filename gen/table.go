package gen

import (
	"strings"

	"github.com/china-uncle/db2gorm/util"
	"gorm.io/gorm"
)

//单个原始表名集合
type Tables []string

/*获取DB下所有的表名*/
func (t *Tables) GetTables(orm *gorm.DB) error {
	res := orm.Raw("show tables").Scan(t)
	if res.Error != nil {
		return res.Error
	}

	//fmt.Printf("%+v", list)
	return nil
}

//单个表在生成struct时所需的全部信息
type TableInfo struct {
	TableName   string //原始表名
	StructName  string //驼峰表名
	PackageName string //全小写的表名
	DirName     string //*.go所在的目录，目前和PackageName一致
	ColumnsInfo []ColumnInfo
}

func GetTableInfo(orm *gorm.DB, tblName string, conf GenConf) (TableInfo, error) {
	var i TableInfo
	i.TableName = tblName
	i.StructName = util.StrCamel(tblName)
	if conf.PackageName != "" {
		i.PackageName = conf.PackageName
	} else {
		i.PackageName = strings.ToLower(i.StructName)
	}
	if conf.DirName != "" {
		i.DirName = conf.DirName
	} else {

		i.DirName = i.PackageName
	}
	info, err := GetTableColumnsInfo(orm, tblName)
	if err != nil {
		return i, err
	}

	i.ColumnsInfo = info

	return i, nil
}

/* 暂时用不到了
func GetTablesInfo(orm *gorm.DB) ([]TableInfo, error){
	var tbls Tables
	if err := tbls.GetTables(orm); err != nil{
		return nil, err
	}

	num := len(tbls)
	tblsInfo := make([]TableInfo, num)
	var err error

	for i, tblName := range tbls{
		tblsInfo[i], err = GetTableInfo(orm, tblName);
		if err != nil{
			return nil, err
		}
	}

	return tblsInfo, nil
}*/
