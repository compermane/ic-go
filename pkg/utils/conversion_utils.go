package utils

import (
	"reflect"

	"github.com/compermane/ic-go/pkg/domain/functions"
)


func ArgToReflectValue(args []interface{}, is_variadic bool, fn *functions.Function) (r_values []reflect.Value) {
	for i, arg := range args {
		value := reflect.ValueOf(arg)

		if fn.ArgTypes[i].Kind().String() == "func" {
			// Caso em que o argumento se trata de uma func
			r_values = append(r_values, reflect.Zero(fn.ArgTypes[i]))
		} else if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
			// Caso em que o argumento não se trata de uma slice nem de um array
			r_values = append(r_values, value)
		} else {
			example_element := value.Index(0).Interface()

			// fmt.Println(reflect.TypeOf(example_element).String())
			switch reflect.TypeOf(example_element).String() {
			case "string":
				if i == len(args) - 1 && is_variadic {
					for i := 0; i < value.Len(); i++ {
						r_values = append(r_values, reflect.ValueOf(value.Index(i).Interface()))
					}
				} else {
					r_array := make([]string, value.Len())
					
					for i := 0; i < value.Len(); i++ {
						v := value.Index(i).Interface()
						str, _ := v.(string)
						r_array[i] = str
					}
					
					r_values = append(r_values, reflect.ValueOf(r_array))
				}
			case "int":
				// r_array := make([]int, 0)
			default:
				// r_array := make([]any, 0)

			}
		}
	}

	return r_values
}

func AnySliceToStringSlice(arg []any) (str_values []string) {
	str_values = make([]string, len(arg))
	for i, v := range arg {
		str, _ := v.(string)
		str_values[i] = str
	}

	return str_values
}