package log

// MapContext is a simple map based context
type MapContext map[string]interface{}

func (ctx MapContext) ContextMap() map[string]interface{} {
	return ctx
}
