package reflectutil

import (
	"bytes"
	"reflect"
)

// MissingFields can be used to ensure that all fields in a struct are populated
// unexported fields are ignored
// The struct is checked recursively ensuring that structs referenced
// or embedded in the root struct also have all relevant fields populated.
// A field is considered populated if
// It is primitive, int/bool/float etc and is non-zero
// It is a string/map/slice/array and is not empty i.e. len(x) != 0
// It is a pointer and is not nil
// It is an interface and is not nil and does not wrap a nil pointer
// i must be a pointer to a struct, function panics on any other type
func FindMissingFields(val interface{}) *MissingFields {
	f := &MissingFields{}
	f.findR(reflect.ValueOf(val), getRootPath(val), false)
	return f
}

// Same as above but fields with `json:"-"` annoations are ignored
func FindMissingFieldsJson(val interface{}) *MissingFields {
	f := &MissingFields{}
	f.findR(reflect.ValueOf(val), getRootPath(val), true)
	return f
}

type MissingFields struct {
	MissingFields []string
}

func (f *MissingFields) findR(v reflect.Value, path string, ignoreIgnoredJson bool) {
	switch v.Kind() {
	case reflect.Bool:
		if !v.Bool() {
			f.Add(path)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == int64(0) {
			f.Add(path)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Uint() == uint64(0) {
			f.Add(path)
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == float64(0) {
			f.Add(path)
		}
	case reflect.Complex64, reflect.Complex128:
		if v.Complex() == complex128(0) {
			f.Add(path)
		}
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() == 0 {
			f.Add(path)
		}
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			f.Add(path)
		} else {
			f.findR(v.Elem(), path, ignoreIgnoredJson)
		}
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			tf := t.Field(i)
			if ignoreIgnoredJson && tf.Tag.Get("json") == "-" {
				continue
			}
			vf := v.Field(i)
			if vf.CanSet() {
				f.findR(vf, path+"."+tf.Name, ignoreIgnoredJson)
			}
		}
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
	}
}

func (f *MissingFields) Add(path string) {
	f.MissingFields = append(f.MissingFields, path)
}

func (c *MissingFields) Remove(path string) {
	for i, p := range c.MissingFields {
		if path == p {
			c.MissingFields = append(c.MissingFields[:i], c.MissingFields[i+1:]...)
		}
	}
}

func (c *MissingFields) String() string {
	if len(c.MissingFields) == 0 {
		return "No missing fields"
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("The following fields are not populated in the test fixture:\n")
	for _, path := range c.MissingFields {
		buf.WriteString(path)
		buf.WriteByte('\n')
	}
	return buf.String()
}

func getRootPath(val interface{}) string {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr {
		panic("Argument must be a pointer to a struct")
	}
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		panic("Argument must be a pointer to a struct")
	}
	return "*" + v.Type().Name()
}
