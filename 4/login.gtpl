<!DOCTYPE html>
<html>
<head>
  <title></title>
</head>
<body>
  <form action="/login" method="post">
    <span>ユーザー名：</span><input type="text" name="username">
    <span>パスワード：</span><input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" name="ログイン">
  </form>
</body>
</html>