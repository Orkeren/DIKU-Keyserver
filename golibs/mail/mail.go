package mail

import (
        "bytes"
        "log"
        "net/smtp"
)

func Send(rcpt, body string) {
		From := `From: DIKU Keys <noreply@dikukeys.dk>
`
		To := "To: " + rcpt + `
`
		Subject := `Subject: Velkommen til DIKU Keys
`
		epost := From + To + Subject + body
        // Connect to the remote SMTP server.
        c, err := smtp.Dial("dikukeys.dk:25")
        if err != nil {
                log.Fatal(err)
        }
        // Set the sender and recipient.
        c.Hello("dikukeys.dk")
        c.Mail("noreply@dikukeys.dk")
        c.Rcpt(rcpt)
        // Send the email body.
        wc, err := c.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        buf := bytes.NewBufferString(epost)
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
}