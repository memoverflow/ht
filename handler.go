package ht

// Handler is the type for middlewares and route handlers to register.
type Handler func(*Request, *Response)
