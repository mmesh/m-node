package update

var RestartRequest = make(chan struct{})
var RestartReady = make(chan struct{})
