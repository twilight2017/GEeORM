package schema

import (
	"final/dialect"
	"go/ast"
	"reflect"
)

//Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string //约束条件
}

//Schema represents a table of database
type Schema struct {
	Model      interface{} //被映射的对象
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field //记录字段名和Field的映射关系，要获取时无需做Fields的遍历
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

//将任意的对象解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.fieldMap[field.Name] = field
		}
	}
	return schema
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldVars []interface{}
	for _, field := range schema.Fields {
		fieldVars = append(fieldVars, destValue.FieldByName(field.Name).Interface())
	}
	return fieldVars
}
