# BigConv - Simple Command Line Utility for Converting Big Numbers

A really simple command line tool written in Go to convert big hexadecimals to integers and back big ints to hexadecimals.

## Hex To Int

```sh
$ go run main.go hex a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
966482230667555116936258103322711973649032657875
```

## Int To Hex

```sh
$ go run main.go int 966482230667555116936258103322711973649032657875
a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
```

## How To Install

Something like this:

```
cd $GOPATH/src/
go get github.com/IcyApril/bigconv
cd github.com/IcyApril/bigconv/
go install
sudo cp $GOPATH/bin/bigconv /usr/local/bin/bigconv
```