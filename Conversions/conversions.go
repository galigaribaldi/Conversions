package conversions

//Variable for Length Units like Yards, Mille and feat
var Leng = Length{}

//Variable for Mass Units like Pound, Ounce and ton
var Mas = Mass{}

/*
---------------------------------------
---------- Length Conversions ---------
---------------------------------------
*/
//Convert Yards to Metter
func YardToMetter(yards float32) float32 {
	newYards := yards * Leng.GetYardInMetter()
	return newYards
}

//Convert Yards to Centimetter
func YardToCentimetter(yards float32) float32 {
	newYards := yards * Leng.GetYardInCentimetter()
	return newYards
}

//Convert Yards to Kilometter
func YardToKilometter(yards float32) float32 {
	newYards := yards * Leng.GetYardInKilometter()
	return newYards
}

//Convert Feat to Metter
func FeatToMetter(feat float32) float32 {
	newFeat := feat * Leng.GetFeatInMetter()
	return newFeat
}

//Convert Feat to Centimetter
func FeatToCentimetter(feat float32) float32 {
	newFeat := feat * Leng.GetFeatInCentimetter()
	return newFeat
}

//Convert Feat to Kilometter
func FeatToKilometter(feat float32) float32 {
	newFeat := feat * Leng.GetFeatInKilometter()
	return newFeat
}

//Convert Mille to Metter
func MilleToMetter(mille float32) float32 {
	newMille := mille * Leng.GetMilleInMetter()
	return newMille
}

//Convert Mille to Centimetter
func MilleToCentimetter(mille float32) float32 {
	newMille := mille * Leng.GetMilleInCentiMetter()
	return newMille
}

//Convert Mille to Kilometter
func MilleToKilometter(mille float32) float32 {
	newMille := mille * Leng.GetMilleInKilometter()
	return newMille
}

//Convert Inche to Metter
func IncheToMetter(inche float32) float32 {
	newInche := inche * Leng.GetIncheInMetter()
	return newInche
}

//Convert Inche to Centimetter
func IncheToCentimetter(inche float32) float32 {
	newInche := inche * Leng.GetIncheInCentimetter()
	return newInche
}

//Convert Inche to Kilometter
func IncheToKilometter(inche float32) float32 {
	newInche := inche * Leng.GetIncheInKilometter()
	return newInche
}

/*
---------------------------------------
---------- Mass Conversions -----------
---------------------------------------
*/

//Convert Pound to Kilograms
func PundToKilograms(pounds float32) float32 {
	newPounds := pounds * Mas.GetPoundInKg()
	return newPounds
}

//Convert Pound to Grams
func PundToGrams(pounds float32) float32 {
	newPounds := pounds * Mas.GetPoundIng()
	return newPounds
}

//Convert Pound to Miligrams
func PundToMiligrams(pounds float32) float32 {
	newPounds := pounds * Mas.GetPoundInmlg()
	return newPounds
}

//Convert Ounce to Kilograms
func OunceToKilograms(ounce float32) float32 {
	newOunce := ounce * Mas.GetOunceInKg()
	return newOunce
}

//Convert Ounce to Grams
func OunceToGrams(ounce float32) float32 {
	newOunce := ounce * Mas.GetOunceIng()
	return newOunce
}

//Convert Ounce to Miligrams
func OunceToMiligrams(ounce float32) float32 {
	newOunce := ounce * Mas.GetOunceInmlg()
	return newOunce
}

//Convert Ton to Kilograms
func TonToKilograms(ton float32) float32 {
	NewTon := ton * Mas.GetOunceInKg()
	return NewTon
}

//Convert Ton to Grams
func TonToGrams(ton float32) float32 {
	NewTon := ton * Mas.GetOunceIng()
	return NewTon
}

//Convert Ton to Miligrams
func TonToMiligrams(ton float32) float32 {
	NewTon := ton * Mas.GetOunceInmlg()
	return NewTon
}
