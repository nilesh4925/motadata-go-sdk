package motadatatypes

import (
	"encoding/json"
	"github.com/nilesh4925/motadata-go-sdk/sdkconstant"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	MotadataString string

	MotadataINT int

	MotadataUINT uint

	MotadataUINT8 uint8

	MotadataUINT16 uint16

	MotadataUINT32 uint32

	MotadataFloat64 float64

	MotadataTimeString string

	MotadataMap map[string]interface{}

	MotadataListMap []map[string]interface{}

	MotadataStringMap map[string]string

	MotadataStringList []string

	MotadataMapList []MotadataMap

	MotadataListInterface []interface{}
)

func (value MotadataTimeString) Format() MotadataString {

	return MotadataString(time.Now().Format(string(value)))

}

func (value MotadataString) ToString() string {

	return strings.TrimSpace(string(value))
}

func (value MotadataString) ToUINT8() MotadataUINT8 {

	result, _ := strconv.ParseUint(value.ToString(), 10, 8)

	return MotadataUINT8(result)
}

func (value MotadataString) ToUINT16() MotadataUINT16 {

	result, _ := strconv.ParseUint(value.ToString(), 10, 16)

	return MotadataUINT16(result)
}

func (value MotadataString) ToUINT32() MotadataUINT32 {

	result, _ := strconv.ParseUint(value.ToString(), 10, 32)

	return MotadataUINT32(result)
}

func (context MotadataMap) GetSliceValue(key string) (result []interface{}) {

	if context.Contains(key) {

		return context[key].([]interface{})
	}

	return
}

func (context MotadataMap) Contains(key string) (result bool) {

	if _, found := context[key]; found {

		return true

	}

	return
}

func (context MotadataStringMap) Contains(key string) (result bool) {

	if _, found := context[key]; found {

		return true

	}

	return
}

func (value MotadataString) ReplaceAll(old MotadataString, new MotadataString) MotadataString {

	return MotadataString(strings.ReplaceAll(ToString(value), ToString(old), ToString(new)))

}

func ToString(value interface{}) (result string) {

	if reflect.TypeOf(value).Name() == "string" {

		result = value.(string)

	} else if reflect.TypeOf(value).Name() == "MotadataString" {

		result = string(value.(MotadataString))

	} else if reflect.TypeOf(value).Name() == "int" {

		result = strconv.Itoa(value.(int))

	} else if reflect.TypeOf(value).Name() == "uint" {

		result = strconv.Itoa(int(value.(uint)))

	} else if reflect.TypeOf(value).Name() == "uint64" {

		result = strconv.FormatUint(value.(uint64), 10)

	} else if reflect.TypeOf(value).Name() == "MotadataUINT16" {

		result = strconv.Itoa(int(value.(MotadataUINT16)))

	} else if reflect.TypeOf(value).String() == "[]uint8" {

		result = string(value.([]uint8))

	} else if reflect.TypeOf(value).String() == "[]byte" {

		result = string(value.([]byte))

	} else if reflect.TypeOf(value).Name() == "float64" {

		result = strconv.Itoa(int(value.(float64)))

	} else if reflect.TypeOf(value).Name() == "bool" {

		result = strconv.FormatBool(value.(bool))

	} else if reflect.TypeOf(value).Name() == "int64" {

		result = strconv.Itoa(int(value.(int64)))

	}

	return

}

