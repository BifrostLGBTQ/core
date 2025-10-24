package constants

import "encoding/json"

type CommandEnvelope struct {
	Version string          `json:"version"` // "v1", "v2" gibi
	Code    string          `json:"code"`    // "chat.send_gif" gibi
	Payload json.RawMessage `json:"payload"` // tip bilinmiyor, sonra parse edilir
}

type TCommandTypes int

const (
	//SYSTEM
	CMD_INITIAL_SYNC = "system.initial_sync"
	// AUTH
	CMD_AUTH_LOGIN    = "auth.login"
	CMD_AUTH_REGISTER = "auth.register"
	CMD_AUTH_LOGOUT   = "auth.logout"
	CMD_AUTH_TEST     = "auth.test"

	// CHAT
	CMD_CHAT_SEND_TEXT    = "chat.send_text"
	CMD_CHAT_SEND_GIF     = "chat.send_gif"
	CMD_CHAT_SEND_CALL    = "chat.send_call"
	CMD_CHAT_SEND_STICKER = "chat.send_sticker"

	// USER
	CMD_USER_UPDATE_PROFILE = "user.update_profile"
	CMD_USER_FETCH_PROFILE  = "user.fetch_profile"
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
