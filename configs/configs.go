/*Package configs 配置文件
用于区分本地开发和服务器部署，顺便可以保密
*/
package configs

//是否为本地开发环境
var dev = true

//SQLConfg 数据库信息
func SQLConfg() string {
	var local = "root:123456@/mym2m"
	var serve = ""
	if dev {
		return local
	}
	return serve
}
