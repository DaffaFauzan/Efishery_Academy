// TUGAS 2 Hari ke 5 masih on progress
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Desc struct {
	Description string `json:"description"`
}
type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}
type Aggr struct {
	Desc Desc
	Jobs []Job
}

func fetch(url string, d interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, d); err != nil {
		return err
	}
	return err
}
func Aggregate() (Aggr, error) {
	defer caculateTime(time.Now())
	var wg sync.WaitGroup
	wg.Add(2)
	var desc Desc
	var descErr error
	go func() {
		descErr = fetch(
			"https://workspace-rho.vercel.app/api/description",
			&desc,
		)
		wg.Done()
	}()
	var jobs []Job
	var jobsErr error
	go func() {
		jobsErr = fetch(
			"https://workspace-rho.vercel.app/api/jobs",
			&jobs,
		)
		wg.Done()
	}()
	wg.Wait()
	if descErr != nil {
		return Aggr{}, descErr
	}
	if jobsErr != nil {
		return Aggr{}, jobsErr
	}
	aggr := Aggr{
		Desc: desc,
		Jobs: jobs,
	}
	return aggr, nil
}
func main() {
	aggr, err := Aggregate()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", aggr)
}
func caculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}
