package constant

const (

	// SeverStarted app startup logs
	SeverStarted           = "########### SERVER STARTED ################"
	StoppingServer         = "########### STOPPING SERVER ################"
	SeverStoppedGracefully = "########### SERVER STOPPED GRACEFULLY ################"
	AbortSignal            = "############### GOT %s SIGNAL. ABORTING.. ###############"
	DBConnectionSuccess    = "SUCCESSFULLY CONNECTED TO DATABASE: "
	DBConnectionFail       = "FAILED TO CONNECT TO DATABASE: "

	StatusOK  = 0
	StatusErr = 1

	PageSize = 4
)
