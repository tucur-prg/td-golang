package main

import (
    "fmt"
    "net/smtp"
    "os"
    "strings"
)

var (
	hostname = "host.docker.internal"
	port = 1025
)

// docker run --name maildev --rm -d -p 8025:1080 -p 1025:1025 maildev/maildev
func main() {
    from := "info@example.net"
    recipients := []string{"foo@example.com", "bar@example.com"}

    subject := "hello"
    body := "Hello World!\nHello Gopher!"

//    auth := smtp.PlainAuth("", "", "", hostname)
    msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body), "\n", "\r\n"))

    if err := smtp.SendMail(
        fmt.Sprintf("%s:%d", hostname, port),
        nil,
        from,
        recipients,
        msg,
    ); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
