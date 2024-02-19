# Dispatch System of QianPBX

```
        _               ______ ______ __   __
       (_)              | ___ \| ___ \\ \ / /
  __ _  _   __ _  _ __  | |_/ /| |_/ / \ V / 
 / _` || | / _` || '_ \ |  __/ | ___ \ /   \ 
| (_| || || (_| || | | || |    | |_/ // /^\ \
 \__, ||_| \__,_||_| |_|\_|    \____/ \/   \/
    | |                                      
    |_|        
Your Best Source Asterisk PBX GUI Solution    
```

![go version](https://img.shields.io/badge/Go-v1.22-blue?logo=Go)


## Usage

> Learn to use `FreePBX` first

先确认 asterisk系统 使用的同名数据库 `asterisk` 里有至少 27 张表（包括 devices）, 如果没有, 可以利用项目目录 `utility` 里的表文件创建

```bash
mysql -D asterisk < tables.sql
```

可执行文件需要和配置文件放在一起

```bash
# tree -C -L 2 .
.
├── config
│   └── config.yaml
├── dispatchAst
└── ...
```


## Reference

[1] [see the LAN IP of the phone](https://community.freepbx.org/t/how-to-sip-show-peers-vs-pjsip-show-endpoints-local-lan-ip/78847/10)

[2] [how to create entension in freepbx without gui](https://community.freepbx.org/t/get-list-of-extension-or-create-extension-through-rest-api/42886/10)