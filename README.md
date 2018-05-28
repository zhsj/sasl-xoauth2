# SASL XOAUTH2

## Build

Go https://console.developers.google.com/ (Credentials -> Create credentials -> OAuth client ID).
Then obtain a pair of OAuth2 ID & secret.

Build dependency: `libsasl2-dev`.

```
make build CLIENTID=YOUR_CLIENT_ID CLIENTSECRET=YOUR_CLIENT_SECRET
```

## Install

```
sudo make install
```

## Use

Use [Mutt](http://www.mutt.org/) to login GMail IMAP/SMTP with
[OAuth2](https://developers.google.com/gmail/imap/xoauth2-protocol).

First run `./sasl-xoauth2` to obtain a token, then set it as a password in `muttrc`, like:

```
set folder      = "imaps://imap.gmail.com/"

set imap_user   = "example@gmail.com"
set imap_pass   = "token"

set smtp_url = "smtps://$imap_user:$imap_pass@smtp.gmail.com"

# set imap_authenticators="XOAUTH2"
set smtp_authenticators="XOAUTH2"
```
