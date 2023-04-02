package v1

import "github.com/kataras/iris/v12/core/router"

func RegisterConfigRouter(party router.Party) {
	party.Handle("GET", "/", nil)
}
