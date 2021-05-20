package wotoSudo

type sudoErr uint16

const (
	notInList     sudoErr = 1
	invalidId     sudoErr = 2
	alreadyInList sudoErr = 3
)
