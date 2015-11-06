package log

type nilContext int

var NilContext = nilContext(0)

func (ctx nilContext) ContextMap() map[string]interface{} {
	return map[string]interface{}{}
}

// MapContext is a simple map based context
type MapContext map[string]interface{}

func (ctx MapContext) ContextMap() map[string]interface{} {
	return ctx
}
