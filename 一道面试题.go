// [1,2,3,[4,5,6, [7,8,9]]]
// 将上面的数值转换为整形切片
// 题记不清楚了，貌似不太对
func ChangtoArr(arr []interface) []int{
	var res []int

	for _, v := range arr {
		k := reflect.TypeOf(v)
		kind := k.Kind()

		switch kind {
		case reflect.Int:
			res = append(res, v)
		case re
		}
	}
}


go reflect反射
主要是获取数据值
TypeOf 获取动态类型信息,返回type类型值
ValueOf 获取数据值
package main

import (
	"fmt"
	"reflect"
)
//定义一个Enum类型
type Enum int
const (
	Zero Enum = 0
)
type Student struct {

	Name string
	Age  int
}

func main() {

	//定义一个Student类型的变量
	var stu Student

	//获取结构体实例的反射类型对象
	typeOfStu := reflect.TypeOf(stu)

	//显示反射类型对象的名称和种类
	fmt.Println(typeOfStu.Name(), typeOfStu.Kind())

	//获取Zero常量的反射类型对象
	typeOfZero := reflect.TypeOf(Zero)

	//显示反射类型对象的名称和种类
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
}


// 查询SQL
func Query(result interface{}, sql string, values ...interface{}) error {
	type1 := reflect.TypeOf(result)
	if type1.Kind() != reflect.Ptr {
		fmt.Println("第一个参数必须是指针")
	}

	type2 := type1.Elem() //指针后面的，表示的值
	if type2.Kind() != reflect.Slice {
		return errors.New("第一个参数必须指向切片")
	}

	type3 := type2.Elem()
	if type3.Kind() != reflect.Ptr {
		return errors.New("切片需要指针类型")
	}

	// 没有封装gorm
	res := []*User{}
	rows, _ := db.Raw(sql, values...).Rows()
	for rows.Next() {
		u := &User{}
		db.ScanRows(row, u)
		res = append(res, u)
	}
	

	// 反射封装
	for rows.Next() {
		elem := reflect.New(type3.Elem())  // 创建一个type3对应的类型对象
		db.ScanRows(rows, elem.Interface()) // 赋值elem
		newSlice := reflect.Append(reflect.ValueOf(res).Elem(), elem)
		reflect.ValueOf(res).Elem().Set(newSlice) 
	}

	rows.Close()
}


// 具体场景调用
type User struct {
	Id int
	Name string
}

func main() {
    result := []*User{}

    // 传入的result的地址
	if err := Query(&result, "select * from user where name=?", "owen"); err == nil {
		for i := 0; i < len(result); i++ {
			fmt.Println(*result[i])
		}
	}
}

