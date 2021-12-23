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

//Return yards in Kilometter
func (yard *Length) GetYardInKilometter() float32 {
	yard.yardInKilometer = 0.0009144
	return yard.yardInKilometer
}

//Return yards in metter
func (yard *Length) GetYardInMetter() float32 {
	yard.yardInMetter = 0.9144
	return yard.yardInMetter
}

func GetYardInMetter2() float32 {
	return 9.2
}

//Return yards in Centimetter
func (yard *Length) GetYardInCentimetter() float32 {
	yard.yardInCentimeters = 91.44
	return yard.yardInCentimeters
}

//Return Inche in Kilometter
func (inche *Length) GetIncheInKilometter() float32 {
	inche.incheInKilometter = 0.0000254
	return inche.incheInKilometter
}

//Return Inche in Metter
func (inche *Length) GetIncheInMetter() float32 {
	inche.incheInMetter = 0.0254
	return inche.incheInMetter
}

//Return Inche in Centimetter
func (inche *Length) GetIncheInCentimetter() float32 {
	inche.incheInCentimetter = 2.54
	return inche.incheInCentimetter
}

//Return Feat in Metter
func (feat *Length) GetFeatInKilometter() float32 {
	feat.featInKilometter = 0.0003048
	return feat.featInKilometter
}

//Return Feat in Metter
func (feat *Length) GetFeatInMetter() float32 {
	feat.featInMetter = 0.3048
	return feat.featInMetter
}

//Return Feat in Centimetter
func (feat *Length) GetFeatInCentimetter() float32 {
	feat.featInCentimetter = 30.48
	return feat.featInCentimetter
}

//Return Mille in KiloMetter
func (mille *Length) GetMilleInKilometter() float32 {
	mille.milleInMetter = 1609.34
	return mille.milleInKilometter
}

//Return Mille in Metter
func (mille *Length) GetMilleInMetter() float32 {
	mille.milleInKilometter = 1.60934
	return mille.milleInMetter
}

//Return Mille in centimetter
func (mille *Length) GetMilleInCentiMetter() float32 {
	mille.milleInCentimetter = 160934
	return mille.milleInCentimetter
}
