package api

import "fmt"

type response struct {
	Code        int
	TypeOf      string
	Description string
}

var UnknownResponse = response{
	Code:        0,
	TypeOf:      "Unknown Response",
	Description: "The response %v, is not yet recognised",
}

var Response100 = response{
	Code:        100,
	TypeOf:      "Continue",
	Description: "An interim response. Indicates to the client that the initial part of the request has been received and has not yet been rejected by the server. The client SHOULD continue by sending the remainder of the request or, if the request has already been completed, ignore this response. The server MUST send a final response after the request has been completed.",
}

var Response101 = response{
	Code:        101,
	TypeOf:      "Switching Protocol",
	Description: "Sent in response to an Upgrade request header from the client, and indicates the protocol the server is switching to.",
}

var Response102 = response{
	Code:        102,
	TypeOf:      "Processing (WebDAV)",
	Description: "Indicates that the server has received and is processing the request, but no response is available yet.",
}

var Response103 = response{
	Code:        103,
	TypeOf:      "Early Hints",
	Description: "Primarily intended to be used with the Link header. It suggests the user agent start preloading the resources while the server prepares a final response.",
}

var Response200 = response{
	Code:        200,
	TypeOf:      "OK",
	Description: "Indicates that the request has succeeded.",
}

var Response201 = response{
	Code:        201,
	TypeOf:      "Created",
	Description: "Indicates that the request has succeeded and a new resource has been created as a result.\n",
}

var Response202 = response{
	Code:        202,
	TypeOf:      "Accepted",
	Description: "Indicates that the request has been received but not completed yet. It is typically used in log running requests and batch processing.",
}

var Response203 = response{
	Code:        203,
	TypeOf:      "Non-Authoritative Information",
	Description: "Indicates that the returned metainformation in the entity-header is not the definitive set as available from the origin server, but is gathered from a local or a third-party copy. The set presented MAY be a subset or superset of the original version.",
}

var Response204 = response{
	Code:        204,
	TypeOf:      "No Content",
	Description: "The server has fulfilled the request but does not need to return a response body. The server may return the updated meta information.\n",
}

var Response205 = response{
	Code:        205,
	TypeOf:      "Reset Content",
	Description: "Indicates the client to reset the document which sent this request.",
}

var Response206 = response{
	Code:        206,
	TypeOf:      "Partial Content",
	Description: "It is used when the Range header is sent from the client to request only part of a resource.",
}

var Response207 = response{
	Code:        207,
	TypeOf:      "Multi-Status (WebDAV)",
	Description: "An indicator to a client that multiple operations happened, and that the status for each operation can be found in the body of the response.",
}

var Response208 = response{
	Code:        208,
	TypeOf:      "Already Reported (WebDAV)",
	Description: "Allows a client to tell the server that the same resource (with the same binding) was mentioned earlier. It never appears as a true HTTP response code in the status line, and only appears in bodies.",
}

var Response226 = response{
	Code:        226,
	TypeOf:      "IM Used",
	Description: "The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.\n",
}

var Response300 = response{
	Code:        300,
	TypeOf:      "Multiple Choices",
	Description: "The request has more than one possible response. The user-agent or user should choose one of them.",
}

var Response301 = response{
	Code:        301,
	TypeOf:      "Moved Permanently",
	Description: "The URL of the requested resource has been changed permanently. The new URL is given by the Location header field in the response. This response is cacheable unless indicated otherwise.",
}

var Response302 = response{
	Code:        302,
	TypeOf:      "Found",
	Description: "The URL of the requested resource has been changed temporarily. The new URL is given by the Location field in the response. This response is only cacheable if indicated by a Cache-Control or Expires header field.",
}

var Response303 = response{
	Code:        303,
	TypeOf:      "See Other",
	Description: "The response can be found under a different URI and SHOULD be retrieved using a GET method on that resource.",
}

var Response304 = response{
	Code:        304,
	TypeOf:      "Not Modified",
	Description: "Indicates the client that the response has not been modified, so the client can continue to use the same cached version of the response.",
}

var Response305 = response{
	Code:        305,
	TypeOf:      "Use Proxy",
	Description: "Indicates that a requested response must be accessed by a proxy.",
}

var Response306 = response{
	Code:        306,
	TypeOf:      "(Unused)",
	Description: "It is a reserved status code and is not used anymore.",
}

