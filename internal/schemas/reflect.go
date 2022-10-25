package schemas

import (
	"reflect"
	"strings"
	"time"

	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
)

// flattenStructValue will invoke SetResourceData on a struct
func flattenStructValue(structValue any) ([]interface{}, error) {
	res := NewFlatResource()
	if err := SetResourceData(structValue, res); err != nil {
		return nil, err
	}
	return []interface{}{res.Flatten()}, nil
}

func flattenSliceValue(slice []interface{}) ([]interface{}, error) {
	if slice == nil {
		return []interface{}{}, nil
	}
	flatSlice := make([]interface{}, 0, len(slice))
	for _, val := range slice {
		mValue := reflect.Indirect(reflect.ValueOf(val))
		if mValue.Kind() == reflect.Struct {
			res := NewFlatResource()
			if err := SetResourceData(mValue.Interface(), res); err != nil {
				return nil, err
			}
			flatSlice = append(flatSlice, res.Flatten())
		} else {
			flatSlice = append(flatSlice, mValue.Interface())
		}
	}
	return flatSlice, nil
}

// SetResourceData uses reflection to set the resource data
// from an IX-API model.
func SetResourceData(model any, res ResourceSetter) error {
	mValue := reflect.Indirect(reflect.ValueOf(model))
	mType := mValue.Type()

	for i := 0; i < mValue.NumField(); i++ {
		field := mType.Field(i)
		fType := field.Type
		val := mValue.Field(i)
		valType := reflect.Indirect(val).Kind()
		valT := val.Type()

		if !field.IsExported() {
			continue
		}
		if valType == reflect.Invalid {
			// Value is NIL, so we skip it
			continue
		}

		// Get prop name from json property name
		propName := field.Tag.Get("tf")
		if propName == "-" {
			continue // skip field
		}
		if propName == "" {
			propName = strings.Split(field.Tag.Get("json"), ",")[0]
		}

		// Rewrite reserved prop names
		if propName == "connection" {
			propName = "network_connection"
		}

		if field.Name == "Type" {
			continue // Exclude polymorphic type
		}

		if fType.String() == "*ixapi.Date" {
			res.Set(propName, val.Interface().(*ixapi.Date).String())
			continue
		}
		if fType.String() == "*time.Time" {
			res.Set(propName, val.Interface().(*time.Time).Format(time.RFC3339))
			continue
		}
		if valT.PkgPath() == "time" && valT.Name() == "Time" {
			res.Set(propName, val.Interface().(time.Time).Format(time.RFC3339))
			continue
		}

		if valType == reflect.Struct {
			// We flatten this struct using set resource data
			propValue, err := flattenStructValue(val.Interface())
			if err != nil {
				return err
			}
			if err := res.Set(propName, propValue); err != nil {
				return err
			}
		} else if valType == reflect.Slice {
			sLen := val.Len()
			sVal := make([]interface{}, sLen)
			for i := 0; i < sLen; i++ {
				sVal[i] = val.Index(i).Interface()
			}
			propValue, err := flattenSliceValue(sVal)
			if err != nil {
				return err
			}
			if err := res.Set(propName, propValue); err != nil {
				return err
			}
		} else if valType == reflect.Map {
			// For now skip map attrs as they are not reflected in the schema.
			continue
		} else {
			propValue := val.Interface()
			if propValue == nil {
				continue
			}
			// Assume struct on interface??
			if val.Kind() == reflect.Interface {
				var err error
				propValue, err = flattenStructValue(propValue)
				if err != nil {
					return err
				}
			} else if val.Kind() == reflect.Pointer {
				propValue = reflect.Indirect(val).Interface()
			}
			if err := res.Set(propName, propValue); err != nil {
				return err
			}
		}
	}

	return nil
}

// FlattenModel will make a flat model
func FlattenModel(model any) (map[string]interface{}, error) {
	res := NewFlatResource()
	if err := SetResourceData(model, res); err != nil {
		return nil, err
	}
	return res.Flatten(), nil
}

// FlattenModels will create a list of flat models
func FlattenModels(models any) ([]interface{}, error) {
	val := reflect.ValueOf(models)
	length := val.Len()
	flat := make([]interface{}, length)
	for i := 0; i < length; i++ {
		item := val.Index(i)
		flatItem, err := FlattenModel(item.Interface())
		if err != nil {
			return nil, err
		}
		flat[i] = flatItem
	}

	return flat, nil
}
