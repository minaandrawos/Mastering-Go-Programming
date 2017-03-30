package hydradblayer

import (
	"testing"
)

func BenchmarkMySQLDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mysql", "gouser:gouser@/Hydra")
	if err != nil {
		b.Fatal("Could not connect to hydra chat system", err)
	}

	allMembersBM(b, dblayer)
}

func BenchmarkMongoDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mongodb", "mongodb://127.0.0.1")
	if err != nil {
		b.Error("Could not connect to hydra chat system", err)
		return
	}

	allMembersBM(b, dblayer)
}

func allMembersBM(b *testing.B, dblayer DBLayer) {

	for i := 0; i < b.N; i++ {
		_, err := dblayer.AllMembers()
		if err != nil {
			b.Error("Query failed ", err)
			return
		}
	}
}
