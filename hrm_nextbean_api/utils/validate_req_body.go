package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ErrorField struct {
	ErrType    string `json:"err-type"`
	Field      string `json:"field"`
	ErrMessage string `json:"err-message"`
}

type ValidateRequestBody struct {
	rule_dict  map[string]map[string]string
	req_body   map[string]interface{}
	list_error []ErrorField
}

func snakeCaseToCamelCase(str string) string {
	var result strings.Builder
	words := strings.Split(str, "-")
	for _, word := range words {
		firstChar := string(word[0])
		uppercaseFirst := strings.ToUpper(firstChar)
		new_word := uppercaseFirst + word[1:]
		result.WriteString(new_word)
	}
	return result.String()
}

func createNewReqBodyMap(req_body map[string]interface{}) map[string]interface{} {
	new_map := make(map[string]interface{})
	for key, value := range req_body {
		new_map[snakeCaseToCamelCase(key)] = value
	}
	return new_map
}

func parseInput(input string) map[string]string {
	result := make(map[string]string)
	parts := strings.Split(input, ",")
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		} else {
			if kv[0] == "required" {
				result[kv[0]] = "true"
			} else {
				result[kv[0]] = ""
			}
		}
	}
	return result
}

func getRuleDictionary(s interface{}) map[string]map[string]string {
	rule_dict := make(map[string]map[string]string)

	val := reflect.ValueOf(s)
	val = reflect.Indirect(val)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {

		field := typ.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		rule_dict[field.Name] = parseInput(tag)
	}
	return rule_dict
}

func CreateValidateRequestBody(req_body map[string]interface{}, s interface{}) *ValidateRequestBody {
	new_req_body := createNewReqBodyMap(req_body)
	new_err_list := []ErrorField{}
	return &ValidateRequestBody{
		req_body:   new_req_body,
		rule_dict:  getRuleDictionary(s),
		list_error: new_err_list,
	}
}

// ? 1. Check for valid fields (in req-body) (are there any field that not exist in struct ?)
func (vrb *ValidateRequestBody) validateReqBody_noNeedField() {
	for key := range vrb.req_body {
		if _, ok := vrb.rule_dict[key]; !ok {
			err_field := ErrorField{ErrType: "extra-field", Field: key, ErrMessage: fmt.Sprintf("field '%v' is invalid", key)}
			vrb.list_error = append(vrb.list_error, err_field)
		}
	}
}

// ? 2. Check required fields (in req-body)
func (vrb *ValidateRequestBody) validateReqBody_missingField() {
	for key, value := range vrb.rule_dict {
		if _, ok := vrb.req_body[key]; !ok {
			_, isRequired := value["required"]
			if isRequired {
				err_field := ErrorField{ErrType: "missing-field", Field: key, ErrMessage: fmt.Sprintf("missing field '%v'", key)}
				vrb.list_error = append(vrb.list_error, err_field)
			}
		}
	}
}

func isNumber(value interface{}) bool {
	numberTypes := []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	valueKind := reflect.ValueOf(value).Kind()
	for _, t := range numberTypes {
		if valueKind == t {
			return true
		}
	}
	return false
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@(gmail\.com|fpt\.edu\.vn)$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func extractEnum(enumString string) string {
	trimmed := strings.TrimPrefix(enumString, "enum(")
	trimmed = strings.TrimSuffix(trimmed, ")")
	return trimmed
}

func checkEnumValue(value interface{}, rule_value string) bool {
	valueStr, ok := value.(string)
	if !ok {
		return false
	}

	enum := extractEnum(rule_value)

	list_enum := strings.Split(enum, " or ")
	for i := 0; i < len(list_enum); i++ {
		if list_enum[i] == valueStr {
			return true
		}
	}

	return false
}

// ? 3. Check the validity of each field (in req-body)
func (vrb *ValidateRequestBody) checkType(key string, value interface{}, rule_value string) {
	switch rule_value {
	case "string":
		if _, ok := value.(string); !ok {
			err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: fmt.Sprintf("field '%v' must be a string", key)}
			vrb.list_error = append(vrb.list_error, err_field)
		}
	case "email":
		if v, ok := value.(string); ok {
			if is_email := isValidEmail(v); !is_email {
				err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: fmt.Sprintf("field '%v' must be a email (@gmail.com or @fpt.edu.vn)", key)}
				vrb.list_error = append(vrb.list_error, err_field)
			}
		}
	case "number":
		if ok := isNumber(value); !ok {
			err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: fmt.Sprintf("field '%v' must be a string", key)}
			vrb.list_error = append(vrb.list_error, err_field)
		}
	default:
		if strings.Contains(rule_value, "enum") {
			if !checkEnumValue(value, rule_value) {
				err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: fmt.Sprintf("field '%v' must be %v", key, extractEnum(rule_value))}
				vrb.list_error = append(vrb.list_error, err_field)
			}
		}
	}

}

func (vrb *ValidateRequestBody) checkMinMax(key string, value interface{}) {
	err_mess_min := ""
	err_mess_max := ""

	rule := vrb.rule_dict[key]
	vl_min, ok_min := rule["min"]
	vl_max, ok_max := rule["max"]

	if ok_min {
		value_min, _ := strconv.Atoi(vl_min)
		if value_string, ok := value.(string); ok {
			if len(value_string) < value_min {
				err_mess_min = fmt.Sprintf("%v must more than %v keyword", key, value_min)
			}
		} else if isNumber(value) {
			value_number := value.(int)

			if value_number < value_min {
				err_mess_min = fmt.Sprintf("%v must more than %v", key, value_min)
			}
		}
	}

	if ok_max {
		value_max, _ := strconv.Atoi(vl_max)
		if value_string, ok := value.(string); ok {
			if len(value_string) > value_max {
				err_mess_max = fmt.Sprintf("%v must less than %v keyword", key, value_max)
			}
		} else if isNumber(value) {
			value_number := value.(int)

			if value_number > value_max {
				err_mess_min = fmt.Sprintf("%v must less than %v", key, value_max)
			}
		}
	}

	if err_mess_min != "" {
		err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: err_mess_min}
		vrb.list_error = append(vrb.list_error, err_field)
		return
	} else if err_mess_max != "" {
		err_field := ErrorField{ErrType: "valid-field", Field: key, ErrMessage: err_mess_max}
		vrb.list_error = append(vrb.list_error, err_field)
		return
	}

}

func (vrb *ValidateRequestBody) validateReqBody_checkValidField() {
	for key, value := range vrb.req_body {
		rule := vrb.rule_dict[key]
		flag := false
		for rule_key, rule_value := range rule {
			switch rule_key {
			case "type":
				vrb.checkType(key, value, rule_value)
			case "min", "max":
				if !flag {
					vrb.checkMinMax(key, value)
					flag = true
				}
			}
		}
	}
}

// ? 4. get status
func (vrb *ValidateRequestBody) GetValidateStatus() (bool, []ErrorField) {
	vrb.validateReqBody_noNeedField()
	vrb.validateReqBody_missingField()
	vrb.validateReqBody_checkValidField()
	if len(vrb.list_error) == 0 {
		return true, []ErrorField{}

	}
	return false, vrb.list_error
}
