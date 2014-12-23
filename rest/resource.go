package rest

import (
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/util"
)

type Resource struct {
	object interface{}
}

func (resource Resource) Form(action Action) html.Node {
	return html.Div()
}

func (resource Resource) Name() string {
	valueReflection := reflect.ValueOf(resource.object)
	fullyQualifiedType := valueReflection.Type().String()
	packageComponents := strings.Split(fullyQualifiedType, ".")
	preciseType := packageComponents[len(packageComponents)-1]

	resourceName := util.CamelToUnderscore(preciseType)
	pluralizedResourceName := inflector.Pluralize(resourceName)

	return pluralizedResourceName
}
