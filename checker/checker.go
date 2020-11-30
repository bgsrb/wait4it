package checker

import (
	"wait4it/model"
)

//Checker ...
type Checker interface {
	BuildContext(cx model.CheckContext)
	Validate() error
	Check() (bool, bool, error)
}
