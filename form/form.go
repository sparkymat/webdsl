package form

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/sparkymat/nadeshiko/html"
	"github.com/sparkymat/webdsl/form/rest"
	"github.com/sparkymat/webdsl/util"
)

type Form struct {
	action rest.Action
	value  interface{}
}

func (form Form) Html() (html.Node, error) {
	var inputNodes []html.Node

	innerValue := reflect.ValueOf(form.value)
	valueType := innerValue.Type()

	for i := 0; i < innerValue.NumField(); i++ {
		field := innerValue.Field(i)

		resourceName, err := form.resourceName()

		if err != nil {
			return html.Node{}, err
		}

		fieldName, err := util.CamelToUnderscore(valueType.Field(i).Name)

		if err != nil {
			return html.Node{}, err
		}

		inputName := fmt.Sprintf("%v[%v]", inflector.Singularize(resourceName), fieldName)

		inputNode := html.Input().Name(inputName).Value(fmt.Sprintf("%v", field.Interface()))

		//log.Print("field = ", valueType.Field(i).Name)
		//log.Print("type = ", field.Type())

		inputNodes = append(inputNodes, inputNode)
	}

	return html.Form().Children(inputNodes...), nil
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
