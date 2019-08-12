geiger
------

[![GoDoc](https://godoc.org/github.com/lukechampine/geiger?status.svg)](https://godoc.org/github.com/lukechampine/geiger)
[![Go Report Card](http://goreportcard.com/badge/github.com/lukechampine/geiger)](https://goreportcard.com/report/github.com/lukechampine/geiger)

```
go get lukechampine.com/geiger
```

Inspired by [this tweet](https://twitter.com/laserallan/status/1159571592332087296).
Just call `go geiger.Count()` at the top of your `main` function. Don't worry, if
you start tasting metal, you're probably just delusional.

This package is in a pretty unpolished state. ~~First of all, it emits a constant
tone, not a series of clicks like a real Geiger counter, because I couldn't
figure out how to get the `beep` package to click the way I wanted.~~ (EDIT: [fixed, thanks @klauspost!](https://github.com/lukechampine/geiger/pull/1)) Second, the
package itself incurs a non-zero number of allocations per second (~100 on
MacOS, presumably different on other platforms) which throws off the reading.
Third, it's completely unconfigurable: you can't adjust the "sensitivity," and
you can only measure the number of objects allocated per second (as opposed to,
e.g., the number of *bytes* allocated per second).

Still, I think the general idea of "use sound to alert the programmer that
they're doing something stupid" has merit. If enough people agree, maybe we can
work together to refine it into an actually useful tool.
