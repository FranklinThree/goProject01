# goProject01
项目结构：

...暂时先实现初级功能吧

服务端

    server
        main        入口
        server      服务端设计
            Start       服务端初始化
            Send        向服务端发送样例
            *local administrator：    本地管理员功能
        structure   存储结构
            Example
            ExampleSet
                Add         添加人脸实例
                Select      查找人脸实例
                Delete      删除人脸实例
                Set         修改人脸实例
                Maintain    管理数据集
        filetools   文件处理工具
            FormatNamer 格式化命名（用于管理自建的数据库）
            Copy        拷贝文件
            

客户端
（以调试为主）

    client
        main        入口
        client      客户端设计
            Start       客户端初始化
            Get         从服务端取得样例
            Use         生成人脸鉴别样例信息
            Check       接收用户选项后，校验答案
        filetools   文件处理工具
            
        
        
            
        
    