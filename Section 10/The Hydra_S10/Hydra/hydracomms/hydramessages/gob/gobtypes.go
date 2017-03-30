package hydragob

type CrewMember struct {
	ID           int32
	Name         string
	SecClearance int32
	Position     string
}

type Ship struct {
	Shipname    string
	CaptainName string
	Crew        []CrewMember
}

