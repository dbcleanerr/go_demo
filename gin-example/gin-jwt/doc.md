### json 停牌结构

- Header
- Payload
- Sgnature

#### Header

> `Header`一般由两部分组成:令牌的类型以及所使用的签名算法，例如HMAC、SHA256、HS256(默认)。bas64URL编码

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

#### Payload

> 有效载荷，就是实际的内容部分,bas64URL编码.

标准定义的基本字段，也可以添加自己的字段
- iss(Issuer): 签发者
- sub(Subject): 主题
- aud(Audience): 受众
- exp(ExpiresAt): 过期时间
- iat(Issued At): 签发时间

#### Signature

> `Signature`是对前两部分的签名，**防止数据篡改**

```javascript
HMACSHA256(
base64UrlEncode(header) + "." +
base64UrlEncode(payload),
secret)
```