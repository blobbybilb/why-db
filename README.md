# why?db

### An extremely simple and lightweight HTTP key-value store. Written in [go](https://github.com/golang/go) using [fiber](https://github.com/gofiber/fiber), because it's _fast_.

#### It can also use a Postgres DB as storage, while still having the same nice & simple HTTP key-value store.

## Why?

Because I needed something simple to store/access data on a local server (for a project), without having to deal with an SQL or other database.

Basically, I wanted to be able to do this:

```bash
go run main.go  # Running on X.X.X.X:6010
```

```javascript
fetch(`http://X.X.X.X:6010/set/example/${key}/${data}`)

fetch(f"http://X.X.X.X:6010/get/example/{key}") // data as text body
```

Then I also thought that it'd be nice to be able to use an SQL DB as storage for the key-value store (mostly to allow using a cloud SQL DB), so now it can only use a Postgres DB as storage.

## Usage

It's very, very simple.

1. Download, extract, open folder.

2. `go mod tidy`, then `go run .`

3. Make http requests. Routes:

   ```markup
   set (replaces value):
   GET /set/<category>/<key>/<value>/
   POST /set/<category>/<key>/     (raw POST body is used as value)

   add (appends to value):
   GET /add/<category>/<key>/<value>/
   POST /add/<category>/<key>/     (raw POST body is used as value)


   get:
   GET /get/<category>/<key>/      (gets stored value as response text)


   del (deletes value):
   GET /del/<category>/<key>/
   ```

   ```bash
   # Examples:
   curl http://127.0.0.1:5000/set/example/somekey/somevalue/
   curl http://127.0.0.1:5000/get/example/somekey/  # somevalue
   ```

## Limitations

Yes.
