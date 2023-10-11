package enums

type StatusType string

const (
	TODO        StatusType = "TODO"
	IN_PROGRESS StatusType = "IN_PROGRESS"
	DONE        StatusType = "DONE"
	CANCELLED   StatusType = "CANCELLED"
)
