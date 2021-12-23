package main

import (
	"fmt"

	conversion "github.com/galigaribaldi/Conversions/Conversions"
)

func main() {
	//Longitud
	//c := conversion.Leng
	//fmt.Println(c.GetFeatInMetter())
	fmt.Println("Conversion de 12 Yardas a Metros:", conversion.YardToMetter(12))
	fmt.Println("Conversion de 12 Toneladas a Kilogramos:", conversion.PundToKilograms(12))

}
