package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	KC            float64 = 273.15
	AbsoluteZeroC Celsius = -273.15
	FreeZingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (f Kelvin) String() string {
	return fmt.Sprintf("%g°K", f)
}
