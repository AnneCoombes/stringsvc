package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    var svc StringService
    svc = stringService{}
    svc = loggingMiddleware{logger, svc}

    // var uppercase endpoint.Endpoint

    // uppercase = makeUppercaseEndpoint(svc)
    // uppercase = loggingMiddleware(log.With(logger, "method", "uppercase")) (uppercase)

    // var count endpoint.Endpoint
    // count = makeCountEndpoint(svc)
    // count = loggingMiddleware(log.With(logger, "method", "count")) (count)

    uppercaseHandler := httptransport.NewServer(
        makeUppercaseEndpoint(svc),
        decodeUppercaseRequest,
        encodeResponse,
    )

    countHandler := httptransport.NewServer(
        makeCountEndpoint(svc), 
        decodeCountRequest, 
        encodeResponse, 
    )

    http.Handle("/uppercase", uppercaseHandler)
    http.Handle("/count", countHandler)
}

