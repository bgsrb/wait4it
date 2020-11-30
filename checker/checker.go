package checker

import (
	"wait4it/config"
)

//Checker ...
type Checker interface {
	BuildContext(cx config.CheckContext)
	Validate() error
	Check() (bool, bool, error)
}
