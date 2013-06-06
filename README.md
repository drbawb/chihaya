chihaya
=======

[![Build Status](https://travis-ci.org/pushrax/chihaya.png?branch=master)](https://travis-ci.org/pushrax/chihaya)

chihaya is a high-performance [BitTorrent tracker](http://en.wikipedia.org/wiki/BitTorrent_tracker) in Go.
It isn't yet ready for prime-time, but these are the features that it'll have:

- Linear horizontal scalability
- Low processing and memory footprint
- IPv6 support
- A publish/subscribe architecture for communication with other components
- Support for multiple database backends


Installing
----------

    $ go get github.com/pushrax/chihaya


Configuration
-------------

Configuration is done in `config.json`, which you'll need to create by copying
`config.json.example`. See [config/config.go](https://github.com/pushrax/chihaya/blob/master/config/config.go)
for a description of each configuration value.


Running
-------

`./chihaya` to run normally, `./chihaya -profile` to generate pprof data for analysis.


Contributing
------------

Style guide: `go fmt`.

If you want to make a smaller change, just go ahead and do it, and when you're
done send a pull request through GitHub. If there's a larger change you want to
make, it would be preferable to discuss it first via a GitHub issue or by
getting in touch on IRC.


Contact
-------

If you have any questions or want to contribute something, come say hi in the
IRC channel: **#chihaya on [freenode](http://freenode.net/)**
([webchat](http://webchat.freenode.net?channels=chihaya)).

