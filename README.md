# CLI URL-Loader

### Task
Необходимо реализовать CLI-утилиту, которая реализует асинхронную обработку входящих URL из файла, переданного в качестве аргумента данной утилите.
Формат входного файла: на каждой строке – один URL. URL может быть очень много! Но могут быть и невалидные URL.

Пример входного файла:
https://myoffice.ru
https://yandex.ru

### Requirements
По каждому URL получить контент и вывести в консоль его размер и время обработки. Предусмотреть обработку ошибок

## The following concepts are applied in util:
- <a href="https://github.com/spf13/cobra">CLI Cobra</a> Cobra is a library for creating powerful modern CLI applications.
- Async
- Docker compose
- CI (GitHub Action)

#### to check this out
```shell
   make run
```

#### to use this out
```shell
   make build
   bin/url-loader -h
```
