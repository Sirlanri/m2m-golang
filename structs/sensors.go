package structs

//ResDataIn 传感器传入数据详情
type ResDataIn struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	Et  string `json:"et"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

//ResData 传感器传入数据大结构体
type ResData struct {
	M2mcin ResDataIn `json:"m2m:cin"`
}
