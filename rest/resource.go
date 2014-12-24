package rest

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/util"
)

type Resource struct {
	object interface{}
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

func (resource Resource) FormFor(action Action) html.Node {
	return resource.FormForWithWrappedInputs(action, func(inputNode html.Node) html.Node {
		return inputNode
	})
}

func (resource Resource) FormForWithWrappedInputs(action Action, wrapper func(html.Node) html.Node) html.Node {
	var inputNodes []html.Node
	innerValue := reflect.ValueOf(resource.object)
	valueType := innerValue.Type()

	if valueType.Kind() != reflect.Struct {
		panic("Not a struct resource")
	}

	for i := 0; i < innerValue.NumField(); i++ {
		field := innerValue.Field(i)

		resourceName := resource.Name()
		fieldName := util.CamelToUnderscore(valueType.Field(i).Name)
		inputName := fmt.Sprintf("%v[%v]", inflector.Singularize(resourceName), fieldName)
		inputNode := html.Input().Name(inputName).Value(fmt.Sprintf("%v", field.Interface()))

		wrappedNode := wrapper(inputNode)
		inputNodes = append(inputNodes, wrappedNode)
	}

	submitNode := html.Input().Type("submit").Value(strings.Title(string(action)))
	inputNodes = append(inputNodes, submitNode)

	return html.Form().Children(inputNodes...)
}
