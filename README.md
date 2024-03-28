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
