package types

import "bifrost/models/post"

type TimelineResult struct {
	Posts      []post.Post `json:"posts"`       // Döndürülen postlar
	NextCursor *int64      `json:"next_cursor"` // Bir sonraki sayfa için cursor (PublicID)
}
