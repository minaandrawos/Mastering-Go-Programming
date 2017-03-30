package Appliances

type Microwave struct{
	typeName string
}

func (mr *Microwave)Start(){
	mr.typeName = " Microwave "
}

func (mr *Microwave)GetPurpose() string{
	return "I am a " + mr.typeName + " I heat stuff up!!"
}

