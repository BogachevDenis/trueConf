# trueConf
## Запуск проекта:
Для запуска приложения выполните следующие команды:
<br>
$ git clone git@github.com:BogachevDenis/trueConf.git
<br>
$ cd trueConf
<br>
$ go run main.go
<br>
#### Приложение запустится на http://localhost:8080/
## Запуск тестов:
$ cd trueConf
<br>
$ go test
#### Покрытие -74.8%

### Примеры запросов для работы с хранилищем
<li>POST Запрос на добавление новых данных
  <br>
  $ curl -X POST -d '{"name":"ivan"}' http://localhost:8080/user
  <br>
<li>GET Запрос на получение всех пользователей
  <br>
  $ curl -X GET  http://localhost:8080/user
  <br>
  <li>GET Запрос на получение пользователя по id
  <br>
  $ curl -X GET  http://localhost:8080/user/1, где 1 - id пользователя
  <br>
