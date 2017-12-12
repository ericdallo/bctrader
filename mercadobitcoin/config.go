package mercadobitcoin

import (
	"github.com/spf13/viper"

	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func BaseURL() string {
	return viper.GetString("mercado_bitcoin_base_url")
}

func TapiId() string {
	return viper.GetString("tapi_id")
}

func TapiMacFor(path string) string {
	secret := viper.GetString("tapi_secret")

	hmac512 := hmac.New(sha512.New, []byte(secret))
	hmac512.Write([]byte(path))

	return hex.EncodeToString(hmac512.Sum(nil))
}

func IsCredentialsSetted() bool {
	if viper.GetString("tapi_id") == "" {
		return false
	}

	if viper.GetString("tapi_secret") == "" {
		return false
	}

	return true
}