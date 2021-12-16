package events

import "github.com/gompus/gompus/models/user"

// UserUpdateEvent is the event sent when
// properties about the user change.
type UserUpdateEvent user.User
