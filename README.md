# zabbix-sender-http
Tool for receive events over http and pass to zabbix. zabbix-sender over http.
zabbix-sender-http does not depend on the zabbix_sender tool and fully implements its interface.

[![build](https://github.com/0xdeface/zabbix-sender-http/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/0xdeface/zabbix-sender-http/actions/workflows/build.yml)
![release](https://img.shields.io/github/v/release/0xdeface/zabbix-sender-http.svg)

**Zabbix-sender-http** - Инструмент для приема событий по http и отправки этих событий в zabbix. 

### (builds) Готовые сборки
[windows_x64_build](dist/zabbix-http.exe)

[linux_build](dist/zabbix-http)


Проверен и совместим с zabbix v5.4    

### (usage) Использование:
Просто запустите **zabbix-http**, это запустит веб сервер готовый принимать данные.
Для отправки данных сформируйте GET запрос. Сервер ожидает
следующие параметры запроса ["server", "key", "value"]

Just run **zabbix-http**, this command will run web server ready to receive data. 
To send data you should make http Get request with these query parameters: ["server", "key", "value"] 
```bash
    curl localhost:8080?server=HOST_NAME&key=ITEM&value=MYVAL
```
##### Пример запуска docker контейнера
```bash
    docker run --name zabbix-sender-http -d \
    -p 3001:8080 \
    -e ZABBIX_SERVER="zabbix-server-pgsql" \
    --restart always \
    --network postgres_network \
    ghcr.io/0xdeface/zabbix-sender:latest
```

### (launch parameters) Параметры запуска

Таблица ниже отображает доступные параметры запуска и их приоритеты. 
The table bellow shows possible launch parameters and their priority. 

| Highest priority       | Middle priority     | Lowest priority   | Description               | 
|------------------------|---------------------|-------------------|---------------------------|
| **Command parameters** | **Env variables**   | **Predefined values** |                       |   
| zabbix-server          | ZABBIX_SERVER       | 127.0.0.1         | set zabbix server address |   
| zabbix-port            | ZABBIX_PORT         | 10051             | set zabbix server port    |   
| http-port              | HTTP_PORT           | 8080              | http server port          |   

### Прием данных в zabbix
Просто создайте элемент данных с типом zabbix trapper. Трапперы используются для приема данных которые пушатся из вне,
а не собираются самим zabbix'ом.    

