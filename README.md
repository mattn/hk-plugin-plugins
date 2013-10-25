# hk-plugin-plugins

List installed plugins for [hk](https://github.com/heroku/hk)

# Installation

1. Add HKPATH environment into your .bashrc or something others.

  If you are using unix, `/usr/local/lib/hk/plugin` is default.

2. Create plugins directory

        $ mkdir -p $HKPATH

3. Install sql command

        $ git clone https://github.com/mattn/hk-plugin-plugins
        $ cd hk-plugin-plugins
        $ go build plugins.go
        $ cp plugins $HKPATH

# Usage

    $ cd /path/to/your/heroku/project
    $ hk plugins

# License

MIT

# Author

Yasuhiro Matsumoto
