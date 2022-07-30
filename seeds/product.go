package seeds

import (
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func (s Seed) ProductSeed() {

	for i := 0; i < 20; i++ {
		//prepare the statement
		a := FakeStruct{}
		f := faker.FakeData(&a)
		if f != nil {
			fmt.Println(f)
		}
		stmt, _ := s.db.Prepare(`INSERT INTO products(name, brand_id, price) VALUES ($1, $2, $3)`)
		// execute query
		_, err := stmt.Exec(&a.Word, rand.Intn(5-1)+1, rand.Intn(10000-1000)+1000)
		if err != nil {
			panic(err)
		}
	}
}
