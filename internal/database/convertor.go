package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func StructToBsonM(item interface{}) (bson.M, error) {
	result := bson.M{}
	v := reflect.ValueOf(item)

	// ต้องแน่ใจว่า item เป็น struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %T", item)
	}

	// วนลูปผ่าน field ของ struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		// ตั้งค่า key และ value ใน bson.M
		result[field.Tag.Get("bson")] = fieldValue.Interface()
	}

	return result, nil
}

func ConvertToBsonItems(items interface{}) ([]interface{}, error) {
	var bsonItems []interface{}
	v := reflect.ValueOf(items)

	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %T", items)
	}

	for i := 0; i < v.Len(); i++ {
		bsonItem, err := StructToBsonM(v.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		bsonItems = append(bsonItems, bsonItem)
	}

	return bsonItems, nil
}
