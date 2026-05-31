package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, got T, want T) {

	t.Helper()

	if !isEqual(got, want) {
		t.Errorf("got %q; want %q", got, want)
	}
}

func True(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Errorf("got %q; want true", got)
	}
}

func Nil(t *testing.T, got any) {
	t.Helper()

	if !isNil(got) {
		t.Errorf("got %q; want nil", got)
	}
}

func isEqual[T any](a T, b T) bool {

	if isNil(a) && isNil(b) {
		return true
	}

	return reflect.DeepEqual(a, b)
}

func isNil(v any) bool {

	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)

	//generico no el especifico
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rv.IsNil()
	}

	return false
}
