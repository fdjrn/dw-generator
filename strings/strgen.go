package strings

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const (
	TransTopUp   = "Top-Up"
	TransPayment = "Payment"
)

func GetUnixTime() string {
	tUnixMicro := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Microsecond)
	return strconv.FormatInt(tUnixMicro, 10)
}

func GetUnixTimeMicro() string {
	tUnixMicro := int64(time.Nanosecond) * time.Now().UnixNano() / 1000
	return strconv.FormatInt(tUnixMicro, 10)
}

func GetUnixTimeNano() string {
	tUnixMicro := int64(time.Nanosecond) * time.Now().UnixNano()
	return strconv.FormatInt(tUnixMicro, 10)
}

func GenerateReceiptNumber(transType string, id string) string {
	tUnix := GetUnixTimeNano()
	var r string

	switch transType {
	case TransTopUp:
		r = fmt.Sprintf("1000%s%s", tUnix, id)
	case TransPayment:
		r = fmt.Sprintf("2000%s%s", tUnix, id)
	}

	return r
}

func GenerateTransNumber() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return fmt.Sprintf("%s%s", time.Now().Format("20060102"), string(b))
}
