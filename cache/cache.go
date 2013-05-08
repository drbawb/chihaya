package cache

import (
	m "github.com/kotokoko/chihaya/models"
	"github.com/kotokoko/storage"
)

type Cache interface {
	Load(s Storage) error

	FindTorrentByInfoHash(infohash string) (*m.Torrent, error)
	FindUserByPasskey(passkey string) (*m.User, error)
	PeerWhitelisted(peerId *m.Peer) (bool, error)

	SaveTorrent(t *m.Torrent) error
	SaveUser(u *m.User) error

	m.StatCollector
}
