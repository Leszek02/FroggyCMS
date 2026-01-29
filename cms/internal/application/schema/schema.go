package schema

import (
	"reflect"
	"strings"
)

type ColumnSchema struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
	IsNullable string `json:"is_nullable"`

	InputType   string `json:"input_type"`
	OptionsFrom string `json:"options_from"`
}

func GenerateSchema(model any) []map[string]any {
	var result []map[string]any
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var processType func(reflect.Type)
	processType = func(typ reflect.Type) {
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)

			if field.Anonymous {
				processType(field.Type)
				continue
			}

			name := resolveFieldName(field)

			jsonTag := field.Tag.Get("json")
			if jsonTag == "-" {
				if field.Tag.Get("form") == "-" {
					continue
				}
			}

			dataType := resolveDataType(field.Type)
			inputType := resolveInputType(field, dataType)
			required := !isNullable(field)

			if inputType == "-" {
				continue
			}

			fieldSchema := map[string]any{
				"name":         name,
				"type":         dataType,
				"input":        inputType,
				"required":     required,
				"options_from": field.Tag.Get("options"),
				"label_field":  field.Tag.Get("labelField"),
				"value_field":  field.Tag.Get("valueField"),
			}
			result = append(result, fieldSchema)
		}
	}

	processType(t)
	return result
}

func getColumnName(gormTag, fieldName string) string {
	parts := strings.Split(gormTag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return toSnakeCase(fieldName)
}

func getDataType(t reflect.Type) string {

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "numeric"
	case reflect.Bool:
		return "boolean"
	default:
		return "string"
	}
}

func toSnakeCase(s string) string {
	var result []rune
	prevLower := false
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			if prevLower {
				result = append(result, '_')
			}
			result = append(result, r+'a'-'A')
			prevLower = false
		} else if r == '_' {
			result = append(result, '_')
			prevLower = false
		} else {
			result = append(result, r)
			prevLower = true
		}
	}
	return string(result)
}

func resolveFieldName(field reflect.StructField) string {
	if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
		return strings.Split(jsonTag, ",")[0]
	}

	if gormTag := field.Tag.Get("gorm"); gormTag != "" {
		for _, part := range strings.Split(gormTag, ";") {
			if strings.HasPrefix(part, "column:") {
				return strings.TrimPrefix(part, "column:")
			}
		}
	}

	return toSnakeCase(field.Name)
}

func resolveDataType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64, reflect.Uint16:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "numeric"
	default:
		return "string"
	}
}

func resolveInputType(field reflect.StructField, dataType string) string {
	formTag := field.Tag.Get("form")

	if formTag == "-" {
		return "-"
	}

	if formTag == "image" {
		return "file"
	}
	if formTag == "select" {
		return "select"
	}
	if field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Struct {
		return "-"
	}

	switch dataType {
	case "boolean":
		return "checkbox"
	case "integer", "numeric":
		return "number"
	default:
		return "text"
	}
}

func isNullable(field reflect.StructField) bool {
	if field.Type.Kind() == reflect.Ptr {
		return true
	}
	return field.Tag.Get("nullable") == "true"
}

func isIDField(name string) bool {
	n := strings.ToLower(name)
	return n == "id" || strings.HasSuffix(n, "_id")
}
