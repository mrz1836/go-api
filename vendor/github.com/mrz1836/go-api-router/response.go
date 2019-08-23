package apirouter

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/matryer/respond"
)

// AllowedKeys is for allowed keys
type AllowedKeys map[string]interface{}

// ReturnResponse helps return a status code and message to the end user
func ReturnResponse(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	respond.With(w, req, code, data)
}

// ReturnJSONEncode is a mixture of ReturnResponse and JSONEncode
func ReturnJSONEncode(w http.ResponseWriter, code int, e *json.Encoder, objects interface{}, allowed []string) (err error) {

	// Set the content if JSON
	w.Header().Set("Content-Type", "application/json")

	// Set the header status code
	w.WriteHeader(code)

	// Attempt to encode the objects
	err = JSONEncode(e, objects, allowed)

	return
}

// JSONEncodeHierarchy will execute JSONEncode for multiple nested objects
func JSONEncodeHierarchy(w io.Writer, objects interface{}, allowed interface{}) error {
	if slice, ok := allowed.([]string); ok {
		return JSONEncode(json.NewEncoder(w), objects, slice)
	} else if obj, ok := allowed.(AllowedKeys); ok {
		data := reflect.ValueOf(objects).Elem().Interface()
		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		numFields := t.NumField()
		_, _ = w.Write([]byte{'{'})
		for i := 0; i < numFields; i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if len(jsonTag) == 0 {
				jsonTag = field.Name
			}
			keys := obj[jsonTag]
			if keys != nil {
				_, _ = w.Write([]byte("\""))
				_, _ = w.Write([]byte(jsonTag))
				_, _ = w.Write([]byte("\": "))
				err := JSONEncodeHierarchy(w, v.Field(i).Interface(), keys)
				if err != nil {
					return err
				}
				if i != numFields-1 {
					_, _ = w.Write([]byte{','})
				}
			}
		}
		_, _ = w.Write([]byte{'}'})
	}
	return nil
}

// JSONEncodeModels will encode only the allowed fields of the models
func JSONEncode(e *json.Encoder, objects interface{}, allowed []string) error {
	var data []map[string]interface{}
	isMulti := false
	count := 0

	if reflect.TypeOf(objects).Kind() == reflect.Slice {
		count = reflect.ValueOf(objects).Len()
		data = make([]map[string]interface{}, count)
		isMulti = true
	}

	if isMulti {
		if count == 0 {
			return e.Encode(make([]interface{}, 0))
		}

		raw := reflect.ValueOf(objects)

		obj := jsonMap(raw.Index(0).Interface())
		toRemove := make([]string, 0)

		for k := range obj {
			if FindString(k, allowed) == -1 {
				toRemove = append(toRemove, k)
			}
		}

		for _, k := range toRemove {
			delete(obj, k)
		}

		data[0] = obj

		for i := 1; i < count; i++ {
			obj = jsonMap(raw.Index(i).Interface())

			for _, k := range toRemove {
				delete(obj, k)
			}

			data[i] = obj
		}

		return e.Encode(data)
	}

	obj := jsonMap(objects)
	toRemove := make([]string, 0)

	for k := range obj {
		if FindString(k, allowed) == -1 {
			toRemove = append(toRemove, k)
		}
	}

	for _, k := range toRemove {
		delete(obj, k)
	}

	return e.Encode(obj)
}

// jsonMap converts an object to a map of string interfaces
func jsonMap(obj interface{}) map[string]interface{} {
	fieldValues := make(map[string]interface{}, 0)

	var s, stringPointer reflect.Value

	// Dereference the obj if it is a pointer
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		stringPointer = reflect.ValueOf(obj)
		s = stringPointer.Elem()
	} else {
		s = reflect.ValueOf(obj)
		stringPointer = reflect.ValueOf(&obj)
	}

	typeOfT := s.Type()
	for i := 0; i < typeOfT.NumField(); i++ {
		structField := typeOfT.Field(i)
		fieldName := structField.Name
		if fieldName[0] != strings.ToUpper(string(fieldName[0]))[0] {
			continue
		}

		// Exclude any field starting with an underscore
		if strings.Index(fieldName, "_") == 0 {
			continue
		}
		val := s.Field(i)
		// Check for embedded types
		if structField.Anonymous {
			subFields := jsonMap(val.Interface())
			for k, v := range subFields {
				fieldValues[k] = v
			}
			continue
		}
		key := SnakeCase(fieldName)
		comps := strings.Split(key, ",")
		key = comps[0]
		fieldType := structField.Type
		if fieldType.Kind() != reflect.Ptr && val.CanAddr() {
			fieldType = reflect.PtrTo(fieldType)
			val = val.Addr()
		}
		fieldValues[key] = val.Interface()
	}

	return fieldValues
}
