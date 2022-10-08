//Package:在这一部分实现隔离不同数据库之间的差异
package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

type Dialect interface {
	//DataTypeOf将GO语言数据类型转换为数据库类型值
	DataTypeOf(typ reflect.Value) string
	//返回某个表是否存在的SQL表名，参数是表名tableName
	TableExistSQL(tableName string) (string, []interface{})
}

//注册dialect实例，该方法可将dialect实例注册到全局
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

//获取dialect实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
