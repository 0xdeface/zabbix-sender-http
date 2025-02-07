# zabbix-sender-http
Tool for receive events over http and pass to zabbix. zabbix-sender over http.
zabbix-sender-http does not depend on the zabbix_sender tool and fully implements its interface.

[![build](https://github.com/0xdeface/zabbix-sender-http/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/0xdeface/zabbix-sender-http/actions/workflows/build.yml)
![release](https://img.shields.io/github/v/release/0xdeface/zabbix-sender-http.svg)


### platform builds: 
[windows_x64_build](dist/zabbix-http.exe)

[linux_build](dist/zabbix-http)


tested with zabbix v5.4, 6.0, 7.2  

### usage:
Just run **zabbix-http**, this command will run web server ready to receive data. 
To send data you should make http Get request with these query parameters: ["server", "key", "value"] 
```bash
    curl localhost:8080?server=HOST_NAME&key=ITEM&value=MYVAL
```
##### example run on docker
```bash
    docker run --name zabbix-sender-http -d \
    -p 3001:8080 \
    -e ZABBIX_HOST="zabbix-server-pgsql" \
    --restart always \
    --network postgres_network \
    ghcr.io/0xdeface/zabbix-sender:latest
```

### launch parameters

The table bellow shows possible launch parameters and their priority. 

| Highest priority  | Middle priority   | Lowest priority       | Description               | 
|-------------------|-------------------|-----------------------|---------------------------|
| **Cmd arguments** | **Env variables** | **Predefined values** |                       |   
| --zabbix-host     | ZABBIX_HOST       | 127.0.0.1             | set zabbix server address |   
| --zabbix-port     | ZABBIX_PORT       | 10051                 | set zabbix server port    |   
| --http-port       | HTTP_PORT         | 8080                  | http server port          |   

### getting data in zabbix
You should create item with type **zabbix trapper**. Trappers it's special item type for receive pushed data.

