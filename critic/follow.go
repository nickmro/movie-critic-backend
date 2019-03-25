package critic

import "time"

// Follow represents a user following another user.
type Follow struct {
	ID         int64      // The follow ID
	FollowerID int64      // The following user ID
	FollowedID int64      // The followed user ID
	CreatedAt  time.Time  // The time of creation
	UpdatedAt  *time.Time // The time of last update
	DeletedAt  *time.Time // The time of deletion
}

// FollowRepository defines operations that may be used to store a follow.
type FollowRepository interface {
	Upsert(follow *Follow) error
	Delete(id int64) error
}
