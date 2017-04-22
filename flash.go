package flash

import (
	"bytes"
	"context"
	"encoding/gob"
	"net/url"
)

// Flash type
type Flash url.Values

func init() {
	gob.Register(&Flash{})
}

// New creates new flash
func New() Flash {
	return make(Flash)
}

// Decode decodes flash data
func Decode(b []byte) (Flash, error) {
	f := New()
	err := gob.NewDecoder(bytes.NewReader(b)).Decode(&f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type contextKey int

const flashKey contextKey = iota

// Get gets flash from context
func Get(ctx context.Context) Flash {
	f, ok := ctx.Value(flashKey).(Flash)
	if !ok {
		return New()
	}
	return f
}

// Set sets flash to context's value
func Set(ctx context.Context, f Flash) context.Context {
	return context.WithValue(ctx, flashKey, f)
}

// Encode encodes flash
func (f Flash) Encode() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(f)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Set sets value to flash
func (f Flash) Set(key, value string) {
	url.Values(f).Set(key, value)
}

// Get gets value from flash
func (f Flash) Get(key string) string {
	return url.Values(f).Get(key)
}

// Add adds value to flash
func (f Flash) Add(key, value string) {
	url.Values(f).Add(key, value)
}

// Del deletes key from flash
func (f Flash) Del(key string) {
	url.Values(f).Del(key)
}

// Clear deletes all data
func (f Flash) Clear() {
	for k := range f {
		url.Values(f).Del(k)
	}
}
