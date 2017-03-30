package Appliances

// define a stove struct, the struct contain a string representing the type name
type Stove struct{
	typeName string
}

//The stove struct implements the start() function
func (sv *Stove)Start(){
	sv.typeName = " Stove "
}

//The stove struct implements the GetPurpose() function
func (sv *Stove)GetPurpose() string{
	return "I am a " + sv.typeName + " I cook food!!"
}

