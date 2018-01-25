# SASL XOAUTH2

## Build

Go https://console.developers.google.com/ and obtain an OAuth2 ID/secret.

Build dependency: `libsasl2-dev`.

```
make build CLIENTID= CLIENTSECRET=
```

## Install

```
sudo make install
```

## Use

Run `./sasl-xoauth2` to obtain a token, then set it as a password in `muttrc`, like:

```
set imap_user   = "example@gmail.com"
set imap_pass   = "token"

# set imap_authenticators="XOAUTH2"
set smtp_authenticators="XOAUTH2"
```
