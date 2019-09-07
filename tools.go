// +build tools

package benchserder

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/mailru/easyjson/easyjson"
	_ "gopkg.in/src-d/proteus.v1/cli/proteus"
)
