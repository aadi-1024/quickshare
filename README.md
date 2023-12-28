### Share files blazingly fast or smth

Quickly share files on the same network

Holy shit the AJAX request was quite the pain.

Installing:
```shell
$  git clone https://github.com/aadi-1024/quickshare
$  cd quickshare
$  go build cmd/*.go
$  sudo mv main /usr/local/bin/quickshare
```

Usage:
```shell
quickshare filename
```

This will display the authentication password and start listening
to requests on port 8080. You can access it from a browser
by going to `http://ipaddr:8080`

Not secure for now as:
- uses http
- the authentication can be easily bypassed, will be fixed by using a randomly generated endpoint instead of /file