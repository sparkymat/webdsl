package rest

import "github.com/sparkymat/webdsl/http"

type Action string

const Create Action = "create"
const New Action = "new"
const Update Action = "update"
const Edit Action = "edit"
const Index Action = "index"
const Show Action = "show"
const Destroy Action = "destroy"

func (action Action) Method() http.Method {
	switch action {
	case Create:
		return http.Put
	case New:
		return http.Get
	case Update:
		return http.Patch
	case Edit:
		return http.Get
	case Index:
		return http.Get
	case Show:
		return http.Get
	case Destroy:
		return http.Delete
	default:
		panic("unknown action")
	}
}
