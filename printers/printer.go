package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintNewUID() string {
	id := uuid.New()
	return fmt.Sprintf("%s\n", id)
}
