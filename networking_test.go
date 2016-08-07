package main

type mockNetworking struct {
	recv chan (*message)
	send chan (*message)
}

func newMockNetworking() *mockNetworking {
	net := &mockNetworking{}
	return net
}

func (net *mockNetworking) listen() error {
	return nil
}

func (net *mockNetworking) disconnect() error {
	return nil
}

func (net *mockNetworking) createSocket(host string, port string) error {
	return nil
}

func (net *mockNetworking) cancelResponse(*expectedResponse) {
}

func (net *mockNetworking) init(self *NetworkNode) {
	net.recv = make(chan (*message))
	net.send = make(chan (*message))
}

func (net *mockNetworking) messagesFin() {

}

func (net *mockNetworking) timersFin() {

}

func (net *mockNetworking) getDisconnect() chan (int) {
	return nil
}

func (net *mockNetworking) getMessage() chan (*message) {
	return nil
}

func (net *mockNetworking) sendMessage(q *message, id int64, expectResponse bool) (*expectedResponse, error) {
	net.recv <- q
	if expectResponse {
		return &expectedResponse{ch: net.send, query: q, node: q.Receiver, id: id}, nil
	} else {
		return nil, nil
	}
}