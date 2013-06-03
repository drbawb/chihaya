chihaya
=======

[![Build Status](https://travis-ci.org/pushrax/chihaya.png?branch=master)](https://travis-ci.org/pushrax/chihaya)

Due to the many inconsistencies AB has with Gazelle, Chihaya is not ready for
general use. Currently the way Chihaya finds out about new and deleted data is
by polling the database server, which is highly inefficent and introduces a
race condition when data is deleted from the source
(due to `INSERT INTO x ON DUPLICATE KEY UPDATE` being used). A pub/sub
architecture is being developed now that will mitigate this.

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
IRC channel: [#chihaya on freenode](irc://chat.freenode.net:6667/chihaya)
([webchat](http://webchat.freenode.net?channels=chihaya)).

