FROM scratch
ADD zabbix-http /main
EXPOSE 8080
ENTRYPOINT ["/main"]
