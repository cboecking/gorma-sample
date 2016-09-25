//************************************************************************//
// API "cellar-aaa": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/cboecking/gorma-sample/design
// --out=$(GOPATH)src/github.com/cboecking/gorma-sample
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// AccountbbbHref returns the resource href.
func AccountbbbHref(accountbbbID interface{}) string {
	paramaccountbbbID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountbbbID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar-aaa/accountbbbs/%v", paramaccountbbbID)
}
