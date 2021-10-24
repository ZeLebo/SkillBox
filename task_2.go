package main

import (
    "fmt"
)


func raceImprovement() {
    var lap int = 4
    var engine int = 254
    var wheels int = 93
    var steeringWheel int = 49
    var wind int = 21
    var rain int = 17
    speed := engine + wheels + steeringWheel - wind - rain

    fmt.Println("===================")
    fmt.Print("Супергонки. Круг ", lap, "\n")
    fmt.Println("===================")
    fmt.Print("Шумахер (", speed, ")\n")
    fmt.Println("===================")
    fmt.Println("Водитель: Шумахер")
    fmt.Print("Скорость: ", speed, "\n")
    fmt.Println("-------------------")
    fmt.Println("Оснащение")
    fmt.Print("Двигатель: +", engine, "\n")
    fmt.Print("Колёса: +", wheels, "\n")
    fmt.Print("Руль: +", steeringWheel, "\n")
    fmt.Println("-------------------")
    fmt.Println("Действия плохой погоды")
    fmt.Print("Ветер: −", wind, "\n")
    fmt.Print("Дождь: −", rain, "\n")
}

func calculator() {
    var price, delivery, discount int
    fmt.Print("Калькулятор стоимости товара со скидкой.\n")
    fmt.Print("Стоимтость товара: ")
    fmt.Scanf("%d", &price)
    fmt.Print("Стоимость доставки: ")
    fmt.Scanf("%d", &delivery)
    fmt.Print("Размер скидки: ")
    fmt.Scanf("%d", &discount)
    fmt.Println("---------")
    price += delivery
    price -= discount
    fmt.Println("Итого:", price)
}

func time() {
    fmt.Println("Расчет количества клиентов, обслуживаемых за смену")
    var duration, clientTime, cashierTime int
    fmt.Print("Введите длительность смены в минутах: ")
    fmt.Scanf("%d", &duration)
    fmt.Print("Сколько минут клиент делает заказ: ")
    fmt.Scanf("%d", &clientTime)
    fmt.Print("Сколько минут кассир собирает заказ: ")
    fmt.Scanf("%d", &cashierTime)
    fmt.Println("-----Считаем-----")
    var result int = duration / (clientTime + cashierTime)
    fmt.Print("За смену длиной ", duration, " Кассир успеет обслужить ", result, " клиентов.\n")
}

func receipt(){
    var summary, entranceAmount, flatAmount int
    fmt.Println("Расчет стоимтости ремонта для каждой квартиры")
    fmt.Print("Сумма, указанная в квитанции в рублях: ")
    fmt.Scanf("%d", &summary)
    fmt.Print("Подъездов в доме: ")
    fmt.Scanf("%d", &entranceAmount)
    fmt.Print("Квартир в каждом подъезде: ")
    fmt.Scanf("%d", &flatAmount)
    result := float32 (summary / (entranceAmount * flatAmount))
    fmt.Println("----Результат----")
    fmt.Println("Каждая квартира должна платить по", result, "рублей.")
}

func main() {
    raceImprovement()
    fmt.Println()
    calculator()
    fmt.Println()
    time()
    fmt.Println()
    receipt()
}
