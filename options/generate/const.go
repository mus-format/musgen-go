package genops

const (
	msigLastParamStream = "w mus.Writer"
	msigLastParam       = "bs []byte"
	mLastParamStream    = "w"
	mLastParamFirst     = "bs"
	mLastParam          = "bs[n:]"

	usigLastParamStream = "r mus.Reader"
	usigLastParam       = "bs []byte"
	uLastParamStream    = "r"
	uLastParamFirst     = "bs"
	uLastParam          = "bs[n:]"

	sksiqLastParamStream = "r mus.Reader"
	sksiqLastParam       = "bs []byte"
	skLastParamStream    = "r"
	skLastParam          = "bs[n:]"

	modImportNameStream  = "mus"
	modImportName        = "mus"
	extPackageNameStream = "ext"
	extPackageName       = "ext"
)
