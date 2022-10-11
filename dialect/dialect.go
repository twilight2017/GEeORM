package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    //用于将Go语言类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) //返回某个表是否存在的SQL语句，参数是表名(table)
}

//注册dialect实例, 若新增了某个数据库，调用RegisterDialect即可注册到全局
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

//get dialect according to the name
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return dialect, ok
}
