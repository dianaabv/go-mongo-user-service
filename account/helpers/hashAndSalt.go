
package helpers
import(
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
    
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
		return "Hash Problem"
		// fmt.Println("Inserted a Single Document: ", err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}