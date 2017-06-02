ShadInfo
========

A Go front-end to display a server’s system info.

“Quick” install
---------------

*Right now*, it is necessary to have one of the most up-to-date versions of Go (1.8)
to run `shadinfo`. **On Debian**, it means [downloading a binary](https://golang.org/doc/install).

```sh
git clone https://github.com/Diti/shadinfo.git /var/www/shadinfo
cd $_
go get -d ./... # Download the dependencies into $(go env GOPATH)/src
go build # Make ./shadinfo
```

Usage
-----

By default, ShadInfo will look for a [Go Template](https://golang.org/pkg/text/template/)
file called `index.html.tmpl` located in your shell’s current “working directory”.
You will probably need to tell `shadinfo` where to find that template:

```sh
shadinfo -template /path/to/index.html.tmpl
```

By default, `shadinfo` will bind on `:8080`, which means it will be listening for
all connections on this port. If you plan on running this server behind a reverse
proxy, I suggest binding to local connections only, and on another port:

```sh
shadinfo -bind localhost:31415
```

From there, you may now write a service to run it at boot, and add `shadinfo` to
your reverse proxy server. On my [home server](http://www.dell.com/dfb/p/optiplex-fx160/pd)
running [Caddy](https://caddyserver.com/), I use these `Caddyfile` settings:

```nginx
home.diti.me,
shade.home.diti.me,
http://109.190.114.172,
http://192.168.1.1 {
    proxy / localhost:31415 {
        transparent
    }
}
```
