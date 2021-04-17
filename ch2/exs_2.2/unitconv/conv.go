package unitconv

func C2F(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func F2C(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func M2Ft(m Meter) Feet         { return Feet(m / 0.3048) }
func Ft2M(ft Feet) Meter        { return Meter(ft * 0.3048) }
func Kg2Lbs(kg Kilogram) Pound  { return Pound(kg / 0.45359237) }
func Lbs2Kg(lbs Pound) Kilogram { return Kilogram(lbs * 0.45359237) }
