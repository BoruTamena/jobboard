package test

import "github.com/BoruTamena/jobboard/initiator"

// your test initiator goes here
type TestInstance struct {
	/*

		Server
		Cache
		Module
		Handler
		PlatformLayer

	*/

	Handler       initiator.Handler
	Module        initiator.Module
	PlatformLayer initiator.Platform
}

func InitiateTest(arg string) TestInstance {

	return TestInstance{}
}
