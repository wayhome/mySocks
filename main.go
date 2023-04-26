package main

import (
	"flag"
	"fmt"
	"github.com/armon/go-socks5"
)

type myCredentialStore struct {
	user     string
	password string
}

func (cs *myCredentialStore) Valid(user, password string) bool {
	return user == cs.user && password == cs.password
}

func main() {
	username := flag.String("u", "", "username for SOCKS5 proxy")
	password := flag.String("P", "", "password for SOCKS5 proxy")
	host := flag.String("h", "0.0.0.0", "host for SOCKS5 proxy")
	port := flag.String("p", "1080", "port for SOCKS5 proxy")
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		return
	}

	auth := socks5.UserPassAuthenticator{
		Credentials: &myCredentialStore{user: *username, password: *password},
	}

	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	listenAddr := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Printf("Starting SOCKS5 proxy server on %s\n", listenAddr)

	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		panic(err)
	}
}
