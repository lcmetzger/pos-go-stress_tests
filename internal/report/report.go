package report

import (
	"fmt"
	"time"
)

const (
	FORMAT = "02/01/2006 15:04:05.000"
)

type Report struct {
	Url          string
	Start        time.Time
	Finish       time.Time
	TotalRequest int64
	MapRequest   map[int]int64
	Concurrency  int
	Requests     int
}

func NewReport() *Report {
	return &Report{
		Start:        time.Now(),
		MapRequest:   make(map[int]int64),
		TotalRequest: 0,
	}
}

func (r *Report) Results() {
	r.Finish = time.Now()

	fmt.Println("\nRelatório de testes de carga")
	fmt.Println("----------------------------")

	fmt.Printf("URL alvo               %s\n", r.Url)
	fmt.Printf("Requisições            %d\n", r.Requests)
	fmt.Printf("Concorrentes           %d\n", r.Concurrency)
	fmt.Printf("Inicio                 %v \n", r.Start.Format(FORMAT))
	fmt.Printf("Final                  %v \n", r.Finish.Format(FORMAT))
	fmt.Printf("Tempo transcorrido     %v\n\n", r.Finish.Sub(r.Start))

	fmt.Println("Mapa de Requisições efetuadas")
	fmt.Println("-----------------------------")

	for k, v := range r.MapRequest {
		fmt.Printf("%d -> %d requisições\n", k, v)
	}

	fmt.Println("-----------------------------")
	fmt.Printf("Total de requisições %d\n", r.TotalRequest)
	fmt.Print("-----------------------------\n\n")

}
