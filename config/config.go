package config

import(
	"os"
	"fmt"
)

var Secret = fmt.Sprintf( " %s ", os.Getenv("SECRET_KEY"))