package money

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	m := Convert(2)
	tt := reflect.TypeOf(m)

	if tt != reflect.TypeOf(m) {
		t.Errorf("got: %d, want: %d.", tt, reflect.TypeOf(m))
	}

	if m != 200 {
		t.Errorf("got: %d, want: %d.", m, 200)
	}
}

func TestFloat64(t *testing.T) {
	m := Convert(2)
	f64 := reflect.TypeOf(m.Float64()).String()

	if f64 != "float64" {
		t.Errorf("got: %s, want: %s.", f64, "float64")
	}

	if m.Float64() != 2.0 {
		t.Errorf("got: %.2f, want: %.2f.", m.Float64(), 2.0)
	}
}

func TestMultiply(t *testing.T) {
	m := Convert(2).Multiply(2)

	if m != 400 {
		t.Errorf("got: %d, want: %d.", m, 400)
	}

	m = Convert(1).Multiply(.19)

	if m != 19 {
		t.Errorf("got: %d, want: %d.", m, 19)
	}
}

func TestFormat(t *testing.T) {
	m := Convert(19.99).Format()

	if m != "19.99" {
		t.Errorf("got: %v (%T), want: %s (string).", m, m, "19.99")
	}
}

func TestMarshal(t *testing.T) {
	m, _ := Convert(19.99).MarshalJSON()

	if string(m) != "19.99" {
		t.Errorf("got: %v (%T), want: %s (string).", string(m), string(m), "19.99")
	}
}

func TestUnmarshal(t *testing.T) {
	var m Money
	json := []byte("19.99")
	m.UnmarshalJSON(json)

	if m != 1999 {
		t.Errorf("got: %v (%T), want: %d (money.Money).", m, m, 1999)
	}
}
