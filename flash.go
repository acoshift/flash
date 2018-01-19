package flash

import (
	"bytes"
	"encoding/gob"
)

// Flash type
type Flash struct {
	v       data
	changed bool
}

type data map[string][]interface{}

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
		return nil, err
	}
	return buf.Bytes(), nil
}

// Set sets value to flash
func (f *Flash) Set(key string, value interface{}) {
	if !f.changed {
		f.changed = true
	}
	if f.v == nil {
		f.v = make(data)
	}
	f.v[key] = []interface{}{value}
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

// Add adds value to flash
func (f *Flash) Add(key string, value interface{}) {
	if !f.changed {
		f.changed = true
	}
	if f.v == nil {
		f.v = make(data)
	}
	f.v[key] = append(f.v[key], value)
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

func cloneValues(src data) data {
	n := make(data, len(src))
	for k, vv := range src {
		n[k] = make([]interface{}, len(vv))
		for kk, vvv := range vv {
			n[k][kk] = vvv
		}
	}
	return n
}
