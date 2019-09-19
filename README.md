# Introduction:

  GoInfo is get os platform information coding by Golang.

  It can help you to know os information.


## Version:

version:0.0.1

## Futures

get linux information

get windows information

get osx information

get freebsd information


## Install:

```sh
$ go get github.com/charz/goInfo
$ go build
```

## Struct:

```sh
type GoInfoObject struct {
	GoOS string
	Kernel string
	Core string
	Platform string
	OS string
	Hostname string
	CPUs int
}
```

## Example:

```sh

$ go run cmd/main.go
```

On Linux:
```sh
    GoOS: linux
    Kernel: Linux
    Core: 4.15.0-52-generic
    Platform: x86_64
    OS: GNU/Linux (ubuntu 16.04)
    Hostname: ubuntu
    CPUs: 1
```

On Windows:

```sh
    GoOS: windows
    Kernel: windows
    Core: 10.0.15063 N/A Build 15063
    Platform: x64-based PC
    OS: Microsoft Windows 10 Pro
    Hostname: windows10
    CPUs: 4
```

On Mac:

```sh
    GoOS: darwin
    Kernel: Darwin
    Core: 18.7.0
    Platform: x86_64
    OS: Darwin
    Hostname: mac
    CPUs: 4
```

##License and Copyright
This software is Copyright 2012-2014 Matis Hsiao.

