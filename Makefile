CLIENTID :=
CLIENTSECRET :=

PKG := github.com/zhsj/sasl-xoauth2
LDFLAGS := -s -w -X $(PKG)/oauth.ClientID=$(CLIENTID) -X $(PKG)/oauth.ClientSecret=$(CLIENTSECRET)

.PHONY: clean

all: build

build: main.go oauth/oauth.go
	go build -buildmode=c-shared -ldflags="$(LDFLAGS)" -o libxoauth2.so
	go build -ldflags="$(LDFLAGS)"

install:
	install -m755 libxoauth2.so /usr/lib/sasl2/

clean:
	rm -rf sasl-xoauth2 libxoauth2.so libxoauth2.h
