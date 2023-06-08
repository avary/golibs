package httputils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/skerkour/golibs/slogutil"
)

var ErrDir = errors.New("path is a folder")
var ErrInvalidPath = errors.New("path is not valid")
var ErrInternalError = errors.New("Internal Server Error")

func WebappHandler(folder fs.FS) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := slogutil.FromCtx(ctx)

		err := tryRead(folder, r.URL.Path, w, r)
		if err == nil {
			logger.Error("httputils.WebappHandler: reading file", slogutil.Err(err))
			handleError(http.StatusInternalServerError, ErrInternalError.Error(), w)
			return
		}

		err = tryRead(folder, "index.html", w, r)
		if err != nil {
			logger.Error("httputils.WebappHandler: reading index.html", slogutil.Err(err))
			handleError(http.StatusInternalServerError, ErrInternalError.Error(), w)
			return
		}
	}
}

func handleError(code int, message string, w http.ResponseWriter) {
	http.Error(w, message, code)
}

func tryRead(fs fs.FS, path string, w http.ResponseWriter, r *http.Request) (err error) {
	path = filepath.Clean(path)
	if strings.Contains(path, "..") || path == "" {
		err = ErrInvalidPath
		return
	}

	if path[0] == '/' {
		path = path[1:]
	}

	file, err := fs.Open(path)
	if err != nil {
		err = ErrInvalidPath
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	if fileInfo.IsDir() {
		err = ErrDir
		return
	}

	modifiedAt := fileInfo.ModTime().UTC().Truncate(time.Second)
	modifiedAtStr := modifiedAt.UTC().Format(http.TimeFormat)

	extension := filepath.Ext(path)

	if extension == ".js" || extension == ".css" || extension == ".woff" || extension == ".woff2" {
		// some webapp's assets files can be cached for very long time because they are versionned by
		// webapp bundler
		w.Header().Set(HeaderCacheControl, CacheControlWebappAssets)
	} else {
		w.Header().Set(HeaderCacheControl, CacheControlDynamic)
	}
	w.Header().Set(HeaderLastModified, modifiedAt.Format(http.TimeFormat))
	w.Header().Set(HeaderETag, fmt.Sprintf(`"%s"`, base64.StdEncoding.EncodeToString([]byte(modifiedAtStr))))

	// handle caching
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		ifModifiedSinceHeader := strings.TrimSpace(r.Header.Get(HeaderIfModifiedSince))
		if ifModifiedSinceHeader != "" {
			ifModifiedSinceTime, errModifiedSinceheader := http.ParseTime(ifModifiedSinceHeader)
			if errModifiedSinceheader == nil {
				ifModifiedSinceTime = ifModifiedSinceTime.Truncate(time.Second)
				if ifModifiedSinceTime.Equal(modifiedAt) || modifiedAt.Before(ifModifiedSinceTime) {
					w.WriteHeader(http.StatusNotModified)
					return
				}
			}
		} else {
			// otherwise we try to fallback to etag which is a base64 encoded version of ifModifiedSinceTime
			// (see below)
			etagHeader := strings.TrimSpace(r.Header.Get(HeaderIfNoneMatch))
			if etagHeader != "" {
				etagHeader = strings.TrimPrefix(etagHeader, "W/")
				etagHeader = strings.TrimLeft(etagHeader, `"`)
				etagHeader = strings.TrimRight(etagHeader, `"`)
				etagValue, errBase64Decode := base64.StdEncoding.DecodeString(etagHeader)
				if errBase64Decode == nil {
					ifModifiedSinceTime, errParseTime := http.ParseTime(string(etagValue))
					if errParseTime == nil {
						ifModifiedSinceTime = ifModifiedSinceTime.Truncate(time.Second)
						if ifModifiedSinceTime.Equal(modifiedAt) || modifiedAt.Before(ifModifiedSinceTime) {
							w.WriteHeader(http.StatusNotModified)
							return
						}
					}

				}
			}
		}
	}

	contentType := mime.TypeByExtension(extension)
	w.Header().Set(HeaderContentType, contentType)

	_, err = io.Copy(w, file)
	if err != nil {
		err = fmt.Errorf("httputils.tryRead: copying content to HTTP response: %w", err)
		return
	}

	return
}
