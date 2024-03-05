package main

import "fmt"

const (
	SmallProd  = "SmallProd"
	MediumProd = "MediumProd"
	LargeProd  = "LargeProd"
)

type SmallProduct struct {
	Cost float64
}

func (sp SmallProduct) GetCost() float64 {
	return sp.Cost
}

type MediumProduct struct {
	Cost float64
}

func (mp MediumProduct) GetCost() float64 {
	return mp.Cost * 1.03
}

type LargeProduct struct {
	Cost float64
}

func (lp LargeProduct) GetCost() float64 {
	return lp.Cost*1.06 + 2500
}

type Product interface {
	GetCost() float64
}

func ProductFactory(productType string, cost float64) Product {
	switch productType {
	case SmallProd:
		return SmallProduct{cost}
	case MediumProd:
		return MediumProduct{cost}
	case LargeProd:
		return LargeProduct{cost}
	default:
		return nil
	}
}

func main() {
	sp1 := ProductFactory(SmallProd, 10.0)
	mp1 := ProductFactory(MediumProd, 100.0)
	lp1 := ProductFactory(LargeProd, 1000.0)

	fmt.Println(sp1.GetCost())
	fmt.Println(mp1.GetCost())
	fmt.Println(lp1.GetCost())
}
