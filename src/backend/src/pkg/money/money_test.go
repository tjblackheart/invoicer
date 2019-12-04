package money

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		in  float64
		out Money
	}{
		{0, Money(0)},
		{1, Money(100)},
		{2.0, Money(200)},
		{3.1415, Money(314)},
		{3.1498, Money(315)},
		{-3.1498, Money(-315)},
	}

	for _, tt := range tests {
		m := Convert(tt.in)
		if m != tt.out {
			t.Errorf("have: %v (%T), want: %v (%T)", m, m, tt.out, tt.out)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		in  Money
		out float64
	}{
		{Money(200), 2.0},
		{Money(1999), 19.99},
		{Money(-1999), -19.99},
	}

	for _, tt := range tests {
		f64 := tt.in.Float64()
		if f64 != tt.out {
			t.Errorf("have: %v (%T), want: %v (%T)", f64, f64, tt.out, tt.out)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		in   Money
		mult float64
		out  Money
	}{
		{Money(200), 2.0, Money(400)},
		{Money(1000), .19, Money(190)},
		{Money(1999), .21, Money(420)},
		{Money(-1999), .19, Money(-380)},
	}

	for _, tt := range tests {
		m := tt.in.Multiply(tt.mult)
		if m != tt.out {
			t.Errorf("have %v (%T), want %v (%T)", m, m, tt.out, tt.out)
		}
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		in  Money
		out string
	}{
		{Money(200), "2.00"},
		{Money(9999), "99.99"},
		{Money(-1000), "-10.00"},
	}

	for _, tt := range tests {
		s := tt.in.Format()
		if s != tt.out {
			t.Errorf("have %v (%T), want %v (%T)", s, s, tt.out, tt.out)
		}
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
