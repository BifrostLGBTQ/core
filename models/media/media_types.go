package media

type MediaRole string
type OwnerType string

const (
	// Role
	RoleProfile   MediaRole = "profile"
	RoleCover     MediaRole = "cover"
	RolePost      MediaRole = "post"
	RoleBlog      MediaRole = "blog"
	RoleChatImage MediaRole = "chat_image"
	RoleChatVideo MediaRole = "chat_video"
	RoleOther     MediaRole = "other"

	// Owner Type
	OwnerUser OwnerType = "user"
	OwnerPost OwnerType = "post"
	OwnerBlog OwnerType = "blog"
	OwnerChat OwnerType = "chat"
	OwnerPage OwnerType = "page"
)
