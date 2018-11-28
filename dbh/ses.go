package dbh

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"strings"

	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var s3Region = envhelp.GetEnv("AWS_DEFAULT_REGION", "us-east-1")

// SeSHandler --
type SeSHandler struct {
	*ses.SES
}

// New --
func (s *SeSHandler) New() error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if err != nil {
		log.Error(err)
		return err
	}
	// Quick verification step
	if _, err = sess.Config.Credentials.Get(); err != nil {
		log.Error(err)
		return err
	}
	s.SES = ses.New(sess)

	return nil

}

// GenerateAndSendEmail --
func (s *SeSHandler) GenerateAndSendEmail(source, destination, subject, message string) error {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// email main header:
	h := make(textproto.MIMEHeader)
	h.Set("From", source)
	h.Set("To", destination)
	h.Set("Return-Path", source)
	h.Set("Subject", subject)
	h.Set("Content-Language", "en-US")
	h.Set("Content-Type", "multipart/mixed; boundary=\""+writer.Boundary()+"\"")
	h.Set("MIME-Version", "1.0")
	_, err := writer.CreatePart(h)
	if err != nil {
		return err
	}

	// body:
	h = make(textproto.MIMEHeader)
	h.Set("Content-Transfer-Encoding", "7bit")
	h.Set("Content-Type", "text/html; charset=utf-8")
	part, err := writer.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write([]byte(message))
	if err != nil {
		return err
	}

	// Strip boundary line before header (doesn't work with it present)
	buff := buf.String()
	if strings.Count(buff, "\n") < 2 {
		return fmt.Errorf("invalid e-mail content")
	}
	buff = strings.SplitN(buff, "\n", 2)[1]

	raw := ses.RawMessage{
		Data: []byte(buff),
	}
	input := &ses.SendRawEmailInput{
		Destinations: []*string{aws.String(destination)},
		Source:       aws.String(source),
		RawMessage:   &raw,
	}

	if ok, err := s.SendRawEmail(input); err != nil {
		log.Error(err, ok)
	}
	return nil
}
