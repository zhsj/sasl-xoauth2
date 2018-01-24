CLIENTID :=
CLIENTSECRET :=

.PHONY: clean

all: build install

build: main.go oauth.go
	go build -buildmode=c-shared -ldflags="-s -w -X main.ClientID=$(CLIENTID) -X main.ClientSecret=$(CLIENTSECRET)"

install:
	install -m755 sasl-xoauth2 /usr/lib/sasl2/libxoauth2.so

clean:
	rm -rf sasl-xoauth2 sasl-xoauth2.h
