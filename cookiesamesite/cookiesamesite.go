
package cookiesamesite

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"log"
	"os"
	"strings"
)

type Rewrite struct {
	Replacement string `json:"replacement,omitempty"`
}

type Config struct {
	Rewrites []Rewrite `json:"rewrites,omitempty"`
}

func CreateConfig() *Config {
	return &Config{}
}

type rewrite struct {
	replacement string
}

type rewriteBody struct {
	name     string
	next     http.Handler
	rewrites []rewrite
}

// New creates and returns a new rewrite body plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	log.SetOutput( os.Stdout ) ;
	log.Println( "LOADED" ) ;

	////////////////////////////////////////////////////////////////

	rewrites := make( [ ]rewrite , len( config.Rewrites ) )

	for i , rewriteConfig := range config.Rewrites {
		rewrites[ i ] = rewrite{
			replacement : rewriteConfig.Replacement ,
		}
	}

	return &rewriteBody {
		name :     name ,
		next :     next ,
		rewrites : rewrites ,
	} , nil

}

func (r *rewriteBody) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	wrappedWriter := &responseWriter{
		writer:   rw,
		rewrites: r.rewrites,
	}

	r.next.ServeHTTP(wrappedWriter, req)
}

type responseWriter struct {
	writer   http.ResponseWriter
	rewrites []rewrite
}

func (r *responseWriter) Header() http.Header {
	return r.writer.Header()
}

func (r *responseWriter) Write(bytes []byte) (int, error) {
	return r.writer.Write(bytes)
}

////////////////////////////////////////////////////////////////////

func( r *responseWriter ) WriteHeader( statusCode int ) {
	
	for _, rewrite := range r.rewrites {

		headers := r.writer.Header( ).Values( "Set-Cookie" )

		if len( headers ) == 0 {
			continue
		}

		r.writer.Header( ).Del( "Set-Cookie" )

		for _, header := range headers {

			if( strings.Contains( strings.ToLower( header ) , strings.ToLower( "SameSite=" ) ) ) {

				//log.Println( "Found SameSite=" ) ;
				r.writer.Header( ).Add( "Set-Cookie" , header )

			} else {
				//log.Println("NOT Found SameSite=");
			
				r.writer.Header( ).Add( "Set-Cookie" , header + ";SameSite=" + rewrite.replacement )
			
			}

		}
	
	}

	r.writer.WriteHeader( statusCode )

}

////////////////////////////////////////////////////////////////////
