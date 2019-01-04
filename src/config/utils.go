/**
* @auth    kunlun
* @date    2019-01-04 11:27
* @version v1.0
* @des     描述：工具类
*
**/
package config

/**
* clear slice, input value address change
* params: s pointer   type *[]interface{}
*
**/
func ClearSlice(s *[]interface{}) {
	*s = append([]interface{}{})
}

/**
* clear slice, input value address not change
* params: s pointer  type *[]interface{}
*
 */
func ClearSlice1(s *[]interface{}) {
	*s = (*s)[0:0]
}
