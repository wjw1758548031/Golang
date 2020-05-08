package drivers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/skyrunner2012/xormplus/xorm"
	"strings"
)

var Engine *xorm.Engine

func init() {
	//当连接数据库是，需要导入数据库驱动，mysql
	db, err := xorm.NewMySQL("mysql", beego.AppConfig.String("dburl"))
	if err != nil {
		logs.Error("connect db error : %s", err.Error())
	}
	//初始化动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	err = db.SetSqlTemplateRootDir("controllers/sqltemplate").InitSqlTemplate(xorm.SqlTemplateOptions{Extension: ".stpl"})
	if err != nil {
		logs.Error("set sql template root dir error : %s", err.Error())
	}
	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	err = db.StartFSWatcher()
	if err != nil {
		logs.Error("set start fs watcher error : %s", err.Error())
	}

	//可以默认自带前缀，在session时会触发
	/*tbMapper := NewCustMapper(core.SnakeMapper{}, "_do", "t_")
	db.SetTableMapper(tbMapper)*/

	//开启pool
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)

	//开启log，方便调试
	db.ShowSQL(true)

	Engine = db

}

func NewCustMapper(mapper core.IMapper, trimSuffix string, addPrefix string) CustMapper {
	return CustMapper{mapper, trimSuffix, addPrefix}
}

type CustMapper struct {
	Mapper        core.IMapper
	TrimObjSuffix string
	AddTabPrefix  string
}

//以下并不知道有什么用

func (mapper CustMapper) Obj2Table(name string) string {
	tabName := mapper.Mapper.Obj2Table(name)
	if strings.HasSuffix(tabName, mapper.TrimObjSuffix) {
		return mapper.AddTabPrefix + tabName[:len(tabName)-len(mapper.TrimObjSuffix)]
	} else {
		return mapper.AddTabPrefix + tabName
	}
}

func (mapper CustMapper) Table2Obj(name string) string {
	return mapper.Mapper.Table2Obj(name[len(mapper.AddTabPrefix):] + mapper.TrimObjSuffix)
}
