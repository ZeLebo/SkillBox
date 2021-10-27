package main

import (
    "fmt"
)

func exam(){
    var result, tmp int = 0, -1
    fmt.Println("Баллы ЕГЭ.")
    fmt.Println("Введите результат первого экзамена:")
    fmt.Scanf("%d", &tmp)
    if tmp > -1 && tmp < 101{
        result += tmp
    }
    fmt.Println("Введите результат второго экзамена:")
    fmt.Scanf("%d", &tmp)
    if tmp > -1 && tmp < 101{
        result += tmp
    }
    fmt.Println("Введите результат третьего экзамена:")
    fmt.Scanf("%d", &tmp)
    if tmp > -1 && tmp < 101{
        result += tmp
    }
    fmt.Println("Сумма проходных баллов: 275")
    fmt.Println("Колличество набранных баллов:", result)
    if result == -1{
        fmt.Println("Введенные даныые неверны, или у вас неправильные экзамены")
    } else if result < 275 {
        fmt.Println("Вы не поступили.")
    } else {
        fmt.Println("Вы поступили.")
    }
}

func nums() {
    fmt.Println("Три числа.")
    var a, b, c int
    fmt.Println("Введите 3 числа")
    fmt.Scanf("%d %d %d", &a, &b, &c)
    if a > 5 || b > 5 || c > 5 {
        fmt.Println("Есть больше 5")
    } else {
        fmt.Println("Числа маловаты")
    }
}

func money() {
    print("Банкомат.\n")
    print("Введите сумму снятия со счета:\n")
    var cash int
    fmt.Scanf("%d", &cash)
    if ( cash % 100 == 0) && (cash < 100001 ) {
        fmt.Println("Операция выполенена успешно")
        fmt.Println("Вы сняли со счета", cash, "рублей.")
    } else {
        fmt.Println("Операция не может быть выполнена")
    }
}

func numsVer2() {
    fmt.Println("Введите 3 числа")
    var a, b, c int
    cnt := 0
    fmt.Scanf("%d %d %d", &a, &b, &c)
    if (a > 4) {
        cnt++
    }
    if (b > 4) {
        cnt++
    }
    if (c > 4) {
        cnt++
    }
    if (cnt != 0) {
        fmt.Println("Среди введенных чисел", cnt, "больше или равны 5.")
    } else {
        fmt.Println("Среди введенных числе нет больших 4")
    }
}

func restaraunt() {
    var day, guestsNumber int
    var summary, discount, addition float32
    fmt.Println("Введите день недели:")
    fmt.Scanf("%d", &day)
    fmt.Println("Введите число гостей:")
    fmt.Scanf("%d", &guestsNumber)
    fmt.Println("Введите сумму чека:")
    fmt.Scanf("%f", &summary)

    if (day == 1) {
        discount = summary * 0.1
        fmt.Println("Скидка по понедельникам:", discount)
    } else if (day == 5) && (summary > 10000) {
        discount = summary * 0.05
        fmt.Println("Скидка на заказ:", discount)
    }
    if (guestsNumber > 5) {
        addition = summary * 0.1
        fmt.Println("Надбавка на обслуживание:", addition)
    }
    fmt.Println("Сумма к оплате:", summary - discount + addition)
}

func student() {
    var n, k, number int = 1, 1, 1
    fmt.Println("Введите число студентов, Колличество групп и номер студента")
    fmt.Scanf("%d %d %d", &n, &k, &number)
    n = number % k
    if ( n == 0) {
        n = k
    }
    fmt.Println(n)
}

func tax(){
    var income, taxPrice float32

    fmt.Println("Сколько ты зарабатываешь?")
    fmt.Scanf("%f", &income)
    if ( income < 10000) {
        taxPrice = income * 0.13
    } else if ( income < 50000) {
        taxPrice = 10000 * 0.13
        taxPrice += (income - 10000) * 0.2
    } else {
        taxPrice = 10000 * 0.13
        taxPrice += 40000 * 0.2
        taxPrice += (income - 50000) * 0.3

    }
    fmt.Println("Ты должен мне:", taxPrice)
}
func main() {
    exam()
    fmt.Println()
    nums()
    fmt.Println()
    money()
    fmt.Println()
    numsVer2()
    fmt.Println()
    restaraunt()
    fmt.Println()
    student()
    fmt.Println()
    tax()
}

