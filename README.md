# gocomposer
之前使用 Toran 作为 composer 的缓存服务,但是它太慢了,而且需要 PHP 环境, 管理不方便.
所以用 Golang 简单写了写,不是完整的代理服务器,只是处理了部分缓存的逻辑,目前 gocomposer 作为缓存服务器,基本已经够用了.

### 配置
使用 .env 作为配置文件
```
# 缓存服务器运行的域名: http://pkg.mydomain.com:8080 末尾不要添加 /
domain="http://pkg.mydomain.com:8080"

# 上游镜像,可以使用packagist或其他国内镜像
mirror="https://repo.packagist.org"

# 生产环境:production 测试:development, 减少日志输出
runmode="production"
```

### 运行
```
./gocomposer
```

### 使用
```
composer config repo.packagist composer http://pkg.mydomain.com:8080
composer install
```

### Dashboard
简单列出所有已经缓存的package
```
http://pkg.mydomain.com:8080/dashboard
```

## 说明
 - 只支持 composer 类型的 repository, 不支持 vcs 类型
 - 如果 pkg.mydomain.com 没有 SSL 证书,需要配置 config 中 "secure-http": false
 - 会替换 package.json 中的 zipball 地址为缓存地址
 - 在 composer 2.0.14 下测试通过
 - 所有的缓存放在 cache 目录下
 - 没有任何身份验证,自行注意安全