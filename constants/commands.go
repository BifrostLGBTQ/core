package constants

import "encoding/json"

type CommandEnvelope struct {
	Version string          `json:"version"` // "v1", "v2" gibi
	Code    string          `json:"code"`    // "chat.send_gif" gibi
	Payload json.RawMessage `json:"payload"` // tip bilinmiyor, sonra parse edilir
}

type TCommandTypes int

const (
	// AUTH
	CMD_AUTH_LOGIN  = "auth.login"
	CMD_AUTH_LOGOUT = "auth.logout"

	// CHAT
	CMD_CHAT_SEND_TEXT    = "chat.send_text"
	CMD_CHAT_SEND_GIF     = "chat.send_gif"
	CMD_CHAT_SEND_CALL    = "chat.send_call"
	CMD_CHAT_SEND_STICKER = "chat.send_sticker"

	// USER
	CMD_USER_UPDATE_PROFILE = "user.update_profile"
	CMD_USER_FETCH_PROFILE  = "user.fetch_profile"
)

const (
	INVALID_COMMAND TCommandTypes = iota
	ACT_AUTH_LOGIN
	ACT_AUTH_REGISTER
	ACT_AUTH_RESET_PASSWORD_REQUEST
	ACT_AUTH_RESET_PASSWORD

	ACT_USER_FETCH
	ACT_USER_TOGGLE_FOLLOW_STATUS
	ACT_USER_TOGGLE_BLOCK_STATUS
	ACT_USER_REPORT
	ACT_USER_GET_USER_BY_NAME
	ACT_USER_FETCH_USER_POSTS
	ACT_USER_FETCH_USER_COMMENTS
	ACT_USER_FETCH_USER_FOLLOWERS
	ACT_USER_FETCH_USER_FOLLOWINGS
	ACT_USER_FETCH_USER_GIFTS
	ACT_USER_UPDATE_USER_PROFILE
	ACT_USER_SUBSCRIBE_REQUEST

	ACT_CHAT_FETCH
	ACT_CHAT_SEND_MESSAGE
	ACT_TEST
)

/*
func main() {
	// Example usage
	command := ACT_ACT_LOGIN
	switch command {
	case ACT_ACT_PROMPT:
		// Handle prompt action
	case ACT_ACT_REGISTER:
		// Handle register action
	case ACT_ACT_LOGIN:
		// Handle login action
	case ACT_ACT_PROFILE:
		// Handle profile action
	case ACT_ACT_REQUEST:
		// Handle request action
	case ACT_ACT_CHECK_AUTH:
		// Handle check auth action
	default:
		// Handle unknown action
	}
}
*/
