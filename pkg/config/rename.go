package config

import (
	"strings"

	"github.com/rancher/mapper/convert"
	"github.com/rancher/mapper/mappers"
)

type FuzzyNames struct {
	mappers.DefaultMapper
	Names map[string]string
}

// TODO: the error is never set in this code, we don't use the return value
func (f *FuzzyNames) ToInternal(data map[string]interface{}) error {
	for k, v := range data {
		if newK, ok := f.Names[k]; ok && newK != k {
			data[newK] = v
		}
	}
	return nil
}

func (f *FuzzyNames) AddName(name, toName string) {
	if f.Names == nil {
		f.Names = make(map[string]string)
	}
	f.Names[strings.ToLower(name)] = toName
	f.Names[convert.ToYAMLKey(name)] = toName
	f.Names[strings.ToLower(convert.ToYAMLKey(name))] = toName
}
