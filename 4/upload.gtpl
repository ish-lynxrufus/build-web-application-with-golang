<!DOCTYPE html>
<html>
<head>
  <title>ファイルアップロード</title>
</head>
<body>
  <form enctype="multipart/form-data" action="/upload" method="post">
    <input type="file" name="uploadfile">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" name="upload">
  </form>
</body>
</html>