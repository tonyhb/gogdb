package gui

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/tonyhb/gogdb/gui/assets"
)

const ROOT = "inspector.html"

func Run(c *cli.Context) (err error) {
	http.HandleFunc("/", serveFile)

	if c.Bool("server") {
		return serveHTTP()
	}

	return serveNodeWebkit()
}

// Serve devtools using node-webkit as a standalone GUI app
func serveNodeWebkit() (err error) {
	var wk nw.NodeWebkit
	if wk, err = nw.New(); err != nil {
		return
	}
	logrus.Debug("starting node-webkit")
	return wk.ListenAndServe(nil)
}

// Serves a requested devtool file from precompiled binary data.
func serveFile(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		byt      []byte
		filename = r.RequestURI[1:]
	)

	if filename == "" {
		filename = ROOT
	}

	if byt, err = assets.Asset(filename); err != nil {
		logrus.WithFields(logrus.Fields{"url": r.URL, "filename": filename}).Warn("404")
		w.WriteHeader(404)
		return
	}

	mime := getMime(filename)
	w.Header().Set("content-type", mime)
	w.Write(byt)

	logrus.WithFields(logrus.Fields{
		"url":      r.URL,
		"filename": filename,
		"mime":     mime,
	}).Debug("serving request")

}

// Serve devtools using golang's built in HTTP server
// This is primarily used in testing
func serveHTTP() (err error) {
	logrus.Debug("starting http server")
	return http.ListenAndServe(":8080", http.DefaultServeMux)
}