func ToMap(value interface{}) (result MotadataMap) {

	if value != nil {

		if reflect.TypeOf(value).Name() == "MotadataMap" {

			result = value.(MotadataMap)

		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {

			result = value.(map[string]interface{})
		}

	}

	return
}

func (value MotadataString) IsNotEmpty() bool {

	if value != sdkconstant.BlankString && len(value.TrimSpace()) > 0 {

		return true
	}

	return false
}

func IsNotEmpty(value string) bool {

	if value != sdkconstant.BlankString && len(strings.TrimSpace(value)) > 0 {

		return true
	}

	return false
}

func ToLower(value string) string {

	return strings.ToLower(value)
}

func (value MotadataString) TrimSpace() MotadataString {

	return MotadataString(strings.TrimSpace(string(value)))
}

func (context MotadataMap) GetMotadataStringValue(key string) (result MotadataString) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "MotadataString" {

			result = value.(MotadataString)

		} else if reflect.TypeOf(value).Name() == "string" {

			result = MotadataString(value.(string))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataString(strconv.Itoa(value.(int)))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = MotadataString(strconv.Itoa(value.(MotadataINT).ToInt()))

		} else if reflect.TypeOf(value).String() == "[]uint8" {

			result = MotadataString(value.([]uint8))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataString(strconv.Itoa(int(value.(float64))))

		} else if reflect.TypeOf(value).String() == "*string" {

			result = MotadataString(*value.(*string))

		} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

			result = MotadataString(strconv.Itoa(int(value.(MotadataFloat64))))

		}
	}

	return
}

