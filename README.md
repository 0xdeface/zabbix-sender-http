# zabbix-sender-http
Tool for receive events over http and pass to zabbix. zabbix-sender over http.

**Zabbix-sender-http** - Инструмент для приема событий по http и отправки этих событий в zabbix. 

В папке [dist](dist) находятся скомпилированные сборки для Windows(x64) и Linux.
zabbix-sender-http не зависит от zabbix_sender и использует реализацию протокола zabbix_sender.

Проверен и совместим с zabbix v5.4    

### Использование:
Просто запустите **zabbix-http**, это запустит веб сервер готовый принимать данные.

Для отправки данных сформируйте GET запрос. Сервер ожидает
следующие параметры запроса ["server", "key", "value"] 
```bash
    curl localhost:8080?server=HOST_NAME&key=ITEM&value=MYVAL
```

Вы можете указать параметры запуска zabbix-sender-http
```
--server 127.0.0.1 IP адрес сервера zabbix
--zabbix-port 10051 порт zabbix сервера
--http-port 8080 порт на котором будет запущен http сервер
``` 
Кроме того вы можете использовать переменные окружения, они **имеют приоритет** над параметрами запуска
```bash
    SERVER=127.0.0.1
    ZABBIX_PORT=10051
    HTTP_PORT=8080
```
     

