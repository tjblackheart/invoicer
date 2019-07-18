package money

import (
	"fmt"
	"strconv"
)

// Money represents a currency value
type Money int64

// Convert converts a float64 to Money
func Convert(f64 float64) Money {
	return Money((f64 * 100) + 0.5)
}

// Float64 converts Money to float64
func (m Money) Float64() float64 {
	return float64(m) / 100
}

// Multiply safely multiplies a Money value by a float64, rounding to the nearest cent.
func (m Money) Multiply(f float64) Money {
	f64 := (float64(m) * f) + 0.5

	return Money(f64)
}

// Format returns a formatted Money value
func (m Money) Format() string {
	f64 := float64(m)

	return fmt.Sprintf("%.2f", f64/100)
}

// MarshalJSON implements the marshalling interface
func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(m.Format()), nil
}

// UnmarshalJSON implements the unmarshalling interface
func (m *Money) UnmarshalJSON(b []byte) error {
	f64, err := strconv.ParseFloat(string(b), 64)

	*m = Convert(f64)

	return err
}
