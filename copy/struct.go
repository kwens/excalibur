/**
 * @Author: kwens
 * @Date: 2022/2/22 15:24
 * @Description:
 */

package copy

import (
	"fmt"
	"reflect"
)

// CopyStruct 拷贝结构体，拷贝两个结构体的字段名和对应类型一致的字段
func CopyStruct(src, dst interface{}, fields ...string) (err error) {
	srcVal := reflect.ValueOf(src).Elem()
	dstType := reflect.TypeOf(dst)
	dstVal := reflect.ValueOf(dst)

	// 简单判断下
	if dstType.Kind() != reflect.Ptr {
		err = fmt.Errorf("util:struct:dst must be a struct pointer")
		return
	}
	dstVal = dstVal.Elem()

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < srcVal.NumField(); i++ {
			name := srcVal.Type().Field(i).Name
			//name := srcType.Field(i).Name
			_fields = append(_fields, name)
		}
	}
	if len(_fields) == 0 {
		return
	}
	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		v := dstVal.FieldByName(name)
		bValue := srcVal.FieldByName(name)
		// 中有同名的字段并且类型一致才复制
		if v.IsValid() && v.Kind() == bValue.Kind() && v.Type() == bValue.Type() && v.CanSet() {
			v.Set(bValue)
		}
	}
	return
}
