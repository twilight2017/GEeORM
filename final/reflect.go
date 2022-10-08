package main

import (
	"fmt"
	"reflect"
)

//reflect demo

// func reflectType(x interface{}) {
// 	//方式1：通过类型断言
// 	//方式2：借助反射
// 	obj := reflect.TypeOf(x)
// 	fmt.Println(obj, obj.Name(), obj.Kind())
// }

// type Cat struct{}
// type Dog struct{}

// func main() {
// 	var a float32 = 1.23
// 	reflectType(a)
// 	var b int8 = 10
// 	reflectType(b)
// 	//结构体类型
// 	var c Cat
// 	var d Dog
// 	reflectType(c)
// 	reflectType(d)
// 	var e []int
// 	reflectType(e)
// }
// func reflectValue(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		//v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
// 		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
// 	case reflect.Float32:
// 		//v.Float()从反射中获取浮点数的原始值，然后通过float32()强制类型转换
// 		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
// 	case reflect.Float64:
// 		//v.Float()从反射中获取浮点型的原始值，然后通过float64进行强制类型转换
// 		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
// 	}
// }

// func main() {
// 	var a float32 = 3.15
// 	var b int64 = 100
// 	reflectValue(a)
// 	reflectValue(b)
// 	c := reflect.ValueOf(10)
// 	fmt.Printf("type c :%T\n", c)
// }
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// func main() {
// 	//*int类型空指针
// 	//IsNil()常用于判断指针是否为空
// 	var a *int
// 	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
// 	//nil值
// 	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
// 	//实例化一个匿名结构体
// 	b := struct{}{}
// 	//尝试从结构体中查找"abc"字段
// 	fmt.Println("不存在的结构体成员", reflect.ValueOf(b).FieldByName("abc").IsValid())
// 	//尝试从结构体中查找"abc"方法
// 	fmt.Println("不存在的结构体方法", reflect.ValueOf(b).MethodByName("abc").IsValid())
// 	c := map[string]int{}
// 	//尝试从map中查找一个不存在的键值
// 	fmt.Println("map中不存在的键", reflect.ValueOf(c).MapIndex(reflect.ValueOf("1")).IsValid())
// }

/*
	当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息
	也可以通过字段名去获取指定的字段信息
*/
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// add two methods to s
func (s *student) Study() string {
	msg := "Study everyday"
	fmt.Println(msg)
	return msg
}

func (s *student) Sleep() string {
	msg := "Sleep right now"
	fmt.Println(msg)
	return msg
}

//printMethod
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name %s\n", t.Method(i).Name)
		fmt.Println("method: %s\n", methodType)
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := student{
		Name:  "Tom",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	//通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) //根据索引i获得结构体中的该字段
		fmt.Printf("name: %s index: %d type: %v, json tag: %v\n", field.Name, field.Type, field.Tag.Get("json"))
	}

	//通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s, index: %d, type: %v, json tag: %v\n", scoreField.Name, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(t)

}
