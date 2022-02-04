package bankutil

import "os"

var PanPrefix = os.Getenv("PAN_PREFIX")

func CensorPaymentString(input string) string {
	if len(input) < 6 {
		return "****" //TODO delete
	}
	return input[0:6] + "****" + input[len(input)-4:]
}
