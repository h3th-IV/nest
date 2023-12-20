## Learing Net and the Good stuff

- [x] httpclient
        a software application enables communication with web servers ussing HTTP

    - Request; An HTTP client sends a request to a server to perform a specific 
        action or retrieve information. The request includes details such as the HTTP 
        method (e.g., GET, POST), headers, and, in the case of a POST request, a payload
        
    - Response: Response: The server processes the request and sends back an HTTP 
        response. The response includes a status code indicating the outcome (e.g: 200 OK, 
        404 Not Found) and may also contain data in the response body.

    - Methods: clients use specific HTTP methods toperform various actions
            * GET- retrieve data from web server
            * POST- submit data to servr fro processing
            * PUT-  update resource or create resource ifit doesn't exist
            * DELETE- request removal of a resource

    - Headers: provide additional info abut request or response

        -- Response Status Code: indicates the outcome of request


- [X] http Server
    a sotware components that listensfor incoming http request and provide 
    response to those request

    - Routing: mappping incoming request to appropriat handlers
        based on the request url.

    - Handler:(a handler is an oobject implementing the htt.Handler interface),  a function 
        that is called when a request with a particular pattern is recieved by the server.
    
    - Pattern in Go Handler Func: a pattern is a string tthatth defines the url pattern or
        route. Patterns areused to to match incoming HTTP Request and baesd on the match,
        the corresponding handler function is called.
    
    - ServeMUx(Serve MUltiplexer): is an object in GO's http package that rooute incoming 
        incoming http request to the appropriate handler baesd on the request url