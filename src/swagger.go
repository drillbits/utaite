package app

import (
	"reflect"
	"strings"

	"github.com/favclip/ucon/swagger"
)

func NewSwaggerPlugin() *swagger.Plugin {
	swInfo := &swagger.Info{
		Title:   "utaite",
		Version: "v1",
	}

	return swagger.NewPlugin(&swagger.Options{
		Object: &swagger.Object{
			Info:    swInfo,
			Schemes: []string{"http"},
		},
		DefinitionNameModifier: func(refT reflect.Type, defName string) string {
			if strings.HasSuffix(defName, "JSON") {
				return defName[:len(defName)-len("JSON")]
			}
			return defName
		},
	})
}
