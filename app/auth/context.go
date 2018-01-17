package auth

import "context"

// UserContext extends auth.Claims with IsAdmin()
// helper and does not transfer jwt issue level
// properties
type UserContext Claims

// IsAdmin returns if the user role is admin
func (uc *UserContext) IsAdmin() bool {
	return uc.Role == "admin"
}

// GetUserContextFrom extracts user claims from a context
// and returns a user context
func GetUserContextFrom(ctx context.Context) *UserContext {
	id := uint(ctx.Value("id").(int))
	role, _ := ctx.Value("role").(string)
	username, _ := ctx.Value("username").(string)

	return &UserContext{ID: id, Role: role, Username: username}
}
