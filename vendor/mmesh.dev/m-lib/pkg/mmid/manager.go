package mmid

type MMManagerID string

func NewMMManagerID(managerID string) MMManagerID {
	return MMManagerID(managerID)
}

func (mmid MMManagerID) String() string {
	return string(mmid)
}
