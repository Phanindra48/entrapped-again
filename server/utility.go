package entrapped

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type command string

const (
	typeData      string = "data"
	typeError     string = "error"
	cmdRegistered string = "registered"
	cmdReady      string = "ready"
	cmdOpen       string = "open"
	cmdResult     string = "result"
	cmdEnemy      string = "enemy"
)

var dataRegex = regexp.MustCompile("^" +
	"(" + typeData + "|" + typeError + "):" +
	"(" + cmdRegistered + "|" + cmdReady + "|" + cmdOpen + "|" + cmdResult + "|" + cmdEnemy + ")" +
	"(:.*)*$")

var logger = log.New(os.Stdout, "[entrapped]", log.Ldate|log.Ltime|log.Lshortfile)

func randomInt(max int) int {
	if max <= 0 {
		max = 1
	}

	num, numErr := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if numErr != nil {
		return 0
	}

	return int(num.Int64())
}

func decodeJSON(req *http.Request, dst interface{}) error {
	return json.NewDecoder(req.Body).Decode(dst)
}

func error500(rw http.ResponseWriter, err error) {
	rw.WriteHeader(500)
	rw.Write([]byte(err.Error()))
}

type data struct {
	dataType string
	command  string
	params   map[string]string
}

func reqParser(msg string) (*data, string) {
	valid := dataRegex.FindStringSubmatch(msg)
	if len(valid) < 3 {
		return nil, "error:invalid params"
	}

	var params map[string]string
	if len(valid) > 3 {
		params = make(map[string]string)
		vals := strings.Split(valid[3], ":")
		for i := 1; i < len(vals); i++ {
			valArr := strings.Split(strings.Replace(strings.Replace(vals[i], "[", "", 1), "]", "", 1), "=")
			if len(valArr) != 2 {
				return nil, "error:invalid params"
			}

			params[valArr[0]] = valArr[1]
		}
	} else {
		params = nil
	}

	return &data{
		dataType: valid[1],
		command:  valid[2],
		params:   params,
	}, ""
}
