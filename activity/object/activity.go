package object

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/engine/mapping/ref/arrayref"
)

// log is the default logger for the Log Activity
var log = logger.GetLogger("activity-tibco-log")

const (
	ivObject = "object"
	ivPath   = "path"
	ovValue  = "value"
)

// ObjectActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type ObjectActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ObjectActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ObjectActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *ObjectActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	obj := context.GetInput(ivObject)
	path := context.GetInput(ivPath).(string)

	if obj == nil {
		return false, nil
	} else {
		v, err := data.CoerceToAny((obj)
		if err != nil {
			return false, fmt.Errorf("Object search must be a json object")
		}
		obj = v
	}

	if strings.HasPrefix(path, "$.") {
		arrayRef := arrayref.ArrayRef{path}
		data, err := arrayRef.EvalFromData(obj)
		if err != nil {
			return false, fmt.Errorf("Object search must be a json object")
		}
		context.SetOutput(ovValue, data)
		return true, nil

	}
		return false, fmt.Errorf("The path must start with $.")

}