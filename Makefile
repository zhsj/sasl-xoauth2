CLIENTID :=
CLIENTSECRET :=

.PHONY: clean

all: build install

build: main.go oauth.go
	go build -buildmode=c-shared -ldflags="-s -w -X main.ClientID=$(CLIENTID) -X main.ClientSecret=$(CLIENTSECRET)" -o libxoauth2.so
	go build -ldflags="-s -w -X main.ClientID=$(CLIENTID) -X main.ClientSecret=$(CLIENTSECRET)"

install:
	install -m755 libxoauth2.so /usr/lib/sasl2/

clean:
	rm -rf sasl-xoauth2 libxoauth2.so libxoauth2.h
