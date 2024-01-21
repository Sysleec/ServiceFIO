package enrich

import (
	"context"
	"sync"
)

type result struct {
	age         int
	gender      string
	nationality string
	err         error
}

func Enrichment(ctx context.Context, name string) (int, string, string, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	res := make(chan result, 3)

	go func() {
		defer wg.Done()
		age, err := GetAge(ctx, name)
		res <- result{age: age, err: err}
	}()

	go func() {
		defer wg.Done()
		gender, err := GetGender(ctx, name)
		res <- result{gender: gender, err: err}
	}()

	go func() {
		defer wg.Done()
		nationality, err := GetNationality(ctx, name)
		res <- result{nationality: nationality, err: err}
	}()

	wg.Wait()
	close(res)

	var age int
	var gender, nationality string
	for r := range res {
		if r.err != nil {
			return 0, "", "", r.err
		}
		if r.age != 0 {
			age = r.age
		}
		if r.gender != "" {
			gender = r.gender
		}
		if r.nationality != "" {
			nationality = r.nationality
		}
	}

	return age, gender, nationality, nil
}
