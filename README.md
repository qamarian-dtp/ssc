# ssc: Shutdown-signification Channel

Shutdown-signification Channel: A channel (in form of a data type) for sending shutdown signal from
one gorountine to another. To use, get a new SSC with function ssc.New (), a pointer of a new SSC
would be returned.

~~~~
ssc := ssc.New ()
~~~~

Pass the _SSC_ to the goroutines that would be using it.

To signal shutdown via the SSC, call its method Shutdown ().

~~~~
ssc.Shutdown ()
~~~~

Note: Once shutdown has been signalled via an SSC, the signal can not be called back.

To check if shutdown has been signalled via an SSC, call method WhatsUp ().

~~~~
signalled := ssc.WhatsUp ()
~~~~

If shutdown has been signalled via the SSC, _true_ would be returned. Otherwise, _false_ would be
returned.
