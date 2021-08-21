### RiRo 
- 资源索引系统
- 通过读取索引文件 直接返回值前端
### repository 需求说明
- repository 下分三种文件
    - 音视频文件 
    - 指向性文件
    - 目录文件
- 音视频根据文件后缀区分
- 指向性文件 repository.direct.toml
    - 定义： 指向性文件为元文件，为 非当前repository MiMo 目录下的文件 路径文件
    - 格式:
    ```
        文件名1.text -> http://127.0.0.1/文件名1.text
            
    ```
- 知识库文件 .repository.index.knowledge.toml
    - 定义: 描述knowledge目录下的知识库信息，根据该文件的顺序和文案直接显示到 web 端
    - 格式:
    ```
    [meta]
        total: 100
        size:  20Mb
    [knowledge]
        [knowledge.1]
        name = 'MiMo开发手册'
        cover: /static/img.png
        creation_time: '20210808121212'
        store = direct:0
        [knowledge.2]
        name = '知识库1'
        cover: /static/img.png
        creation_time: '20210808111111'
        store = direct:0
  ```
  - 解释
    - store 
        - direct:0 
- 知识库下文件索引 .repository.index.file.toml
    - 定义: 描述所在知识库下的文件的顺序和文案，且直接显示到 web 端
    - 格式:
    ```
    [meta]
        total: 100
        title: 
        author: [uid.nikename, uid.nikename]
        last_revision: datetime
        creation_time: datetime
    [resource]
        [resource.1]
        name = 'test'
        creation_time: '20210808121212'
        store = 'direct:0'
        type = '0'
        upload = '0'
        [resource.1]
        name = '知识库1'
        creation_time: '20210808111111'
        [resource.2]
        store = 'remote:2'
  ```
  - 解释
    - store 
        - direct:0 
            - repository.cfg.toml 中第0个库 repository 默认在MiMo.exe同目录
        - direct:0:relative_path
            - repository.cfg.toml 中第0个库 repository 默认在MiMo.exe同目录
            - 有相对目录
        - direct:1 repository.cfg.toml 中第1个库 指向性的(本地操作系统)
        - remote:1 在远端 1 代表远端文件主键 
    - type 资源类型
        - '0' 文件
        - '1' 音频
        - '2' 视频
        - '3' 图片
        - '4' 目录
    - upload 是否已上传 1 已上传 0 未上传