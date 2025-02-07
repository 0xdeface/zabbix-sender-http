package http

import (
	"context"
	"fmt"
	zabbix "github.com/0xdeface/zabbix/sender"
	"log"
	"net/http"
)

var required = []string{"server", "key", "value"}

func RunServer(ctx context.Context, port string, msgCh chan zabbix.Message, errCh chan error, debug bool) {

	server := &http.Server{Addr: ":" + port, Handler: handler(msgCh, errCh, debug)}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		errCh <- err
	}
}

func handler(msgCh chan zabbix.Message, errCh chan error, debug bool) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		for _, val := range required {
			if _, ok := q[val]; !ok {
				_, err := fmt.Fprintf(writer, "%v: shouldn t be empty \n", val)
				if err != nil {
					errCh <- err
				}
				return
			}
		}
		if debug {
			fmt.Printf("#GOT %+v\n", q)
		}
		msgCh <- zabbix.CreateMessage(q["server"][0], q["key"][0], q["value"][0])
		_, err := fmt.Fprintf(writer, "\"{\"status\": \"ok\"}")
		if err != nil {
			errCh <- err
		}
	}
}
