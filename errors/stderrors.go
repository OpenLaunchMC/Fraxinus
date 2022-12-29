package errors

var InvalidToken = YggdrasilErrorResponse{
	StatusCode: 403,
	Error: "ForbiddenOperationException",
	ErrorMessage: "Invalid token.",
	Cause: "Token is Invalid",
}

var InvalidCredentials = YggdrasilErrorResponse{
	StatusCode: 403,
	Error: "ForbiddenOperationException",
	ErrorMessage: "Invalid credentials. Invalid username or password.",
	Cause: "Credentials are Invalid",
}

var InvalidJSONFormat = YggdrasilErrorResponse{
	StatusCode: 400,
	Error: "JSONMappingException",
	// You should replace this with your own JSON Error Massage
	ErrorMessage: "Unexpected character",
}

var UserMigrated = YggdrasilErrorResponse{
	StatusCode: 400,
	Error: "ForbiddenOperationException",
	ErrorMessage: "Account migrated, use email as username.",
	Cause: "UserMigratedException",
}

var InvalidRequestBody = YggdrasilErrorResponse{
	StatusCode: 403,
	Error: "ForbiddenOperationException",
	ErrorMessage: "Forbidden",
}

var InvalidRequestMethod = YggdrasilErrorResponse{
	StatusCode: 405,
	Error: "Method Not Allowed",
	ErrorMessage: "Method Not Allowed",
}

var UserGone = YggdrasilErrorResponse{
	StatusCode: 410,
	Error: "Gone",
	ErrorMessage: "Migrated",
}

var InvalidMediaType = YggdrasilErrorResponse{
	StatusCode: 415,
	Error: "Unsupported Media Type",
	ErrorMessage: "The server is refusing to service the request because the entity of the request is in a format not supported by the requested resource for the requested method.",
}

var TokenAlreadyAssigned = YggdrasilErrorResponse{
	StatusCode: 400,
	Error: "IllegalArgumentException",
	ErrorMessage: "Access token had already been assigned.",
}

var InvalidAssignment = YggdrasilErrorResponse{
	StatusCode: 403,
	Error: "ForbiddenOperationException",
	ErrorMessage: "The profile requested to be assigned is not owned by this account",
}

var InvalidProfile = YggdrasilErrorResponse{
	StatusCode: 403,
	Error: "ForbiddenOperationException",
	ErrorMessage: "Invalid Token.",
}

var Success = YggdrasilErrorResponse{
	StatusCode: 200,
}

var LimitReached = YggdrasilErrorResponse{
	StatusCode: 429,
	Error: "TooManyRequestsException",
	ErrorMessage: "Too Many Requests",
}