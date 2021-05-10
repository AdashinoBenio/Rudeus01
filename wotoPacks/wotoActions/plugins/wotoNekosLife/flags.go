package wotoNekosLife

// the flags for the wotoNekosLife plugin.
// please notice that all flags should begin with
// prefex `--`, but in the pTools, this prefex will be
// removed.
const (
	PV_FLAG      = "pv"      // won't work if message is replied
	ADD_FLAG     = "add"     // it works only for sudo people
	RM_FLAG      = "rm"      // it works only for sudo people
	REM_FLAG     = "rem"     // it works only for sudo people
	REMOVE_FLAG  = "remove"  // it works only for sudo people
	HENTAI_FLAG  = "hentai"  // it works only for sudo people
	PRIVATE_FLAG = "private" // won't work if message is replied
	DELETE_FLAG  = "delete"  // won't work if message is not replied
	DEL_FLAG     = "del"     // won't work if message is not replied
)
