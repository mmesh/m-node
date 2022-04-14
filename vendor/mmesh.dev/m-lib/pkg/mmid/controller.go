package mmid

type MMControllerID string

func NewMMControllerID(controllerID string) MMControllerID {
	return MMControllerID(controllerID)
}

func (mmid MMControllerID) String() string {
	return string(mmid)
}
