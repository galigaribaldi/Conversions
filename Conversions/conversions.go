package conversions

type Length struct {
	//Yards in Metters, Kilometters and Centimeters
	yardInMetter      float32
	yardInKilometer   float32
	yardInCentimeters float32

	//Feats in Metters, Kilometters and Centimeters
	featInKilometter  float32
	featInMetter      float32
	featInCentimetter float32

	//Milles in Metters, Kilometters and Centimeters
	milleInMetter      float32
	milleInKilometter  float32
	milleInCentimetter float32

	//Inche in Metters, Kilometters and Centimeters
	incheInKilometter  float32
	incheInMetter      float32
	incheInCentimetter float32
}

func NewLength() *Length {
	return &Length{
		yardInMetter:      0.9144,
		yardInKilometer:   0.0009144,
		yardInCentimeters: 91.44,

		featInKilometter:  0.0003048,
		featInMetter:      0.3048,
		featInCentimetter: 30.48,

		milleInMetter:      1609.34,
		milleInKilometter:  1.60934,
		milleInCentimetter: 160934,

		incheInKilometter:  0.0000254,
		incheInMetter:      0.0254,
		incheInCentimetter: 2.54,
	}
}
func (yard *Length) GetYardsInKilometter() float32 {
	return yard.yardInKilometer

}
