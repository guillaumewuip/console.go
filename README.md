# console.go [![godoc reference](https://godoc.org/github.com/guillaumewuip/console.go?status.png)](http://godoc.org/github.com/guillaumewuip/console.go)

Console is a lightweight logging tool inspired by the NodeJS console object
and [Scribe.js](http://github.com/bluejamesbond/scribe.js) project.

```golang
console.Info("Some info")
console.Tag("Hello").Time().File().Log("Hello World")
```

Console provided `Info`, `Log`, `Warning` and `Error` method and `Tag`, `Time`, `File` optional information.

With use of [ansi](https://github.com/mgutz/ansi) colors, it output something like below in terminal :

```
[Tag1][Tag2] [main.go:48] 2015-02-15T09:22:06+01:00 Hello World
```

**Godoc :** http://godoc.org/github.com/guillaumewuip/console.go

