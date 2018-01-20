package flash

import (
	"bytes"
	"encoding/gob"
)

// Flash type
type Flash struct {
	v       Data
	changed bool
}

// Data is the flash data
type Data map[string][]interface{}

// New creates new flash
func New() *Flash {
	return &Flash{}
}

// Decode decodes flash data
func Decode(b []byte) (*Flash, error) {
	f := New()
	if len(b) == 0 {
		return f, nil
	}

	err := gob.NewDecoder(bytes.NewReader(b)).Decode(&f.v)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Encode encodes flash
func (f *Flash) Encode() ([]byte, error) {
	// empty flash encode to empty bytes
	if len(f.v) == 0 {
		return []byte{}, nil
	}

	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(f.v)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// Values returns flash's data
func (f *Flash) Values() Data {
	return f.v
}

// Set sets value to flash
func (f *Flash) Set(key string, value interface{}) {
	if !f.changed {
		f.changed = true
	}
	if f.v == nil {
		f.v = make(Data)
	}
	f.v[key] = []interface{}{value}
}

// Add adds value to flash
func (f *Flash) Add(key string, value interface{}) {
	if !f.changed {
		f.changed = true
	}
	if f.v == nil {
		f.v = make(Data)
	}
	f.v[key] = append(f.v[key], value)
}

// Get gets value from flash
func (f *Flash) Get(key string) interface{} {
	if f.v == nil {
		return nil
	}
	if len(f.v[key]) == 0 {
		return nil
	}
	return f.v[key][0]
}

func (f *Flash) Value(key string) []interface{} {
	if f.v == nil {
		return []interface{}{}
	}
	if f.v[key] == nil {
		return []interface{}{}
	}
	return f.v[key]
}

// GetString gets string from flash
func (f *Flash) GetString(key string) string {
	r, _ := f.Get(key).(string)
	return r
}

// GetInt gets int from flash
func (f *Flash) GetInt(key string) int {
	r, _ := f.Get(key).(int)
	return r
}

// GetInt64 gets int64 from flash
func (f *Flash) GetInt64(key string) int64 {
	r, _ := f.Get(key).(int64)
	return r
}

// GetFloat32 gets float32 from flash
func (f *Flash) GetFloat32(key string) float32 {
	r, _ := f.Get(key).(float32)
	return r
}

// GetFloat64 gets float64 ffrom flash
func (f *Flash) GetFloat64(key string) float64 {
	r, _ := f.Get(key).(float64)
	return r
}

// Del deletes key from flash
func (f *Flash) Del(key string) {
	if !f.Has(key) {
		return
	}
	if !f.changed {
		f.changed = true
	}
	f.v[key] = nil
}

// Has checks is flash has a given key
func (f *Flash) Has(key string) bool {
	if f.v == nil {
		return false
	}
	return len(f.v[key]) > 0
}

// Clear deletes all data
func (f *Flash) Clear() {
	if f.Count() > 0 {
		f.changed = true
	}
	f.v = nil
}

// Count returns count of flash's keys
func (f *Flash) Count() int {
	return len(f.v)
}

// Clone clones flash
func (f *Flash) Clone() *Flash {
	return &Flash{v: cloneValues(f.v)}
}

// Changed returns true if value changed
func (f *Flash) Changed() bool {
	return f.changed
}

func cloneValues(src Data) Data {
	n := make(Data, len(src))
	for k, vv := range src {
		n[k] = make([]interface{}, len(vv))
		for kk, vvv := range vv {
			n[k][kk] = vvv
		}
	}
	return n
}
