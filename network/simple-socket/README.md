# TCP Simple Command
TCP本身的循环逻辑比较简单
Server:
Listen() → Accept() → go connHandler(conn) → conn.Read() → conn.Write() → conn.Close()

Client:
Dial() → connHandler(conn) → conn.Read() → conn.Write()

Server和Client端都是通过读写循环来收发命令的, 整个过程跟WebSocket相似.
这里能明显的感受到TCP和上层HTTP的区别, 理论上两者功能相似, 都是通过收发命令来传输信息. 但是TCP这层需要考虑循环逻辑, 由于是block IO过程, 所以处理不当容易导致程序卡死. 而HTTP在TCP上添加了更多规范

UDP和TCP的区别: UDP由于为无连接协议, 所以每次都获取conn地址


## reference 
https://juejin.im/entry/5aa8ebe46fb9a028de4467bd