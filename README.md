# idalloc 
常用 id 分配微服务，支持 golang + mysql ,golang + redis 

##配置文件 conf

通过修改配置，实现不同的存储
idalloc_info:
  save_mode : 1 # default 0 mysql  1 redis 2 file
  
  
