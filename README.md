# dcs
Distributed Cache Service

通讯协议基于 `ABNF` ，定义如下：
```
Command = Op Key | KeyValue; 请求协议
Op = "S" / "G" / "D";
Length = 1*DIGIT;
Key = Length SP 1*VCHAR;
Value = Length SP 1*VCHAR;
KeyValue = Key Value;
Response = [Error] Value; 返回协议
Error = Length SP *VCHAR
```
