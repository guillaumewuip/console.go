# console.go [![godoc reference](https://godoc.org/github.com/guillaumewuip/console.go?status.png)](http://godoc.org/github.com/guillaumewuip/console.go)

Console is a lightweight logging tool inspired by the NodeJS console object
and [Scribe.js](http://github.com/bluejamesbond/scribe.js) project.

*As a js dev, I was lost in the Go world without my `console.log`.*

```go
console.Info("Some info")
console.Tag("Hello").Time().File().Log("Hello World")
```

With use of [ansi](https://github.com/mgutz/ansi) colors, it output something like this in terminal :

```text
[Hello] [main.go:48] 2015-02-15T09:22:06+01:00 Hello World
```

----

Console provided :

- `Info`, `Log`, `Warning` and `Error` logging method
- `Tag`, `Time`, `File` optional method
- hook

**See Godoc :** http://godoc.org/github.com/guillaumewuip/console.go

**Related :**
- [a console.go hook to send log over AMQP](https://github.com/guillaumewuip/console_scribeAmqp.go)
