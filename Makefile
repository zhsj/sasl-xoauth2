CLIENTID :=
CLIENTSECRET :=

.PHONY: clean

all: build

oauth/conf.go: oauth/conf.go.in
	sed 's|##CLIENTID##|$(CLIENTID)|g; s|##CLIENTSECRET##|$(CLIENTSECRET)|g' $< > $@

build: main.go oauth/oauth.go oauth/conf.go
	go build -buildmode=c-shared -ldflags="-s -w" -o libxoauth2.so
	go build -ldflags="-s -w"

cgo: main.go oauth/oauth.go oauth/conf.go
	go-7 build -buildmode=c-shared -o libxoauth2.so
	go-7 build

install:
	install -m755 libxoauth2.so /usr/lib/sasl2/

clean:
	rm -rf sasl-xoauth2 libxoauth2.so *.h oauth/conf.go
