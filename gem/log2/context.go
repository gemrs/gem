package log

type MapContext map[string]interface{}

func (ctx MapContext) ContextMap() map[string]interface{} {
	return ctx
}
