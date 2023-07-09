package main

import "fmt"

//////////// Interface ///////////////////////////

// OTP = One Time Password
type iOTP interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

//////////// Template ////////////////////////////

// Important: The combination of the iOTP interface and the otp struct provides the capabilities of an ABSTRACT CLASS in Go!!!

// generic OTP template = abstract class
type otp struct {
	iOTP iOTP
}

// template method
func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOTP.genRandomOTP(otpLength)
	o.iOTP.saveOTPCache(otp)
	message := o.iOTP.getMessage(otp)
	err := o.iOTP.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOTP.publishMetric()
	return nil
}

/////////////// iOTP Implementations /////////////

/////////////////// SMS //////////////////////////

type sms struct {
	otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "0815"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

///////////////// Email //////////////////////////

type email struct {
	otp
}

func (s *email) genRandomOTP(len int) string {
	randomOTP := "4711"
	fmt.Printf("Email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}

func (s *email) publishMetric() {
	fmt.Printf("Email: publishing metrics\n")
}

////////////// Main //////////////////////////////

func main() {
  // SMS
	smsOTP := &sms{}
	o := otp{iOTP: smsOTP}
	o.genAndSendOTP(4)

  // Problem!!!
  // smsOTP.genAndSendOTP(4) // compiler allows this but will crash
	fmt.Println()

  // email
	emailOTP := &email{}
	o = otp{iOTP: emailOTP}
	o.genAndSendOTP(4)
}
