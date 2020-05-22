# gt-gin-vue

POST   /auth/login      登录         curl http://host:port/auth/login -d '{"name":"gt","pwd":"123"}'
POST   /auth/register   注册         curl http://host:port/auth/register -d '{"username":"gt","pwd":"123"}'

GET    /user/test       测试token    curl http://host:port/user/test -H "token:abc"
GET    /user/resetpwd   重置用户密码  curl http://host:port/user/resetpwd -H "token:abc"(此处token内会存储用户id)
POST   /user/delete     删除用户		 curl http://host:port/user/delete -H "token:abc" -d "name=gt"
POST   /user/update     更新用户信息  curl http://host:port/user/test -H "token:abc" -d "name=gt&pwd=123" 


