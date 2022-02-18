package main

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Validator map[string]interface{}

func (instance Validator) ValidateString(item interface{}) bool {
	if str, ok := item.(string); ok {
		if str != "" {
			return true
		}
	}

	return false
}

func (instance Validator) ValidateInteger(item interface{}) bool {
	switch t := item.(type) {
	case int64:
		return true

	case float32:
		return true

	case float64:
		return true

	case int:
		return true

	case string:
		_, err := strconv.ParseInt(t, 10, 64)
		return err == nil

	default:
		return false
	}
}

func (instance Validator) ValidateV4UUID(item interface{}) bool {
	if val, ok := item.(string); ok {
		r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
		return r.MatchString(val)
	} else {
		fmt.Print("Failed to parse string")
	}

	return false
}

func (instance Validator) validateReqBody(requestBody io.Reader) ([]string, map[string]interface{}) {
	body, _ := io.ReadAll(requestBody)
	jsonData := map[string]interface{}{}
	errs := []string{}

	if err := json.Unmarshal(body, &jsonData); err != nil {
		errs = append(errs, "Failed to unmarshal body to json.\n")
		return errs, nil
	}

	for key, validationType := range instance {
		if match, ok := jsonData[key]; ok {
			switch validationType {
			case "string":
				if !instance.ValidateString(match) {
					errs = append(errs, key+" is not of type string.\n")
				}
				break
			case "int":
				if !instance.ValidateInteger(match) {
					errs = append(errs, key+" is not of type int.\n")
				}
				break
			case "uuid":
				if !instance.ValidateV4UUID(match) {
					errs = append(errs, key+" is not of type uuid.\n")
				}
				break
			default:
				errs = append(errs, "Unknown validation type.\n")
			}
		} else {
			errs = append(errs, key+" was not found in the body.\n")
		}
	}

	return errs, jsonData
}
