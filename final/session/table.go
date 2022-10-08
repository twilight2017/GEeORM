//这里存放操作数据库的相关代码
package session

import (
	"final/log"
	"final/schema"
	"fmt"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	//nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

/*实现数据库表的创建、删除和判断是否存在的功能
三者的实现方式是统一的：利用RefTable返回的数据库表和字段的信息，拼接出SQL语句，调用原生SQL接口执行
*/
func (s *Session) CreateTable() error {
	table := s.RefTable() //RefTable保存的是上一次的解析值
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
