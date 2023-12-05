/**
 * @Author: kwens
 * @Date: 2022-08-10 16:39:33
 * @Description:
 */
package oprlog

import "github.com/google/uuid"

func genUuid() string {
	u4 := uuid.New()
	return u4.String()
}
