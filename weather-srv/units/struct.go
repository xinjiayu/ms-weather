package units

import "reflect"

//通过反射给结构体字段赋值。
func ReflectSetData(s interface{}, f, v string) {
	// 得到 reflect.Value
	reVal := reflect.ValueOf(s)
	// 修改值必须是指针类型否则不可行
	if reVal.Kind() != reflect.Ptr {
		//不是指针类型，没法进行修改操作
		return
	}
	reVal = reVal.Elem()
	name := reVal.FieldByName(f)

	if name.Kind() == reflect.String {
		name.SetString(v)
	}
}
