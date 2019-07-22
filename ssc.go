package ssc

func New () (*SSC) { // This function can be used to create new instances of this data type.
	return &SSC {false}
}

type SSC struct {
	shutdown bool /* True would mean shutdown has been signalled, false would mean shutdown is
		yet to be signal. */
}

func (ssc *SSC) WhatsUp () (bool) { // To check if a shutdown has been signalled, call this method.
	return ssc.shutdown
}

func (ssc *SSC) Shutdown () { /* To signal shutdown, call this method. Once a shutdown has been
	signalled, it can not be undone. */

	ssc.shutdown = true
}
