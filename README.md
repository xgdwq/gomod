# gomod
test go mod 
> 
参考： 
> 
https://www.jianshu.com/p/98feb7c36e1i8 
> 
1,  go mod 一个模块多版本共存 
> 
1）不同版本创建不同分支，并修改go.mod中对应模块名字，如 moduleName/v2, moduleName/v3 等 
> 
2）分支push改动后，创建tag，并将tag push 到master 
> 
git tag -m "v2.0.0” v2.0.0 
> 
git push origin master --tags 
> 
3）调用地方除了import导入，需用require指令修改go.mod文件 
> 
Go文件： 
> 
import gomodv5 "github.com/xgdwq/gomod/v5" 
> 
go.mod: 
> 
require github.com/xgdwq/gomod/v5 v5.0.0 
