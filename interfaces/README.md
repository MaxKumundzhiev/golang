# Интерфейсы
Через интерфейсы в го реализован полиморфизм. Полиморфизм - возможность работать с разными оббектами через 1 интерфейс. Напмер интерфейс животного, обьекты Собака и Кошка или интерфес оплаты, обьекты Visa и Mastercard карты.

```
type Payment interface {
    Pay(amount float64)
}

type CardPayment struct{}
func (CardPayment) Pay(amount float64) {
    fmt.Println("Paid by Card:", amount)
}

type PayPalPayment struct{}
func (PayPalPayment) Pay(amount float64) {
    fmt.Println("Paid by PayPal:", amount)
}

func ProcessPayment(p Payment) {
    p.Pay(100)
}

func main() {
    ProcessPayment(CardPayment{})
    ProcessPayment(PayPalPayment{})
}
```