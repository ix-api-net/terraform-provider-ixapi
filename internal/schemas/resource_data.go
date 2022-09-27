package schemas

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceSetter is an interface, which resource data
// and flatable resource data implements.
type ResourceSetter interface {
	Set(key string, value interface{}) error
}

// FlatResource implements the resource setter and can be
// used to retrieve flattened resource data by applying
// the model specific ...SetResourceData function.
type FlatResource struct {
	res map[string]interface{}
}

// NewFlatResource creates a new flat resouce
func NewFlatResource() *FlatResource {
	return &FlatResource{
		res: make(map[string]interface{}),
	}
}

// Set implements the ResourceSetter interface
func (r *FlatResource) Set(key string, value interface{}) error {
	r.res[key] = value
	return nil
}

// Flatten returns the flattened resource
func (r *FlatResource) Flatten() map[string]interface{} {
	return r.res
}

// ResourceData extends the terraform resource data with
// additional typed getters.
type ResourceData struct {
	*schema.ResourceData
}

// GetString retrievs a string from the resource
func (res *ResourceData) GetString(key string) string {
	return res.Get(key).(string)
}

// GetStringOpt gets an optional string
func (res *ResourceData) GetStringOpt(key string) *string {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	ret := val.(string)
	if ret == "" {
		return nil
	}
	return &ret
}

// GetInt retrieves an integer from the resource
func (res *ResourceData) GetInt(key string) int {
	return res.Get(key).(int)
}

// GetIntOpt retrievs an optional integer from the resource
func (res *ResourceData) GetIntOpt(key string) *int {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	v := val.(int)
	return &v
}

// GetBool retrievs a boolean
func (res *ResourceData) GetBool(key string) bool {
	return res.Get(key).(bool)
}

// GetBoolOpt retrievs an optional boolean
func (res *ResourceData) GetBoolOpt(key string) *bool {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	ret := val.(bool)
	return &ret
}

// GetStringList gets list of strings from the resource data
func (res *ResourceData) GetStringList(key string) []string {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	return MustStringListFromAny(val)
}

// GetIntList gets a list of integers from resource data
func (res *ResourceData) GetIntList(key string) []int {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	return MustIntListFromAny(val)
}

// GetResource retrievs an embedded resource
func (res *ResourceData) GetResource(key string) Resource {
	val, ok := res.GetOk(key)
	if !ok {
		return nil
	}
	resVals := val.([]interface{})
	if len(resVals) == 0 {
		return nil
	}
	sres := Resource(resVals[0].(map[string]interface{})) // first embedded
	return sres
}

// Resource is an embedded resource
type Resource map[string]interface{}

// GetString retrievs a string from the resource
func (res Resource) GetString(key string) string {
	return res[key].(string)
}

// GetStringOpt gets an optional string
func (res Resource) GetStringOpt(key string) *string {
	val, ok := res[key]
	if !ok {
		return nil
	}
	sval := val.(string)
	if sval == "" {
		return nil
	}
	return &sval
}

// GetStringOptDefault returns a string pointer to an optional string
func (res Resource) GetStringOptDefault(key string, def string) *string {
	var s string
	val, ok := res[key]
	if !ok {
		s = def
	} else {
		s = val.(string)
	}
	return &s
}

// GetIntOpt returns an optional integer if present
func (res Resource) GetIntOpt(key string) *int {
	val, ok := res[key]
	if !ok {
		return nil
	}
	ival := val.(int)
	return &ival
}

// GetInt get an integer if present
func (res Resource) GetInt(key string) int {
	val, ok := res[key]
	if !ok {
		return 0
	}
	return val.(int)
}

// GetResource retrievs an embedded embedded resource
func (res Resource) GetResource(key string) Resource {
	val, ok := res[key]
	if !ok {
		return nil
	}
	resVals := val.([]interface{})
	if len(resVals) == 0 {
		return nil
	}
	sres := Resource(resVals[0].(map[string]interface{})) // first embedded
	return sres
}

// Timestamp returns the current unix timestamp
func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// MustStringListFromAny will convert an interface to a list of strings,
// interface must be any -> []any -> []string
func MustStringListFromAny(in any) []string {
	list := in.([]any)
	result := make([]string, len(list))
	for i, elem := range list {
		result[i] = elem.(string)
	}
	return result
}

// MustIntListFromAny will convert an interface to a list of ints,
// interface must be any -> []any -> []int
func MustIntListFromAny(in any) []int {
	list := in.([]any)
	result := make([]int, len(list))
	for i, elem := range list {
		result[i] = elem.(int)
	}
	return result
}
