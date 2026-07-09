package apify

import (
	"reflect"
	"strings"
	"sync"
)

// knownKeysCache memoizes the JSON key sets per struct type so the reflection cost is
// paid once per type rather than on every unmarshal.
var knownKeysCache sync.Map // map[reflect.Type]map[string]struct{}

// knownJSONKeys returns the set of JSON object keys that the struct type of v maps to its
// fields. Used by the models' UnmarshalJSON to decide which incoming keys are "extra"
// (unmodelled) and should be collected into the Extra map.
//
// The argument is a zero value of the model's alias type; only its type is inspected.
func knownJSONKeys(v any) map[string]struct{} {
	t := reflect.TypeOf(v)
	if cached, ok := knownKeysCache.Load(t); ok {
		return cached.(map[string]struct{})
	}

	keys := make(map[string]struct{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue // e.g. the Extra field itself
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			name = field.Name
		}
		keys[name] = struct{}{}
	}

	knownKeysCache.Store(t, keys)
	return keys
}
