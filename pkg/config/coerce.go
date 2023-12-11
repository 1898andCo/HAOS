package config

import (
	"github.com/rancher/mapper"
	"github.com/rancher/mapper/convert"
	"github.com/rancher/mapper/mappers"
)

type Converter func(val interface{}) interface{}

type typeConverter struct {
	mappers.DefaultMapper
	converter Converter
	fieldType string
	mappers   mapper.Mappers
}

func (t *typeConverter) ToInternal(data map[string]interface{}) error {
	return t.mappers.ToInternal(data)
}

func NewTypeConverter(fieldType string, converter Converter) mapper.Mapper {
	return &typeConverter{
		fieldType: fieldType,
		converter: converter,
	}
}

func NewToMap() mapper.Mapper {
	return NewTypeConverter("map[string]", func(val interface{}) interface{} {
		if m, ok := val.(map[string]interface{}); ok {
			obj := make(map[string]string, len(m))
			for k, v := range m {
				obj[k] = convert.ToString(v)
			}
			return obj
		}
		return val
	})
}

func NewToSlice() mapper.Mapper {
	return NewTypeConverter("array[string]", func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return []string{str}
		}
		return val
	})
}

func NewToBool() mapper.Mapper {
	return NewTypeConverter("boolean", func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return str == "true"
		}
		return val
	})
}
