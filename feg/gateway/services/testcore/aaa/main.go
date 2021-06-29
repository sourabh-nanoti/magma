package main

import (
	"log"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {
	handler := func(w radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)

		var code radius.Code
		if r.Packet.Code == radius.CodeAccessRequest {
			log.Printf("Received Access Request for Username %s ", username)
			code = radius.CodeAccessAccept
		} else if r.Packet.Code == radius.CodeDisconnectRequest {
			log.Printf("Received Disconnect Request for Username %s ", username)
			code = radius.CodeDisconnectACK
		}
		log.Printf("Writing %v to %v", code, r.RemoteAddr)
		w.Write(r.Response(code))
	}

	server := radius.PacketServer{
		Addr:         "127.0.0.1:50008",
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(`123456`)),
	}

	log.Printf("Starting server on :50008")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
