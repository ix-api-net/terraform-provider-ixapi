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
