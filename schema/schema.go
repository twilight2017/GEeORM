package schema

import (
	"GReORM/dialect"
	"go/ast"
	"reflect"
)

//Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field          //所有字段
	FieldNames []string          //所有字段名
	fieldMap   map[string]*Field //所有字段名和字段的映射关系
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

//Parse函数将任意的对象解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	//reflect.Indirect是为了获取指针对象指向的实例，modelType返回的是结构体类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), //获取对应结构体的名称作为表名
		fieldMap: make(map[string]*Field),
	}
	//NumField()获取实例的字段的个数
	for i := 0; i < modelType.NumField(); i++ {
		//通过下标获取到对应字段
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				//DataTypeOf转换为数据库的字段类型
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
