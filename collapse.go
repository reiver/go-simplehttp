package simplehttp


import (
	"reflect"
)


// collapse turns different kinds of data into a map[string]interface{}.
// (Include structs.)
func collapse(tagName string, cascade ...interface{}) map[string]interface{} {

	// Deal with datum.
	data := make(map[string]interface{})


	// Collapse.
	for _,x := range cascade {
		switch xx := x.(type) {
		case map[string]interface{}:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]string:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]int:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]int8:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]int16:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]int32:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]int64:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]uint:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]uint8:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]uint16:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]uint32:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]uint64:
			for key, value := range xx {
				data[key] = value
			}
		case map[string]bool:
			for key, value := range xx {
				data[key] = value
			}
		case string:
			data["text"] = xx
		default:
			reflectedValue := reflect.ValueOf(xx)

			moreData := collapseReflectedValue(tagName, reflectedValue)

			for key, value := range moreData {
				data[key] = value
			}
		}
	}


	// Return.
	return data
}


func collapseReflectedValue(tagName string, reflectedValue reflect.Value) map[string]interface{} {

	// Initialize.
	data := make(map[string]interface{})


	// Collapse.
	switch reflectedValue.Kind() {
	case reflect.Struct:
		moreData := collapseReflectedStruct(tagName, reflectedValue)
		for k,v := range moreData {
			data[k] = v
		}
	case reflect.Ptr:
		moreData := collapseReflectedPtr(tagName, reflectedValue)
		for k,v := range moreData {
			data[k] = v
		}
	default:
		// Nothing here.
	}


	// Return.
	return data
}


func collapseReflectedStruct(tagName string, reflectedStruct reflect.Value) map[string]interface{} {

	// Initialize.
	data := make(map[string]interface{})


	//
	typeOfValue := reflectedStruct.Type()

	numFields := reflectedStruct.NumField()
	for i:=0; i<numFields; i++ {
		fieldValue := reflectedStruct.Field(i)

		fieldType := typeOfValue.Field(i)
		fieldTypeTag := fieldType.Tag
		key := ""
		if "" != tagName {
			key = fieldTypeTag.Get(tagName)
		}
		if "" == key {
			key = fieldType.Name
		}

		value := fieldValue.Interface()


		data[key] = value
	}


	// Return.
	return data
}


func collapseReflectedPtr(tagName string, reflectedPtr reflect.Value) map[string]interface{} {

	reflectedValue := reflectedPtr.Elem()

	data := collapseReflectedValue(tagName, reflectedValue)

	return data
}
