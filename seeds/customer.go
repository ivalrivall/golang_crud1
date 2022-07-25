package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) CustomerSeed() {

	for i := 0; i < 5; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO customers(name) VALUES ($1)`)
		// execute query
		_, err := stmt.Exec(faker.Name())
		if err != nil {
			panic(err)
		}
	}
}
