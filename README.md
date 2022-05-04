# MacOS 顶栏的 GitHub Star 数统计

## 打包

修改 config/.env.json 文件

```json
{
  "repo": "hyperf/hyperf",
  "token": "ghp_xxx"
}
```

执行以下命令进行打包

```shell
go build -o star-bar main.go
```

编写对应的 plist

> ProgramArguments 换成对应的脚本地址

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>KeepAlive</key>
        <dict>
            <key>SuccessfulExit</key>
            <false/>
        </dict>
        <key>Label</key>
        <string>cn.limingxinleo.star-bar</string>
        <key>ProgramArguments</key>
        <array>
            <string>/Users/limingxin/.bin/star-bar</string>
        </array>
        <key>RunAtLoad</key>
        <true/>
        <key>WorkingDirectory</key>
        <string>/tmp</string>
    </dict>
</plist>
```