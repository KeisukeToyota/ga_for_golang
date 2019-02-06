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

func GetPopulation(geneLength int, individualLength int) Population {
	rand.Seed(time.Now().UnixNano())
	var population Population

	for i := 0; i < individualLength; i++ {
		var genes Genes
		for j := 0; j < geneLength; j++ {
			genes = append(genes, Gene(rand.Intn(2)))
		}
		population = append(population, Individual{
			Fitness: 0,
			Genes:   genes,
		})
	}
	return population
}

func (ind Individual) CalcFitness() int {
	r := 0
	for _, gene := range ind.Genes {
		r += int(gene)
	}
	return r
}

func (p Population) Evaluate() {
	for i := 0; i < p.Len(); i++ {
		p[i].Fitness = p[i].CalcFitness()
	}
	sort.Sort(sort.Reverse(p))
}

func main() {
	population := GetPopulation(10, 100)
	population.Evaluate()
	fmt.Println(population)
}
