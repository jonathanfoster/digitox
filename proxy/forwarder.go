package proxy

import (
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Forwarder represents a forwarding proxy.
type Forwarder struct {
}

// NewForwarder creates a new instance of Forwarder.
func NewForwarder() *Forwarder {
	return &Forwarder{}
}

func (f *Forwarder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("%v %v %v", r.Method, r.URL, r.Proto)

	if r.Method == http.MethodConnect {
		f.serveTunnel(w, r)
	} else {
		f.serveHTTP(w, r)
	}
}

func (f *Forwarder) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// ReverseProxy already handles hop-by-hop headers
	// so instead of rewriting that logic, configure the
	// reverse proxy as a forward proxy by setting the
	// proxied request URL to the original request URL
	rp := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			target := r.URL
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path
			req.URL.RawQuery = target.RawQuery
		},
	}

	rp.ServeHTTP(w, r)
}

func (f *Forwarder) serveTunnel(w http.ResponseWriter, r *http.Request) {
	// Default to port 80 if none provided
	addr := r.Host
	if !strings.Contains(addr, ":") {
		addr = addr + ":80"
	}

	dst, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorf("error dialing host %v: %v", addr, err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		log.Error("error hijacking connection: hijacking not supported: ", err.Error())
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	src, _, err := hijacker.Hijack()
	if err != nil {
		log.Error("error hijacking connection: ", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	dstTCP := dst.(*net.TCPConn)
	srcTCP := src.(*net.TCPConn)

	go transfer(dstTCP, srcTCP)
	go transfer(srcTCP, dstTCP)
}

func transfer(dst, src *net.TCPConn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Error("error copying bytes from source to destination: ", err.Error())
	}

	dst.CloseWrite() // nolint: errcheck, gas
	src.CloseRead()  // nolint: errcheck, gas
}