func (context MotadataMap) GetINTValue(key string) (result MotadataINT) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = value.(MotadataINT)

		} else if reflect.TypeOf(value).Name() == "uint" {

			result = MotadataINT(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = MotadataINT(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = MotadataINT(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = MotadataINT(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = MotadataINT(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataINT(value.(int))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = MotadataINT(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = MotadataINT(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = MotadataINT(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataINT(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataINT(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

			result = MotadataINT(value.(MotadataFloat64))
		}
	}

	return
}

func (value MotadataINT) ToInt() int {

	return int(value)

}

func (context MotadataMap) GetContext(key string) (result MotadataMap) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "MotadataMap" {

			result = value.(MotadataMap)

		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {

			result = value.(map[string]interface{})
		}

	}

	return
}

func (context MotadataMap) GetInterfaceContext(key string) (result MotadataListInterface) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).String() == "[]interface {}" {

			result = value.([]interface{})
		} else if reflect.TypeOf(value).String() == "[]interface {}" {

			result = value.([]interface{})
		}
	}

	return
}

func (value MotadataListInterface) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (context MotadataMap) GetUINTValue(key string) (result MotadataUINT) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "uint" {

			result = MotadataUINT(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = MotadataUINT(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = MotadataUINT(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = MotadataUINT(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = MotadataUINT(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataUINT(value.(int))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = MotadataUINT(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = MotadataUINT(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = MotadataUINT(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataUINT(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataUINT" {

			result = value.(MotadataUINT)

		} else if reflect.TypeOf(context[key]).Name() == "string" {

			output, _ := strconv.ParseInt(strings.TrimSpace(context[key].(string)), 10, 64)

			result = MotadataUINT(output)
		}
	}

	return
}

func (value MotadataUINT16) ToUInt16() uint16 {

	return uint16(value)
}

func (value MotadataUINT16) ToInt() int {

	return int(value)
}

func (value MotadataUINT8) ToUInt16() uint16 {

	return uint16(value)
}

func (value MotadataUINT8) ToUInt8() uint8 {

	return uint8(value)
}

func (value MotadataUINT32) ToUInt32() uint32 {

	return uint32(value)
}

func (context MotadataMap) GetStringValue(key string) (result string) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "string" {

			result = value.(string)

		} else if reflect.TypeOf(value).Name() == "MotadataString" {

			result = string(value.(MotadataString))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = strconv.Itoa(value.(int))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = strconv.Itoa(value.(MotadataINT).ToInt())

		} else if reflect.TypeOf(value).String() == "[]uint8" {

			result = string(value.([]uint8))

		} else if reflect.TypeOf(value).String() == "uint" {

			result = strconv.Itoa(ToInt(value.(uint)))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = strconv.Itoa(int(value.(float64)))

		} else if reflect.TypeOf(value).String() == "*string" {

			result = *value.(*string)

		}
	}

	return
}

func ToInt(value interface{}) (result int) {

	if reflect.TypeOf(value).Name() == "string" {

		output, _ := strconv.ParseInt(strings.TrimSpace(value.(string)), 10, 64)

		result = int(output)

	} else if reflect.TypeOf(value).Name() == "MotadataUINT" {

		result = int(value.(MotadataUINT))

	} else if reflect.TypeOf(value).Name() == "uint" {

		result = int(value.(uint))

	} else if reflect.TypeOf(value).Name() == "uint8" {

		result = int(value.(uint8))

	} else if reflect.TypeOf(value).Name() == "MotadataUINT16" {

		result = int(value.(MotadataUINT16))

	} else if reflect.TypeOf(value).Name() == "uint16" {

		result = int(value.(uint16))

	} else if reflect.TypeOf(value).Name() == "uint32" {

		result = int(value.(uint32))

	} else if reflect.TypeOf(value).Name() == "uint64" {

		result = int(value.(uint64))

	} else if reflect.TypeOf(value).Name() == "int" {

		result = value.(int)

	} else if reflect.TypeOf(value).Name() == "int8" {

		result = int(value.(int8))

	} else if reflect.TypeOf(value).Name() == "int16" {

		result = int(value.(int16))

	} else if reflect.TypeOf(value).Name() == "int32" {

		result = int(value.(int32))

	} else if reflect.TypeOf(value).Name() == "int64" {

		result = int(value.(int64))

	} else if reflect.TypeOf(value).Name() == "float64" {

		result = int(value.(float64))
	}

	return
}

func ToFloat64(value interface{}) (result float64) {

	if reflect.TypeOf(value).Name() == "string" {

		value, _ := strconv.ParseFloat(ToString(value), 64)

		result = value
	}

	return
}

func (value MotadataStringMap) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (context MotadataMap) GetUINT16Value(key string) (result MotadataUINT16) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "float64" {

			result = MotadataUINT16(context[key].(float64))
		} else if reflect.TypeOf(context[key]).Name() == "int" {

			result = MotadataUINT16(context[key].(int))
		} else if reflect.TypeOf(context[key]).Name() == "string" {

			output, _ := strconv.ParseInt(strings.TrimSpace(context[key].(string)), 10, 16)

			result = MotadataUINT16(output)
		} else if reflect.TypeOf(context[key]).Name() == "MotadataUINT16" {
			result = context[key].(MotadataUINT16)
		}
	}

	return
}

func (context MotadataMap) GetUINT32Value(key string) (result MotadataUINT32) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "MotadataUINT32" {

			result = context[key].(MotadataUINT32)

		}
	}

	return
}

func (context MotadataMap) GetUINT8Value(key string) (result MotadataUINT8) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT8(value.(float64))
		} else if reflect.TypeOf(context[key]).Name() == "int" {

			result = MotadataUINT8(context[key].(int))
		} else if reflect.TypeOf(context[key]).Name() == "string" {

			output, _ := strconv.ParseInt(strings.TrimSpace(context[key].(string)), 10, 8)

			result = MotadataUINT8(output)
		}
	}

	return
}

func (value MotadataMap) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (elements MotadataStringList) Contains(value MotadataString) bool {

	for _, element := range elements {

		if MotadataString(element) == value {

			return true
		}
	}
	return false
}

func (value MotadataStringList) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (value MotadataMapList) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (context MotadataMap) GetMapSliceValue(key string) (result []MotadataMap) {

	if context.Contains(key) {

		return context[key].([]MotadataMap)
	}

	return
}

func IsNotEmptyMapSlice(values []MotadataMap) bool {

	if values != nil && len(values) > 0 {

		return true
	}

	return false
}

func (context MotadataMap) GetStringSliceValue(key string) (result []interface{}) {

	if context.Contains(key) {

		return context[key].([]interface{})
	}

	return
}

func IsNotEmptyInterfaceSlice(values []interface{}) bool {

	if values != nil && len(values) > 0 {

		return true
	}

	return false
}

func IndexCheck(value string, key string) bool {

	if strings.Index(value, key) > 0 {

		return true
	}

	return false
}

func ToDate(value int64) string {

	return time.UnixMilli(value).Format(sdkconstant.TimeFormat)
}

func ToJson(value MotadataMap) MotadataString {

	result, err := json.Marshal(value)

	if err == nil {
		return MotadataString(result)
	}

	return sdkconstant.BlankString
}
