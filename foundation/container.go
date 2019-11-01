package foundation

import (
	"lanvard/support"
	"reflect"
)

type bindings map[string]interface{}
type instances map[string]interface{}

type ContainerStruct struct {

	// The container's bindings.
	bindings bindings

	// The registered type aliases.
	aliases map[string]interface{}

	// The registered aliases keyed by the abstract name.
	abstractAliases map[string]map[string]interface{}

	// The container's shared instances.
	instances instances
}

func Container() ContainerStruct {
	return ContainerStruct{}
}

// Determine if the given abstract type has been bound.
func (c *ContainerStruct) Bound(abstract string) bool {
	_, bind := c.bindings[abstract]
	_, instance := c.instances[abstract]

	if bind || instance || c.IsAlias(abstract) {
		return true
	}

	return false
}

// Determine if a given string is an alias.
func (c *ContainerStruct) IsAlias(name string) bool {
	if _, ok := c.aliases[name]; ok {
		return true
	}

	return false
}

// Register a binding with the container.
func (c *ContainerStruct) Bind(abstract interface{}, concrete interface{}) {
	if c.bindings == nil {
		c.bindings = make(bindings)
	}
	abstractString := support.Name(abstract)

	c.bindings[abstractString] = concrete
}

// Register a shared binding in the container.
func (c *ContainerStruct) Singleton(abstract interface{}, concrete interface{}) {
	c.Bind(abstract, concrete)
}

// Register an existing instance as shared in the container.
func (c ContainerStruct) Instance(abstract interface{}, instance interface{}) {
	abstractName := support.Name(abstract)

	c.removeAbstractAlias(abstractName)

	_, ok := c.aliases[abstractName]
	if ok {
		delete(c.aliases, abstractName)
	}

	if c.instances == nil {
		c.instances = make(instances)
	}

	c.instances[abstractName] = instance
}

// Get the container's bindings.
func (c ContainerStruct) GetBindings() bindings {
	return c.bindings
}

// Resolve the given type from the container.
func (c *ContainerStruct) Make(abstract interface{}) interface{} {
	return c.resolve(abstract)
}

// Resolve the given type from the container.
func (c *ContainerStruct) resolve(abstract interface{}) interface{} {
	abstractString := reflect.TypeOf(abstract).Elem().String()

	object, present := c.bindings[abstractString]

	if present {
		return object
	}

	panic("Can't resole container")
}

// Remove an alias from the contextual binding alias cache.
func (c ContainerStruct) removeAbstractAlias(abstract string) {
	if _, ok := c.aliases[abstract]; !ok {
		return
	}

	panic("Todo, implement removeAbstractAlias")
}
