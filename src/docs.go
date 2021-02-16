// Telegram CORE API interlayer
//
// Documentation.
//
//     Schemes: http
//     BasePath: /
//     Version: 0.0.1
//     Host: localhost:5000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package main

// swagger:route GET /articles/search/:query?apiKey=;page=;pageSize=;
// Returns found articles with pagination.
// responses:
//   200: Article[]

// swagger:response GetArticlesResponse
type GetArticlesWrapper struct {
	// in:body
	Body []Article
}

// swagger:parameters GetArticlesEndpoint
type foobarParamsWrapper struct {
	// in:body
	Body IGetArticles
}
