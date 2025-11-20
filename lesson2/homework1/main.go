package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p *Product) IsInStock() bool {
	return (p.Stock > 0)
}
func (p *Product) Info() string {
	return fmt.Sprintf("商品: %s,单价: %f,库存: %d", p.Name, p.Price, p.Stock)
}
func (p *Product) Restock(amount int) {
	p.Stock += amount
}
func (p *Product) Sell(amount int) (success bool, message string) {
	if p.Stock >= amount {
		p.Stock -= amount
		return true, "售卖成功"

	} else {
		return false, "库存不足"
	}
}

func main() {
	a := Product{
		Name:  "Go语言编程",
		Price: 89.5,
		Stock: 10,
	}
	s, m := a.Sell(5)
	fmt.Println(s, m, "剩余库存：", a.Stock)
	a.Restock(20)
	fmt.Println("当前库存：", a.Stock)
	s1, m1 := a.Sell(30)
	fmt.Println(s1, ",", m1)
	str := a.Info()
	fmt.Println(str)
	fmt.Printf("库存总值：%.2f", float64(a.Stock)*a.Price)

}
