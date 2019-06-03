package object

import (
	"github.com/project-flogo/legacybridge"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "tibco-log",
  "type": "flogo:activity",
  "ref": "github.com/lixingwang/flogo/activity/object",
  "version": "0.0.1",
  "title": "Get Ojbect",
  "description": "Get object value using mapper path",
  "homepage": "https://github.com/lixingwang/flogo/tree/master/activity/object",
  "inputs":[
    {
      "name": "object",
      "type": "any",
      "value": ""
    },
    {
      "name": "path",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "value",
      "type": "any"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	legacybridge.RegisterLegacyActivity(NewActivity(md))
}
