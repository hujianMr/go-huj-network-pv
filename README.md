## go-huj-network-pv 客户端公网IP实时更新服务

### 背景
&nbsp;&nbsp;&nbsp;&nbsp; 由于我想在外网ssh访问个人的linux服务器, 服务器的端口在路由器映射的, 那么可以通过公网ip+端口访问，但是，公网ip是多变的，所以希望能实时的知道服务器的公网ip

### 项目结构

&nbsp;&nbsp;&nbsp;&nbsp; FetchPublicIP.go 文件是编译后部署客户端定时获取公网ip后推送到服务端        
&nbsp;&nbsp;&nbsp;&nbsp; FetchPublicIpServer.go 文件编译后部署到服务端，比如租的阿里云的服务器     
&nbsp;&nbsp;&nbsp;&nbsp; config.ini具体看配置描述     