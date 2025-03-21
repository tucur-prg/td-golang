package main

import (
	"fmt"
	"os"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

/*
* openssl genrsa -out server.key 4096
* openssl req -new -key server.key -out server.csr -subj "/C=JP/ST=Tokyo/L=Minato City/O=Local/CN=host.docker.internal"
* openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt -subj "/C=JP/ST=Tokyo/L=Minato City/O=local/CN=host.docker.internal" -extfile <(echo "subjectAltName=DNS:host.docker.internal")
*/
// docker run --name mailpit --rm -d -p 8026:8025 -p 1026:1025 -v `pwd`/cert:/cert -e MP_SMTP_TLS_CERT=/cert/server.crt -e MP_SMTP_TLS_KEY=/cert/server.key -e MP_SMTP_AUTH_FILE=/cert/.htpasswd axllent/mailpit
func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "info@example.com", "テスト")
	m.SetAddressHeader("Reply-to", "", "テスト")
	m.SetAddressHeader("To", "tester@example.com", "")
	m.SetHeader("Subject", "日本語メールだよ")
	m.SetBody("text/html", `<html><body>
メールの内容だよ<br>
ほげ<br>
ほげ<br>
<br>
<a href="http://localhost:4000/">Localhost</a><br>
<br>
ふが<br>
ふが<br>
</body></html>`)

	d := gomail.NewDialer("host.docker.internal", 1026, "mailpit", "mailpit")
	fmt.Printf("%T\n", d)
	fmt.Println(d.Auth, d.SSL, d.TLSConfig)

	w, err := os.OpenFile("tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName: "host.docker.internal",
		KeyLogWriter: w,
	}

	sender, err := d.Dial()
	if err != nil {
		panic(err.Error())
	}
	defer sender.Close()

	fmt.Printf("%T: %v\n", sender, sender)

	if err := gomail.Send(sender, m); err != nil {
		panic(err.Error())
	}

	fmt.Println("sendmail")
}
