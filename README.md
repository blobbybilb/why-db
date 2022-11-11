# why?db
### An extremely simple and lightweight key-value store. Written in [go](https://github.com/golang/go) using [fiber](https://github.com/gofiber/fiber), because it's *fast*.

## Why?

Because I needed something simple to store/access data on a local server (for a project), and a proper database is too complicated.

Basically, I wanted to be able to do this:

```bash
go run main.go  # Running on X.X.X.X:5000
```

```python
requests.get(f"http://X.X.X.X:5000/set/example/{key}/{data}")

requests.get(f"http://X.X.X.X:5000/get/example/{key}").text  # data
```

## Usage

It's very, very simple.

1. Download, extract, open folder.

2. Run main.go:  `go run .`

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
