package main

import (
	"github.com/kataras/iris/v12"
	"io"
	"os"
)

const maxSize = 5 << 20 // 5MB

func main() {
	app := iris.New()

	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		// Get the file from the request.
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile("../uploads/"+fname,
			os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()

		io.Copy(out, file)
	})

	// start the server at http://localhost:8080 with post limit at 5 MB.
	app.Run(iris.Addr(":8080") /* 0.*/, iris.WithPostMaxMemory(maxSize))
}
