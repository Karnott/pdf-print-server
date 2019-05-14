# PDF print server

It's a web server permitting to convert an HTML file to a PDF file.

## Endpoint

```
[POST] /pdf
```

#### Params

The only param accepted is an html file (as binary data)

#### Response

Return a PDF file


## How to build it ?

If you want to use it in a development env, you can use the `Dockerfile.dev`

```bash
docker build -t [your-image-name] . -f Dockerfile.dev
```

After that, you can run the image:

```bash
docker run --rm -p 4000:80 -v $(pwd):/go/src/github.com/karnott/pdf-print-server [your-image-name] go run main.go
```

When the server is running, you can try to convert an html file to a pdf file

```bash
curl -X POST http://localhost:4000/pdf --data-binary "@path-to-html-file" -o test.pdf
```

For production env, you can use the `Dockerfile` file

