CLIENTID :=
CLIENTSECRET :=

LIBDIR := /usr/lib/$(shell gcc -print-multiarch)/sasl2


.PHONY: clean

all: build

oauth/conf.go: oauth/conf.go.in
	sed 's|##CLIENTID##|$(CLIENTID)|g; s|##CLIENTSECRET##|$(CLIENTSECRET)|g' $< > $@

build: main.go oauth/oauth.go oauth/conf.go
	go build -buildmode=c-shared -ldflags="-s -w" -o libxoauth2.so
	go build -ldflags="-s -w"

install:
	mkdir -p $(LIBDIR)
	install -m644 libxoauth2.so $(LIBDIR)

uninstall:
	rm -rf $(LIBDIR)/libxoauth2.so

clean:
	rm -rf sasl-xoauth2 libxoauth2.so *.h oauth/conf.go
