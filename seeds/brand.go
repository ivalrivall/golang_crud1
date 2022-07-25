package seeds

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

func (s Seed) BrandSeed() {

	for i := 0; i < 5; i++ {
		a := FakeStruct{}
		f := faker.FakeData(&a)
		if f != nil {
			fmt.Println(f)
		}
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO brands(name) VALUES ($1)`)
		// execute query
		_, err := stmt.Exec(&a.Word)
		if err != nil {
			panic(err)
		}
	}
}
