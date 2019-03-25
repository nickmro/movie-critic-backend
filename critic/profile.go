package critic

import "time"

// Profile represents a user's profile information.
type Profile struct {
	ID          int64      // The profile ID
	UserID      int64      // The user ID
	FirstName   string     // The first name
	LastName    string     // The last name
	Description string     // The description
	CreatedAt   time.Time  // The time of creation
	UpdatedAt   *time.Time // The time of last update
	DeletedAt   *time.Time // The time of deletion
}
