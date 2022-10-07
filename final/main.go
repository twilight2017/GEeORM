// package main

// import (
// 	"fmt"
// 	_ "github.com/mattn/go-sqlite3"
// 	"os"
// 	"reflect"
// 	"strings"
// )

// type Config struct {
// 	Name    string `json:"server-name"`
// 	IP      string `json:"server-ip"`
// 	URL     string `json:"server-url"`
// 	Timeout string `json:"timeout"`
// }

// func readConfig() *Config {
// 	//read from json
// 	config := Config{}
// 	typ := reflect.TypeOf(config)                       //main.Config
// 	value := reflect.Indirect(reflect.ValueOf(&config)) //{}
// 	for i := 0; i < typ.NumField(); i++ {
// 		f := typ.Field(i) //利用反射获取每个字段的Tag属性
// 		if v, ok := f.Tag.Lookup("json"); ok {
// 			//拼接出对应环境变量的名称
// 			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
// 			if env, exist := os.LookupEnv(key); exist {
// 				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
// 			}
// 		}
// 	}
// 	return &config
// }

// func main() {
// 	os.Setenv("CONFIG_SERVER_NAME", "global_server")
// 	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
// 	os.Setenv("CONFIG_SERVER_URL", "geektutu.com")
// 	c := readConfig()
// 	fmt.Printf("%+v", c)
// }
package main

import (
	"final/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//NewEngine 完成数据库新建操作
	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	//新建sql会话
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
