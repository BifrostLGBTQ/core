package constants

type ErrorCode int

const (
	UnknownError ErrorCode = iota
	FileNotFound
	PermissionDenied
	InvalidInput
	NetworkError
	DatabaseError
	ResourceNotAvailable
	InvalidAction
	InvalidPassword
)

var ErrorMessages = map[ErrorCode]string{
	UnknownError:         "Unknown error occurred.",
	FileNotFound:         "The specified file could not be found.",
	PermissionDenied:     "Permission denied for the requested action.",
	InvalidInput:         "Invalid input provided.",
	NetworkError:         "A network error occurred.",
	DatabaseError:        "An error occurred while accessing the database.",
	ResourceNotAvailable: "The requested resource is not available.",
	InvalidAction:        "The action is not valid.",
	InvalidPassword:      "Invalid Password",
}

func (e ErrorCode) String() string {
	return ErrorMessages[e]
}

/*
func main() {
    errorCode := FileNotFound
    errorMessage := errorCode.String()
    println("Error Code:", errorCode, "- Error Message:", errorMessage)
}
*/
