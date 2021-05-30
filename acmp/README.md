# acmp
Необходимо реализовать функцию *Difficulty*, которая на вход принимает url задачи с сайта https://acmp.ru.
Пример адреса: https://acmp.ru/index.asp?main=task&id_task=1
Функция должна вернуть сложность задачи в виде *float64*.
Если ссылка на задачу некорректна, произошла любая ошибка, либо сложность задачи на странице не найдена, функция должна вернуть -1.

Для получения страницы на английском языке необходимо передать Cookie
```cookie
English:1
```

![Screenshot](img.png?raw=true "acmp.ru")

Подсказки:
* https://golang.org/pkg/net/http/#NewRequest
* [Пример запросов](https://blog.logrocket.com/making-http-requests-in-go/)
* Для парсинга сложности можно воспользоваться регулярным выражением:
```regesp
Difficulty: (\d+)%\)
```
* Также можно воспользоваться пакетом [goquery](https://github.com/PuerkitoBio/goquery)