package main

import (
	"clients_ms/internal/api"
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var t []*api.Table
	//var u []*api.UAR_Mass

	t = append(t, &api.Table{
		TableName:    "users",
		TableComment: "users"})

	t = append(t, &api.Table{
		TableName:    "users1",
		TableComment: "users1"})

	slcIntf, err := ConvertToInterfaceSlice(t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(slcIntf)
}

func ConvertToInterfaceSlice(data interface{}) ([][]interface{}, error) {
	sliceValue := reflect.ValueOf(data)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("data is not a slice")
	}
	sliceLen := sliceValue.Len()
	if sliceLen == 0 {
		return nil, nil
	}
	elemValue := sliceValue.Index(0)
	if !elemValue.IsValid() {
		return nil, fmt.Errorf("slice element is not valid")
	}
	if !elemValue.Elem().CanInterface() {
		return nil, fmt.Errorf("slice element fields are not exported")
	}
	rows := make([][]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		elemValue = sliceValue.Index(i)
		if !elemValue.Elem().CanInterface() {
			return nil, fmt.Errorf("slice element fields are not exported")
		}
		fieldsValue := elemValue.Elem()
		fieldCount := fieldsValue.NumField() - 3
		row := make([]interface{}, fieldCount)
		for j := 0; j < fieldCount; j++ {
			field := fieldsValue.Field(j)
			if !field.CanInterface() {
				fmt.Errorf("field %s is not exported", fieldsValue.Type().Field(j).Name)
				continue
			}
			row[j] = field.Interface()
		}
		rows[i] = row
	}
	return rows, nil
}

func convertPointerSliceToInterfaceSlice(slice interface{}) ([][]interface{}, error) {
	// Получаем значение слайса

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("value is not a slice")
	}

	// Получаем тип элемента слайса
	elementType := sliceValue.Type().Elem()
	if elementType.Kind() != reflect.Ptr {
		return nil, errors.New("slice element type is not a pointer")
	}
	elementType = elementType.Elem()

	// Создаем новый слайс [][]interface{}
	result := make([][]interface{}, sliceValue.Len())

	// Итерируем по элементам слайса и заполняем новый слайс
	for i := 0; i < sliceValue.Len(); i++ {
		elementValue := sliceValue.Index(i)
		if !elementValue.IsValid() {
			continue
		}
		if elementValue.Type().Elem() != elementType {
			return nil, errors.New("slice contains elements of different types")
		}

		row := make([]interface{}, 0, elementValue.Elem().NumField())
		for j := 0; j < elementValue.Elem().NumField(); j++ {

		}
		result[i] = row
	}

	return result, nil
}
