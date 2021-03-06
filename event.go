package highway

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/nzmprlr/highway/toll"
)

func Event(toll *toll.Toll) {
	if r := recover(); r != nil {
		toll.SetStatus(http.StatusInternalServerError)
		toll.Label("panic", fmt.Sprintf("%s\n%s", r, debug.Stack()))
	}

	toll.End()
}
