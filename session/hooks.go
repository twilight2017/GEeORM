package session

import (
	"geeorm/log"
	"reflect"
)

//Hooks constants
onst(
	BeforeQuery = "BeforeQuery"
	AfterQuery = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert = "AfterInsert"
)

//CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}){
	fmt := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil{
		fm =reflect.Valueof(value).MethodByName(method)
	}
	params := []reflect.Value(reflect.ValueOf(s))
	if fm.isValid(){
		if v := fm.Call(params); len(v) > 0{
			if err, ok := v[0].Interface{}.(error);ok{
				log.Error(err)
			}
		}
	}
	return
}
