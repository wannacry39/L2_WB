package main

import (
	"errors"
	"fmt"
)

type Book struct {
	name  string
	price float64
}

func (bo Book) GetPrice() float64 {
	return bo.price
}
func (bo Book) GetName() string {
	return bo.name
}

type BookShop struct {
	name  string
	books []Book
}

func (b BookShop) Sell(user User, book_name string) error {
	for _, val := range b.books {
		if val.GetName() == book_name {
			fmt.Println("Книга есть в магазине!")
			if user.CheckCash() >= val.GetPrice() {
				fmt.Println("Покупка совершена!")
				return nil
			} else {
				fmt.Println("Недостаточно средств")
				return errors.New("not enough money")
			}

		} else {
			fmt.Println("Книга не найдена")
			return errors.New("book not found")
		}
	}
	return nil
}

type User struct {
	name string
	cash float64
}

func (u User) CheckCash() float64 {
	return u.cash
}

func main() {
	var (
		book1 = Book{"Onegin", 145}
		book2 = Book{"1984", 200}
		user  = User{"Oleg", 100}
		shop  = BookShop{"Labirint", []Book{book1, book2}}
	)

	shop.Sell(user, "Onegin") //facade pattern

}
