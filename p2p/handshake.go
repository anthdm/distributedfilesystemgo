package p2p

// HandshakeFunc... ?
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
