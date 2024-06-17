package middleware

import (
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func verifyJWTMiddleware(token string) (map[string]interface{}, error) {
	token_payload, err := utils.VerifyJWT(token)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	if value, ok := token_payload["exp_date"]; ok {
		currentUnixTime := time.Now().Unix()
		expDateFloat, ok_int := value.(float64)
		if !ok_int {
			return nil, fmt.Errorf("unauthorized")
		}
		expDateUnix := int64(expDateFloat)
		if currentUnixTime > expDateUnix {
			return nil, fmt.Errorf("token_expired")
		}
	} else {
		return nil, fmt.Errorf("unauthorized")
	}
	return token_payload, nil
}
