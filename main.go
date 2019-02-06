package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Gene int

type Genes []Gene

type Individual struct {
	Fitness int
	Genes   Genes
}

type Population []Individual

func (p Population) Len() int {
	return len(p)
}

func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Population) Less(i, j int) bool {
	return p[i].Fitness < p[j].Fitness
}

func (p Population) CalcFitness() {
	for i := 0; i < p.Len(); i++ {
		r := 0
		for _, gene := range p[i].Genes {
			r += int(gene)
		}
		p[i].Fitness = r
	}
}

func (p Population) Evaluate() {
	p.CalcFitness()
	sort.Sort(sort.Reverse(p))
}

func GetPopulation(geneLen int, indLen int) Population {
	rand.Seed(time.Now().UnixNano())
	var population Population

	for i := 0; i < indLen; i++ {
		var genes Genes
		for j := 0; j < geneLen; j++ {
			genes = append(genes, Gene(rand.Intn(2)))
		}
		population = append(population, Individual{
			Fitness: 0,
			Genes:   genes,
		})
	}
	return population
}

func main() {
	population := GetPopulation(10, 100)
	population.Evaluate()
	fmt.Println(population)
}
