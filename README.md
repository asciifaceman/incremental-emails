# Incremental Emails

Loosely developing a weird Golang incremental game using the Fyne.io UI ecosystem.


# Design

* [Game Design Doc](design/GDD.md)

# Build

Fyne uses cgo due to its low level graphics handling. You will require some dependencies to compile or xcompile.
This project is intended to be built on Linux or OSX which is what the build automation targets. I don't use windows so...

* Ubuntu/Linux
  * see [ubuntu-build-requirements.sh](ubuntu-build-requirements.sh)
* OSX
  * tbd
* Windows
  * tbd