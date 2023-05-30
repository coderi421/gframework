# 错误码

！！系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrConnectDB | 100601 | 500 | Init db error |
| ErrConnectGRPC | 100602 | 500 | Connect to grpc error |
| ErrGoodsNotFound | 100501 | 404 | Goods not found |
| ErrCategoryNotFound | 100502 | 404 | Category not found |
| ErrEsUnmarshal | 100503 | 500 | Es unmarshal error |
| ErrInventoryNotFound | 100601 | 404 | Inventory not found |
| ErrInvSellDetailNotFound | 100602 | 400 | Inventory sell detail not found |
| ErrInvNotEnough | 100603 | 400 | Inventory not enough |
| ErrShopCartItemNotFound | 100701 | 404 | ShopCart item not found |
| ErrSubmitOrder | 100702 | 400 | Submit order error |
| ErrNoGoodsSelect | 100703 | 404 | No Goods selected |
| ErrUserNotFound | 100401 | 404 | User not found |
| ErrUserAlreadyExists | 100402 | 400 | User already exists |
| ErrUserPasswordIncorrect | 100403 | 400 | User password incorrect |
| ErrSmsSend | 100404 | 400 | Send sms error |
| ErrCodeNotExist | 100405 | 400 | Sms code incorrect or expired |
| ErrCodeInCorrect | 100406 | 400 | Sms code incorrect |
| ErrCodeInvalidParam | 100407 | 400 | Invalid param |

