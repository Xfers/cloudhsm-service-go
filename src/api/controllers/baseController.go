package controllers

// baseData is data that is stored once on server load and holds keys
type baseData struct {
	Keys *map[string]interface{}
}

// Base controller
// Only meant to be instantiated once as singleton when setting the server
type baseController struct {
	data *baseData
}

var instantiated *baseController = nil

func NewBaseController(keys *map[string]interface{}) *baseController {
	if instantiated == nil {
		instantiated = &baseController{
			data: &baseData{
				Keys: keys,
			},
		}
	}
	return instantiated
}

func GetBaseController() *baseController {
	return instantiated
}

func (c *baseController) getKey(name string) interface{} {
	return (*c.data.Keys)[name]
}
