package main

import (
    "fmt"
)

type Item struct {
    Name  string
    Price float64
}

type Cart struct {
    Items []Item
}

func (c *Cart) AddItem(item Item) {
    c.Items = append(c.Items, item)
}

func (c *Cart) Total() float64 {
    total := 0.0
    for _, item := range c.Items {
        total += item.Price
    }
    return total
}

func (c *Cart) PrintReceipt() {
    fmt.Println("Struk Belanja:")
    for _, item := range c.Items {
        fmt.Printf("%s: Rp%.2f\n", item.Name, item.Price)
    }
    fmt.Printf("Total: Rp%.2f\n", c.Total())
}

func main() {
    cart := Cart{}

    cart.AddItem(Item{Name: "Buku", Price: 50000})
    cart.AddItem(Item{Name: "Pensil", Price: 2000})
    cart.AddItem(Item{Name: "Penghapus", Price: 1000})

    cart.PrintReceipt()
}
