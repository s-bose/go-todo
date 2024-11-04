package enums

type StatusType string
type PriorityType string

const (
	TODO        StatusType = "TODO"
	IN_PROGRESS StatusType = "IN PROGRESS"
	DONE        StatusType = "DONE"
	CANCELLED   StatusType = "CANCELLED"

	HIGH     PriorityType = "HIGH"
	MEDIUM   PriorityType = "MEDIUM"
	LOW      PriorityType = "LOW"
	VERY_LOW PriorityType = "VERY LOW"
)
