package controller

import (
	"cd-pipeline-handler-controller/pkg/controller/stage"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, stage.Add)
}