var Response307 = response{
	Code:        307,
	TypeOf:      "Temporary Redirect",
	Description: "Indicates the client to get the requested resource at another URI with same method that was used in the prior request. It is similar to 302 Found with one exception that the same HTTP method will be used that was used in the prior request.",
}

var Response308 = response{
	Code:        308,
	TypeOf:      "Permanent Redirect",
	Description: "Indicates that the resource is now permanently located at another URI, specified by the Location header. It is similar to 301 Moved Permanently with one exception that the same HTTP method will be used that was used in the prior request.",
}

var Response400 = response{
	Code:        400,
	TypeOf:      "Bad Request",
	Description: "The request could not be understood by the server due to incorrect syntax. The client SHOULD NOT repeat the request without modifications.",
}

var Response401 = response{
	Code:        401,
	TypeOf:      "Unauthorised",
	Description: "Indicates that the request requires user authentication information. The client MAY repeat the request with a suitable Authorization header field",
}

var Response402 = response{
	Code:        402,
	TypeOf:      "Payment Required (Experimental)",
	Description: "Reserved for future use. It is aimed for using in the digital payment systems.",
}

var Response403 = response{
	Code:        403,
	TypeOf:      "Forbidden",
	Description: "Unauthorized request. The client does not have access rights to the content. Unlike 401, the client’s identity is known to the server.",
}

var Response404 = response{
	Code:        404,
	TypeOf:      "Not Found",
	Description: "The server can not find the requested resource.",
}

var Response405 = response{
	Code:        405,
	TypeOf:      "Method Not Allowed",
	Description: "The request HTTP method is known by the server but has been disabled and cannot be used for that resource.",
}

var Response406 = response{
	Code:        406,
	TypeOf:      "Not Acceptable",
	Description: "The server does not find any content that conforms to the criteria given by the user agent in the Accept header sent in the request.",
}

var Response407 = response{
	Code:        407,
	TypeOf:      "Proxy Authentication Required",
	Description: "Indicates that the client must first authenticate itself with the proxy.",
}

var Response408 = response{
	Code:        408,
	TypeOf:      "Request Timeout",
	Description: "Indicates that the server did not receive a complete request from the client within the server’s allotted timeout period.",
}

var Response409 = response{
	Code:        409,
	TypeOf:      "Conflict",
	Description: "The request could not be completed due to a conflict with the current state of the resource.",
}

var Response410 = response{
	Code:        410,
	TypeOf:      "Gone",
	Description: "The requested resource is no longer available at the server.",
}

var Response411 = response{
	Code:        411,
	TypeOf:      "Length Required",
	Description: "The server refuses to accept the request without a defined Content- Length. The client MAY repeat the request if it adds a valid Content-Length header field.",
}

var Response412 = response{
	Code:        412,
	TypeOf:      "Precondition Failed",
	Description: "The client has indicated preconditions in its headers which the server does not meet.\n",
}

var Response413 = response{
	Code:        413,
	TypeOf:      "Request Entity Too Large",
	Description: "Request entity is larger than limits defined by server.",
}

var Response414 = response{
	Code:        414,
	TypeOf:      "Request-URI Too Long",
	Description: "The URI requested by the client is longer than the server can interpret.",
}

var Response415 = response{
	Code:        415,
	TypeOf:      "Unsupported Media Type",
	Description: "The media-type in Content-type of the request is not supported by the server.",
}

var Response416 = response{
	Code:        416,
	TypeOf:      "Requested Range Not Satisfiable",
	Description: "The range specified by the Range header field in the request can’t be fulfilled.",
}

var Response417 = response{
	Code:        417,
	TypeOf:      "Expectation Failed",
	Description: "The expectation indicated by the Expect request header field can’t be met by the server.",
}

var Response418 = response{
	Code:        418,
	TypeOf:      "I’m a teapot (RFC 2324)",
	Description: "It was defined as April’s fool joke and is not expected to be implemented by actual HTTP servers.",
}

var Response420 = response{
	Code:        420,
	TypeOf:      "Enhance Your Calm (Twitter)",
	Description: "Returned by the Twitter Search and Trends API when the client is being rate limited.",
}

