/**
 * @Author: kwens
 * @Date: 2022-08-10 15:57:26
 * @Description:
 */
package stringid

import (
	"math/rand"

	"github.com/google/uuid"
)

// GenUuid 生成Uuid
func GenUuid() string {
	u4 := uuid.New()
	return u4.String()
}

// GenID 生成随机N位ID
func GenID(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(90-65))
	}
	return string(bytes)
}
