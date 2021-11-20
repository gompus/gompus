package integration

//go:generate stringer -type=ExpireBehavior -trimprefix=ExpireBehavior

type ExpireBehavior int

const (
	ExpireBehaviorRemoveRole ExpireBehavior = 0
	ExpireBehaviorKick                      = 1
)
