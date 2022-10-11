//该文件实现操作数据库相关的代码
package session

import (
	"GReORM/schema"
	"fmt"
	"log"
	"reflect"
	"strings"
)

//解析操作比较耗时，因此将解析的结果保存在成员变量reftable中
func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update reftable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Panic("Model is not set")
	}
	return s.refTable
}

//以下实现数据库表的创建、删除和判断
//利用RefTable()返回的数据库表和字段的信息，拼接出SQL语句，调用原生SQL接口执行
func (s *Session) CreateTable() error {
	//获取数据库表
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s, %s, %s", field.Name, field.Type, field.Type))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
