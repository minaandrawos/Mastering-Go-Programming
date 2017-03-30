package Appliances

// define a fridge struct, the struct contain a string representing the type name
type Fridge struct{
	typeName string
}

//The fridge struct implements the start() function
func (fr *Fridge)Start(){
	 fr.typeName = " Fridge "
}

//The fridge struct implements the start() function
func (fr *Fridge)GetPurpose() string{
	return "I am a " + fr.typeName + " I cool stuff down!!"
}
