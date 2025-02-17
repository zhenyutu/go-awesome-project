package geecache

/**
 * Peer
 */
type PeerGetter interface {
	PeerGet(group string, key string) ([]byte, error)
}

type PeerPicker interface {
	Pick(key string) (PeerGetter, error)
}
