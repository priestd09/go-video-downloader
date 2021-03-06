/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * builtin-cast.go: Checked casts of interface{} to various types
 */

package runtime

import "fmt"
import "reflect"

// CastToInt casts to int, handles boxing, panics with extractorError
func CastToInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return v
	case OptInt:
		return v.Get()
	default:
		panic(NewCastError(val, "int"))
	}
}

// CastToString casts to string, handles boxing, panics with extractorError
func CastToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case OptString:
		return v.Get()
	default:
		panic(NewCastError(val, "string"))
	}
}

// CastToFloat casts to float, handles boxing, panics with extractorError
func CastToFloat(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case OptFloat:
		return v.Get()
	default:
		panic(NewCastError(val, "float64"))
	}
}

// CastToOptInt casts to OptInt, handles boxing, panics with extractorError
func CastToOptInt(val interface{}) OptInt {
	switch v := val.(type) {
	case nil:
		return OptInt{}
	case int:
		return AsOptInt(v)
	case OptInt:
		return v
	default:
		panic(NewCastError(val, "OptInt"))
	}
}

// CastToOptString casts to OptString, handles boxing, panics with extractorError
func CastToOptString(val interface{}) OptString {
	switch v := val.(type) {
	case nil:
		return OptString{}
	case string:
		return AsOptString(v)
	case OptString:
		return v
	default:
		panic(NewCastError(val, "OptString"))
	}
}

// CastToOptFloat casts to OptFloat, handles boxing, panics with extractorError
func CastToOptFloat(val interface{}) OptFloat {
	switch v := val.(type) {
	case nil:
		return OptFloat{}
	case float64:
		return AsOptFloat(v)
	case OptFloat:
		return v
	default:
		panic(NewCastError(val, "OptFloat"))
	}
}

// IsList checks if val is a list (slice)
func IsList(val interface{}) bool {
	return reflect.TypeOf(val).Kind() == reflect.Slice
}

// NewCastError creates extractorError's for failed casts
func NewCastError(val interface{}, destType string) error {
	return newExtractorError(fmt.Sprintf("Cannot cast %T to %s", val, destType))
}

// GetIntField tries to get an int typed value from dict, handles boxing, panics with extractorError
func GetIntField(dict SDict, field string, required bool, def int) int {
	if val, ok := dict[field]; ok {
		return CastToInt(val)
	} else if required {
		panic(newExtractorError(fmt.Sprintf("No %v found in the result dict", field)))
	} else {
		return def
	}
}

// GetStringField tries to get a string typed value from dict, handles boxing, panics with extractorError
func GetStringField(dict SDict, field string, required bool, def string) string {
	if val, ok := dict[field]; ok {
		return CastToString(val)
	} else if required {
		panic(newExtractorError(fmt.Sprintf("No %v found in the result dict", field)))
	} else {
		return def
	}
}

// GetFloatField tries to get a float typed value from dict, handles boxing, panics with extractorError
func GetFloatField(dict SDict, field string, required bool, def float64) float64 {
	if val, ok := dict[field]; ok {
		return CastToFloat(val)
	} else if required {
		panic(newExtractorError(fmt.Sprintf("No %v found in the result dict", field)))
	} else {
		return def
	}
}
