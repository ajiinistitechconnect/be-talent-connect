package request

type ChangePassword struct {
	NewPassword     string
	CurrentPassword string
}
