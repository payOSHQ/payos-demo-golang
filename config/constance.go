package config

import (
	"os"

	"github.com/joho/godotenv"
)

const YOUR_DOMAIN = "http://localhost:8080"
const PAYOS_BASE_URL = "https://api-merchant.payos.vn"

var PAYOS_CLIENT_ID string
var PAYOS_API_KEY string
var PAYOS_CHECKSUM_KEY string

func init() {
	godotenv.Load(".env")

	PAYOS_CLIENT_ID = os.Getenv("PAYOS_CLIENT_ID")
	PAYOS_API_KEY = os.Getenv("PAYOS_API_KEY")
	PAYOS_CHECKSUM_KEY = os.Getenv("PAYOS_CHECKSUM_KEY")

}
