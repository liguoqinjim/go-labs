package main

import (
	"github.com/kataras/iris/v12"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	app := iris.New()

	// Handle the post request from the upload_form.html to the server.
	app.Post("/upload", func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// second argument is totally optionally,
		// it can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		ctx.UploadFormFiles("../uploads", beforeSave)
	})

	app.Post("/upload_manual", func(ctx iris.Context) {
		// Get the max post value size passed via iris.WithPostMaxMemory.
		maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

		err := ctx.Request().ParseMultipartForm(maxSize)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}

		form := ctx.Request().MultipartForm

		files := form.File["files[]"]
		failures := 0
		for _, file := range files {
			_, err = saveUploadedFile(file, "../uploads")
			if err != nil {
				failures++
				ctx.Writef("failed to upload: %s\n", file.Filename)
			}
		}
		ctx.Writef("%d files uploaded", len(files)-failures)
	})

	// start the server at http://localhost:8080 with post limit at 32 MB.
	app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(32<<20))
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}