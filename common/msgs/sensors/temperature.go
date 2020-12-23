package sensors

import (
	"github.com/tombenke/axon-go/common/msgs/common"
)

type Temperature struct {
	Header common.Header
	Body   common.Float64VarBody
}
