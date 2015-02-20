package css

import "fmt"

type Class string

func (class Class) Selector() string {
	return "." + string(class)
}

func (class Class) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: class, PseudoClass: pseudoClass}
}

func (class Class) Style(properties ...Property) RuleSet {
	return For(class).Set(properties...)
}

func (class Class) Child(childClass string) Class {
	return Class(fmt.Sprintf("%v__%v", string(class), childClass))
}

func (class Class) Variant(variantClass string) Class {
	return Class(fmt.Sprintf("%v--%v", string(class), variantClass))
}
