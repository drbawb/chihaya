// Copyright 2013 The Chihaya Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

// Package babou provides a driver for a BitTorrent tracker to interface
// with the postgres database used by babou (github.com/drbawb/babou).
package babou

import (
	"database/sql"
	"fmt"

	"github.com/pushrax/chihaya/config"
	"github.com/pushrax/chihaya/storage/web"

	bridge "github.com/drbawb/babou/bridge"
	libBabou "github.com/drbawb/babou/lib"

	_ "github.com/bmizerany/pq"
)

type driver struct{}

func (d *driver) New(conf *config.DataStore) web.Conn {
	// The instance caches will initially be populated using
	// data from babou's database.
	//
	// Babou's database is optimized for a read-heavy workflow.
	// Writes to the database will be grouped into batches.
	//
	// The batches have dynamically adjusting timers so that
	// the write queues do not need to fill up before being
	// persisted to disk.
	//
	// It is important that the babou storage driver has a shutdown
	// hook so that this write queue can be safely persisted.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Password,
		conf.Schema,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("babou: failed to open connection to postgres")
	}

	if conf.MaxIdleConns != 0 {
		db.SetMaxIdleConns(conf.MaxIdleConns)
	}

	// Listen for messages from other babous.
	//
	// Babous are designed to work together as a `pack`.
	// This pack communicates by passing messages on what is known
	// as the `event bridge`.
	//
	// This driver will act as a "backend" on the event bridge.
	// It will subscribe to messages from the frontend which may
	// require writes or evictions of the chihaya cache.

	// TODO: Read from configuration file
	br0Settings := &libBabou.TransportSettings{}
	br0Settings.Transport = libBabou.TCP_TRANSPORT
	br0Settings.Socket = "0.0.0.0"
	br0Settings.Port = 5000

	br0 := bridge.NewBridge(br0Settings)

	return &Conn{db, br0}
}

type Conn struct {
	*sql.DB
	*bridge.Bridge
}

// Start's the event-bridge listener.
func (c *Conn) Start() error {
	return nil
}

func (c *Conn) RecordAnnounce(delta *web.AnnounceDelta) error {
	return nil
}

func (c *Conn) Close() error {
	return nil
}

func init() {
	web.Register("babou", &driver{})
}
