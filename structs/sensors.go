package structs

//ReqDataIn 传感器传入数据详情
type ReqDataIn struct {
	Rn  string `json:"rn"`  //资源名称
	Ty  int    `json:"ty"`  //类型
	Ri  string `json:"ri"`  //资源ID与资源名称对应
	Pi  string `json:"pi"`  //父ID
	Ct  string `json:"ct"`  //(创建时间）2020年11月21日 10时01分 33秒创建，下同
	Lt  string `json:"lt"`  //最近更新时间
	Et  string `json:"et"`  //过期时间
	St  int    `json:"st"`  //状态标签
	Cnf string `json:"cnf"` //内容格式
	Cs  int    `json:"cs"`  //内容大小）con的长度，con是386的时候cs就是3
	Con string `json:"con"` //内容正文！
}

//SensorData 发送给传感器Wifi的数据
type SensorData struct {
	M2m ReqDataIn `json:"m2m:cin"`
}

//ReqData 传感器传入数据大结构体
type ReqData struct {
	M2mcin ReqDataIn `json:"m2m:cin"`
}

//Response 返回给传感器的数据
type Response struct {
	M2mcin ReqDataIn `json:"m2m:cin"`
}

//RequestInt 传感器传入子数据 int
type RequestInt struct {
	Con int `json:"con"`
}

//RequestFloat 传感器传入子数据 float
type RequestFloat struct {
	Con float32 `json:"con"`
}

//RequestSting 传感器传入子数据 float
type RequestSting struct {
	Con string `json:"con"`
}

//RequestData2 传感器传入的完整数据
type RequestData2 struct {
	M2m RequestSting `json:"m2m:cin"`
}

//WifiPost 传给wifi模块的子数据
type WifiPost struct {
	Con string `json:"con"`
}

//WifiPostData 传给wifi模块的数据
type WifiPostData struct {
	M2m WifiPost `json:"m2m:cin"`
}
