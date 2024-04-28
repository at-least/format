package format

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

var cases = []string{
	"Api",
	"Id",
	"Http",
	"Https",
	"Pdf",
	"Ip",
	"Json",
	"Sql",
	"Vat",
	"Tcp",
	"Tls",
	"Udp",
	"Ui",
	"Uid",
	"Uuid",
	"Uri",
	"Url",
	"Utf8",
	"Pk",
	"Sk",
}

// GoName returns a name formatted for Go.
func GoName(s string) string {
	s = strcase.ToCamel(s)
	for _, c := range cases {
		if strings.HasSuffix(s, c) {
			s = strings.Replace(s, c, strings.ToUpper(c), 1)
		}
	}
	return s
}

// GoInputType returns the name of a method input type
func GoInputType(types, method string) string {
	if len(types) == 0 {
		return fmt.Sprintf("%sInput", GoName(method))
	}
	return fmt.Sprintf("%s.%sInput", types, GoName(method))
}

// JsName returns a name formatted for JS.
func JsName(s string) string {
	return strcase.ToLowerCamel(s)
}
