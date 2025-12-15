package constants

import "errors"

var (
	ErrChatRetrieval = errors.New("error retrieving chats")
	ErrChatInput     = errors.New("error inserting chat")
	ErrBooking       = errors.New("error inserting booking schedue")
)
