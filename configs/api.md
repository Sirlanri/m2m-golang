# 数据格式

{
   "m2m:cin" : {
      "rn" : "cin_7099962211655080384",              (资源名称）
      "ty" : 4,                                                              （类型）
      "ri" : "/server/cin-7099962211655080384", （资源ID)      资源ID与资源名称对应
      "pi" : "/server/cnt-8057999773803139777",  (父ID)
      "ct" : "20201121T100133",                                (创建时间）2020年11月21日 10时01分 33秒创建，下同
      "lt" : "20201121T100133",                              （最近更新时间）
      "et" : "20201121T100133",                             （过期时间）
      "st" : 0,                                                               （状态标签）
      "cnf" : "text/plain:0",                                       （内容格式）
      "cs" : 4,                                                              （内容大小）con的长度，con是386的时候cs就是3
      "con" : "1024"                                                  （内容）
   }
}

# 完整的post过程

connecting to 192.168.43.87
POST /server/mydevice1/luminosity HTTP/1.1
Host: 192.168.43.87
X-M2M-Origin: Cae_device1
Content-Type: application/json;ty=4
Content-Length: 25
Connection: close

{"m2m:cin":{"con":"688"}}

HTTP/1.1 201 Created
Connection: close
Date: Sat, 21 Nov 2020 04:49:26 GMT
Set-Cookie: JSESSIONID=node0caxl6ue5iu1l19rypa5jyf5uh142.node0;Path=/
Expires: Thu, 01 Jan 1970 00:00:00 GMT
Access-Control-Allow-Origin: *
Access-Control-Allow-Headers: X-M2M-Origin,Content-Type,X-M2M-Key
Access-Control-Allow-Methods: DELETE, PUT, GET, POST
Content-Location: /server/cin-3369042773441023094
X-M2M-Origin: /server
X-M2M-RSC: 2001
Content-Type: application/json
Content-Length: 369
Server: Jetty(9.4.7.v20170914)

{
   "m2m:cin" : {
      "rn" : "cin_3369042773441023094",
      "ty" : 4,
      "ri" : "/server/cin-3369042773441023094",
      "pi" : "/server/cnt-3286409875966090923",
      "ct" : "20201121T124926",
      "lt" : "20201121T124926",
      "et" : "20201121T124926",
      "st" : 0,
      "cnf" : "text/plain:0",
      "cs" : 3,
      "con" : "688"
   }
}
closing connection



# 完整的response过程

▶▶▶▶▶
POST http://127.0.0.1:8080/server/mydevice1/led
{ 'm2m:cin': { con: 'ON' } }
◀◀◀◀◀
201
{ 'm2m:cin':
   { rn: 'cin_1567083670256467599',
     ty: 4,
     ri: '/server/cin-1567083670256467599',
     pi: '/server/cnt-2796593338684513035',
     ct: '20201121T125316',
     lt: '20201121T125316',
     et: '20201121T125316',
     st: 0,
     cnf: 'text/plain:0',
     cs: 2,
     con: 'ON'

 } 

}

