package httputils

const (
	HeaderLastModified            = "Last-Modified"
	HeaderContentType             = "Content-Type"
	HeaderCacheControl            = "Cache-Control"
	HeaderContentLength           = "Content-Length"
	HeaderETag                    = "ETag"
	HeaderServer                  = "Server"
	HeaderIfModifiedSince         = "If-Modified-Since"
	HeaderExpires                 = "Expires"
	HeaderIfNoneMatch             = "If-None-Match"
	HeaderAccept                  = "Accept"
	HeaderUserAgent               = "User-Agent"
	HeaderAuthorization           = "Authorization"
	HeaderAltSvc                  = "Alt-Svc"
	HeaderStrictTransportSecurity = "Strict-Transport-Security"
	HeaderContentDisposition      = "Content-Disposition"
)

const (
	CacheControlDynamic = "public, no-cache, must-revalidate" // TODO https://web.dev/http-cache/, https://web.dev/love-your-cache/
	// CacheControlDynamic    = "max-age=0, must-revalidate, public" // TODO https://web.dev/http-cache/, https://web.dev/love-your-cache/
	CacheControlWebappAssets = "max-age=31536000, public" // 1 year, theme assets SHOULD be versionned
	// CacheControlNoCache       = "no-cache, no-store, no-transform, must-revalidate" // "no-cache, no-store, no-transform, must-revalidate, private, max-age=0"
	CacheControlNoCache    = "no-cache, no-store, must-revalidate" // "no-cache, no-store, no-transform, must-revalidate, private, max-age=0"
	CacheControl5Minutes   = "max-age=300, public"
	CacheControl10Minutes  = "max-age=600, public"
	CacheControlPortal     = "max-age=60, s-max-age=3600, public"
	CacheControlWebappPage = CacheControlDynamic
	// CacheControlWebappJsCss = CacheControlThemeAsset
)

const (
	MediaTypeText = "text/plain"
	MediaTypeXml  = "application/xml"
	MediaTypeJson = "application/json"

	MediaTypeHtmlUtf8       = "text/html; charset=UTF-8"
	MediaTypeJsonUtf8       = "application/json; charset=utf-8"
	MediaTypeJavascriptUtf8 = "application/javascript;charset=utf-8"
	MediaTypeTextUtf8       = "text/plain; charset=UTF-8"

	MediaTypeJsonFeed = "application/feed+json"
	MediaTypeRSS      = "application/rss+xml"
	MediaTypeAtom     = "application/atom+xml"

	MediaTypeAudio = "audio"
	MediaTypeVideo = "video"
	MediaTypePNG   = "image/png"
	MediaTypeJPEG  = "image/jpeg"
	MediaTypeSVG   = "image/svg+xml"

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// alternative value from the 'file' command line program: text/x-shellscript
	MediaTypeShellScript = "application/x-sh"
)
