package form

import (
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/sparkymat/webdsl/form/rest"
	"github.com/sparkymat/webdsl/util"
)

type Form struct {
	action rest.Action
	value  interface{}
}

func (form Form) resourceName() (string, error) {
	valueReflection := reflect.ValueOf(form.value)
	fullyQualifiedType := valueReflection.Type().String()
	components := strings.Split(fullyQualifiedType, ".")
	preciseType := components[len(components)-1]

	resourceName, err := util.CamelToUnderscore(preciseType)

	if err != nil {
		return "", err
	}

	pluralizedResourceName := inflector.Pluralize(resourceName)

	return pluralizedResourceName, nil
}
