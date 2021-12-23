package conversions

type Mass struct {
	//Pounds in Kilograms, Grams and Miligrams
	poundInKg, poundIng, poundInmlg float32
	//Ounce in Kilograms, Grams and Miligrams
	ounceInKg, ounceIng, ounceInmlg float32
	//Ton in Kilograms, Grams and Miligrams
	tonInKg, tonIng, tonInmlg float32
}

func (pound *Mass) GetPoundInKg() float32 {
	pound.poundInKg = 0.454
	return pound.poundInKg
}
func (pound *Mass) GetPoundIng() float32 {
	pound.poundIng = 454
	return pound.poundIng
}
func (pound *Mass) GetPoundInmlg() float32 {
	pound.poundInmlg = 454000
	return pound.poundInmlg
}

func (ounce *Mass) GetOunceInKg() float32 {
	ounce.ounceInKg = 0.0283
	return ounce.ounceInKg
}
func (ounce *Mass) GetOunceIng() float32 {
	ounce.ounceIng = 28.3
	return ounce.ounceIng
}
func (ounce *Mass) GetOunceInmlg() float32 {
	ounce.ounceInmlg = 28300
	return ounce.ounceInmlg
}

func (ton *Mass) GetTonInKg() float32 {
	ton.tonInKg = 907
	return ton.tonInKg
}

func (ton *Mass) GetTonIng() float32 {
	ton.tonIng = 907000
	return ton.tonIng
}

func (ton *Mass) GetTonInmlg() float32 {
	ton.tonInmlg = 907000000
	return ton.tonInmlg
}
