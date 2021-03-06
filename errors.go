package togo

import (
	"fmt"
	"reflect"
)

// CacheAddError is an error type describing an error while adding an element to a cache
type CacheAddError struct {
	cacheName   string
	elementName string
}

func (c CacheAddError) Error() string {
	return fmt.Sprintf("Error in caching %s to cache %s", c.elementName, c.cacheName)
}

// GenericCacheError is a generic error in handling of the cache as verified by the message
type GenericCacheError struct {
	cacheName string
	element   string
	message   string
}

func (c GenericCacheError) Error() string {
	return fmt.Sprintf("Error while working with cache %s and element %s: %s",
		c.cacheName, c.element, c.message)
}

// NameConflictError denotes the error thrown when a name is conflicting,
// and new name cannot be created
type NameConflictError struct {
	name string
}

func (nce NameConflictError) Error() string {
	return fmt.Sprintf("Name %s is already taken and cannot successfully create a new name",
		nce.name)
}

// FieldError indicates that the Field operation has failed due to some issue.
type FieldError struct {
	field   Field
	message string
}

func (fe FieldError) Error() string {
	return fmt.Sprintf("Field %s errored due to %s", fe.field.name, fe.message)
}

// UnsupportedType is an error type whenever an unsupported data type is encountered
// during creation of a Field or GoStruct
type UnsupportedType struct {
	data interface{}
}

func (ut UnsupportedType) Error() string {
	tp := reflect.TypeOf(ut.data).Kind()
	return fmt.Sprintf("Data type %s is not recognized", tp)
}

// GoStructError indicates that operation on the GoStruct failed due to some issue.
type GoStructError struct {
	gs      GoStruct
	message string
}

func (gse GoStructError) Error() string {
	return fmt.Sprintf("GoStruct %s errored due to %s", gse.gs.Name, gse.message)
}
