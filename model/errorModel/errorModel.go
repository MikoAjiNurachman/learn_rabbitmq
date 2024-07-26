package errorModel

import "log"

func FailOnErr(err error, message string) {
	log.Fatalf(`%s: %s`, message, err)
}
