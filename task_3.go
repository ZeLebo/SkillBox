package main

import (
    "fmt"
)

/*
Просто запрашиваем данные от пользователя и подставляем в текст
Да, тут имена читаются по одному слову, возможность сделать их многословными не реализована
Для этого же надо написать всё через scanner ("os" и "bufio")?
*/
func firstTask() {
    var planetName, starName, rangerName string = "Awesome planet", "Awesome star", "Awesome ranger"
    var money int = 1000000
    fmt.Println("Write the planet, star and ranger name and the amount of money")
    fmt.Scanf("%s %s %s %d", &planetName, &starName, &rangerName, &money)
    fmt.Print("Как вам, ", rangerName, ", известно, мы раса мирная, поэтому на наших военных кораблях используются\nнаёмники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на\nнаши корабли людей с планеты " , planetName,  " системы ", starName, ".\n" )
    fmt.Print("Но случилась неприятность: в связи с большими потерями в последнее время престиж\nпрофессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар\nпланеты ", planetName, ", впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с\nзакруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва! ")
    fmt.Print(rangerName, ", вы должны прибыть на планету ", planetName," системы ", starName, " и помочь выполнить план\nпризыва. Мы готовы выплатить вам премию в ",money ," кредитов за эту маленькую услугу.\n")
}


/*
Делаем массив для остановок
Считаем в цикле людей, что зашли и вышли
Для удобства счиатаем, что люди платят при входе
Для результат используем float64, ибо мало ли мы кортеж Земли повезем
*/
func secondTask() {
    fmt.Println("Программа симмуляции маршрутного такси")
    stations := [4]string {"Zeleboba's station","ZhoRa's station","New station","Old station"}
    peopleIn, peopleOut, peopleStay := 0, 0, 0
    result := 0.0

    for i := 0; i < 4; i++ {
        fmt.Print("Прибываем на остановку '", stations[i], "'.\n")
        fmt.Println("В салоне пассажиров:", peopleStay)
        fmt.Print("Сколько пассажиров вышло на остановке: ")
        fmt.Scanf("%d", &peopleOut)
        peopleStay -= peopleOut
        fmt.Print("Сколько пассажиров вошло на остановке: ")
        fmt.Scanf("%d", &peopleIn)
        peopleStay += peopleIn
        result += float64 ( 20 * peopleIn )
        fmt.Print("Отправляемся с остановки '", stations[i], "'\n")
        fmt.Println("В салоне пассажиров:", peopleStay)
        if ( i != 3) {
            fmt.Println("--------------Едем-------------")
        }
    }
    fmt.Print("Всего заработали: ", result, " рублей.\n")
    fmt.Print("Зарплата водителя: ", result / 4 , " рублей.\n")
    fmt.Print("Расходы на топливо: ", result / 5, " рублей.\n")
    fmt.Print("Налоги: ", result / 5, " рублей.\n")
    fmt.Print("Расходы на ремонт машины: ", result / 5, " рублей.\n")
    fmt.Print("Итого доход: ", result * 0.15," рублей.\n")
}

/*
Первый обмен максимально в лоб
Просто делаем новую переменную и меняем через нее
*/

func thirdTask() {
    fmt.Println("Variables changing")
    a := 42
    b := 153
    fmt.Println("a:", a)
    fmt.Println("b:", b)

    tmp := a
    a = b
    b = tmp

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}

/*
Даем пользователю возможность вводить все данные
Если он пропускает эту шикарную возможность, то используем уже известные значения

Дальше выводим длину бамбука на 3,5 день ( тип float32 (должно хватить  =) ))
После считаем в цикле день, на который бамбук созреет
Он может созреть к вечеру, поэтому есть проверка на это
В случае если он созреет к вечеру, то соотвествующее сообщение будет выведенно

Логика с dayTmp нужна для вывода корректной формы слова "день" в конце строки
Я не знаю, зачем я это сделал, но просто захотелось ( лучше бы на английском написал...)
Для удобства использовал switch-case и while ( только for, на самом деле )
*/
func forthTask(){
    var speed, eatingSpeed, firstLen, lenToCut float32 = 50.0, 20.0, 100.0, 300.0
    print("На сколько за день растет бамбук?\n")
    fmt.Scanf("%f", &speed)
    print("Сколько гусеницы съежают за ночь?\n")
    fmt.Scanf("%f", &eatingSpeed)
    print("Какая высота саженца?\n")
    fmt.Scanf("%f", &firstLen)
    print("Бамбук какой высоты можно продать?\n")
    fmt.Scanf("%f", &lenToCut)
    thirdDayLen := float32 (firstLen + (speed * 3.5 - eatingSpeed * 2))
    fmt.Println("Длина бамбука в середине третьего дня =", thirdDayLen, "см.")


    if ( firstLen < lenToCut && speed <= eatingSpeed) {
        fmt.Println("Бамбук не созреет, злые гусеницы съедят его раньше...")
        return
    }

    day := 1
    dayString := "дня."
    word := "через"
    for firstLen < lenToCut {
        firstLen += speed
        if firstLen >= lenToCut{
            word = "в конце"
            break
        }
        firstLen -= eatingSpeed
        day++
    }

    dayTmp := day
    for dayTmp > 100 {
        dayTmp %= 10
    }
    for dayTmp > 21 {
        dayTmp %= 20
    }
    if ( word == "через") {
        switch dayTmp {
        case 1:
            dayString = "день."
        case 2:
            dayString = "дня."
        case 3:
            dayString = "дня."
        case 4:
            dayString = "дня."
        default:
            dayString = "дней."
        }
    }
    fmt.Println("Бамбук созреет", word, day, dayString)
}


/*
Делаем смену в стиле python
Будет ругаться на разные типы данных
*/
func fifthTask(){
    fmt.Println("Variables changing")
    a := 42
    b := 153
    fmt.Println("a:", a)
    fmt.Println("b:", b)

    a, b = b, a

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}


/*
Делаем через swap на указателях в стиле С
*/
func swap(a, b *int) {
    tmp := *a
    *a = *b
    *b = tmp
}

func fifthTaskSwap(){
    fmt.Println("Variables changing")
    a := 42
    b := 153
    fmt.Println("a:", a)
    fmt.Println("b:", b)

    swap(&a, &b)

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}

/*
Просто свалка для фунцкий и их вызова
*/
func main () {
    firstTask()
    fmt.Println()
    secondTask()
    fmt.Println()
    thirdTask()
    fmt.Println()
    forthTask()
    fmt.Println()
    fifthTask()
    fmt.Println()
    fifthTaskSwap()
}