var Response422 = response{
	Code:        422,
	TypeOf:      "Unprocessable Entity (WebDAV)",
	Description: "The server understands the content type and syntax of the request entity, but still server is unable to process the request for some reason.",
}

var Response423 = response{
	Code:        423,
	TypeOf:      "Locked (WebDAV)",
	Description: "The resource that is being accessed is locked.",
}

var Response424 = response{
	Code:        424,
	TypeOf:      "Failed Dependency (WebDAV)",
	Description: "The request failed due to failure of a previous request.",
}

var Response425 = response{
	Code:        425,
	TypeOf:      "Too Early (WebDAV)",
	Description: "Indicates that the server is unwilling to risk processing a request that might be replayed.",
}

var Response426 = response{
	Code:        426,
	TypeOf:      "Upgrade Required",
	Description: "The server refuses to perform the request. The server will process the request after the client upgrades to a different protocol.",
}

var Response428 = response{
	Code:        428,
	TypeOf:      "Precondition Required",
	Description: "The origin server requires the request to be conditional.",
}

var Response429 = response{
	Code:        429,
	TypeOf:      "Too Many Requests",
	Description: "The user has sent too many requests in a given amount of time (“rate limiting”).",
}

var Response431 = response{
	Code:        431,
	TypeOf:      "Request Header Fields Too Large",
	Description: "The server is unwilling to process the request because its header fields are too large.",
}

var Response444 = response{
	Code:        444,
	TypeOf:      "No Response (Nginx)",
	Description: "The Nginx server returns no information to the client and closes the connection.",
}

var Response449 = response{
	Code:        449,
	TypeOf:      "Retry With (Microsoft)",
	Description: "The request should be retried after performing the appropriate action.",
}

var Response450 = response{
	Code:        450,
	TypeOf:      "Blocked by Windows Parental Controls (Microsoft)",
	Description: "Windows Parental Controls are turned on and are blocking access to the given webpage.",
}

var Response451 = response{
	Code:        451,
	TypeOf:      "Unavailable For Legal Reasons",
	Description: "The user-agent requested a resource that cannot legally be provided.\n",
}

var Response499 = response{
	Code:        499,
	TypeOf:      "Client Closed Request (Nginx)",
	Description: "The connection is closed by the client while HTTP server is processing its request, making the server unable to send the HTTP header back.",
}

var Response500 = response{
	Code:        500,
	TypeOf:      "Internal Server Error",
	Description: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
}

var Response501 = response{
	Code:        501,
	TypeOf:      "Not Implemented",
	Description: "The HTTP method is not supported by the server and cannot be handled.",
}

var Response502 = response{
	Code:        502,
	TypeOf:      "Bad Gateway",
	Description: "The server got an invalid response while working as a gateway to get the response needed to handle the request.",
}

var Response503 = response{
	Code:        503,
	TypeOf:      "Service Unavailable",
	Description: "The server is not ready to handle the request.",
}

var Response504 = response{
	Code:        504,
	TypeOf:      "Gateway Timeout",
	Description: "The server is acting as a gateway and cannot get a response in time for a request.",
}

var Response505 = response{
	Code:        505,
	TypeOf:      "HTTP Version Not Supported (Experimental)",
	Description: "The HTTP version used in the request is not supported by the server.",
}

var Response506 = response{
	Code:        506,
	TypeOf:      "Variant Also Negotiates (Experimental)",
	Description: "Indicates that the server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper endpoint in the negotiation process.",
}

var Response507 = response{
	Code:        507,
	TypeOf:      "Insufficient Storage (WebDAV)",
	Description: "The method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request.\n",
}

var Response508 = response{
	Code:        508,
	TypeOf:      "Loop Detected (WebDAV)",
	Description: "The server detected an infinite loop while processing the request.\n",
}

var Response510 = response{
	Code:        510,
	TypeOf:      "Not Extended",
	Description: "Further extensions to the request are required for the server to fulfill it.",
}

var Response511 = response{
	Code:        511,
	TypeOf:      "Network Authentication Required",
	Description: "Indicates that the client needs to authenticate to gain network access.\n",
}

func (r *response) fmtUnknownResponse(code int) {
	if r != &UnknownResponse {
		return
	}
	r.Description = fmt.Sprintf(UnknownResponse.Description, code)
}

func (r response) IsOK() bool {
	return r == Response200
}
