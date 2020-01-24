package helpers
import (
	"gokit-example/account/models"
)
// type sd models.User

func CheckValues(user models.User) bool {
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Lastname == "" || user.Phone == "" || user.Dob == ""  || user.Country == ""  || user.Bio == "" {
		return false
	}
	return true
}