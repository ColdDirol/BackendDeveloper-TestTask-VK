# Профильное задание на Backend-разработчика в VK

## Запуск
```bash
docker-compose up --build app
```
* Работает скрипт по ожиданию инициализации базы данных при развертывании в докере. Просьба дождаться выполнения:
```
{"time":"2024-03-28T19:33:09.55887512Z","level":"INFO","msg":"You are in mode:","env":"dev"}
```
```
{"time":"2024-03-28T19:33:09.561525511Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"1","attempt":"200"}
{"time":"2024-03-28T19:33:10.562767308Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"2","attempt":"200"}
{"time":"2024-03-28T19:33:11.56417207Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"3","attempt":"200"}
{"time":"2024-03-28T19:33:12.565753759Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"4","attempt":"200"}
{"time":"2024-03-28T19:33:13.567280756Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"5","attempt":"200"}
{"time":"2024-03-28T19:33:14.569261673Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"6","attempt":"200"}
{"time":"2024-03-28T19:33:15.570587334Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"7","attempt":"200"}
{"time":"2024-03-28T19:33:16.571931352Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"8","attempt":"200"}
{"time":"2024-03-28T19:33:17.573103624Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"9","attempt":"200"}
{"time":"2024-03-28T19:33:18.574986354Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"10","attempt":"200"}
{"time":"2024-03-28T19:33:19.576431624Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"11","attempt":"200"}
{"time":"2024-03-28T19:33:20.57829764Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"12","attempt":"200"}
{"time":"2024-03-28T19:33:21.579429814Z","level":"INFO","msg":"PostgreSQL is unavailable - wait","attempt":"13","attempt":"200"}
```
```
{"time":"2024-03-28T19:33:22.652487305Z","level":"INFO","msg":"PostgreSQL is up - executing command"}
{"time":"2024-03-28T19:33:22.800036343Z","level":"INFO","msg":"Server start listening on port: ","port":"8080"}
```
## Навигация
* [Примеры запросов](https://github.com/ColdDirol/BackendDeveloper-TestTask-VK/blob/main/requests.http)
* [JWT реализация](https://github.com/ColdDirol/BackendDeveloper-TestTask-VK/blob/main/pkg/auth/jwt/jwt.go)
* [Конфигурационный файл](https://github.com/ColdDirol/BackendDeveloper-TestTask-VK/blob/main/config.json)

## Методы:
```http
POST GET http://localhost:8080/login
POST GET http://localhost:8080/registration

POST /advertisement
GET /advertisement/{id}
GET /advertisement/sortByDate/DESC/{pageNum}
GET /advertisement/sortByDate/ASC/{pageNum}
GET /advertisement/sortByCost/DESC/{pageNum}
GET /advertisement/sortByCOST/ASC/{pageNum}
PUT /advertisement
DELETE /advertisement/{id}
```
[использование](https://github.com/ColdDirol/BackendDeveloper-TestTask-VK/blob/main/requests.http)

`id` - id объекта таблицы advertisements

`pageNum` - страница, содержащая N эементов (где N - `page_elements` из [файла конфигурации](https://github.com/ColdDirol/BackendDeveloper-TestTask-VK/blob/main/config.json))

## Правила конфигурации
### Конфигурационный файл
```json
{
  "env": "dev",
  "page_elements": 10,
  "http_server": {
    "host": "0.0.0.0",
    "port": "8080"
  },
  "auth": {
    "secret_key": "secret_key",
    "salt": "salt"
  },
  "database": {
    "host": "postgres",
    "port": "5432",
    "username": "postgres",
    "password": "postgres",
    "db_name": "postgres",
    "ssl_mode": "disable"
  }
}
```

#### ENV:
В переменную env можно установить значения: `local`, `dev`, `prod`.
```json
"env": "dev"
```
От значения зависит уровень и способ логирования: `local` - уровень логирования - Debug, `dev` - уровень логирования - Debug, `prod` - уровень огирования - Info.

#### PAGE_ELEMENTS:
Переменная page_elements должно содержать целое неотрицательное число - количество элементов, которые вмещает страница (по умолчанию 10).
```json
"page_elements": 10
```

#### HTTP_SERVER:
```json
"http_server": {
  "host": "0.0.0.0",
  "port": "8080"
}
```
`host` - хост адреса, на котором поднимается сервер (`localhost` для локального запуска, `0.0.0.0` для запуска в докере)

`port` - порт адреса, на котором поднимается сервер

#### AUTH:
```json
"auth": {
  "secret_key": "secret_key",
  "salt": "salt",
  "secure_mode": "secure"
}
```
`secret_key` - secret key для составления jwt токена

#### DATABASE:
```json
"database": {
  "host": "db",
  "port": "5432",
  "username": "postgres",
  "password": "postgres",
  "db_name": "postgres",
  "ssl_mode": "disable"
}
```
`host` - хост, на котором расположена база данных (`localhost` для локального запуска, `db` для запуска в докере)

`post` - порт базы данных

`username`, `password`, `db_name` - данные к базе данных

`ssl_mode` - протокол шифрования: `disable` - не используется
