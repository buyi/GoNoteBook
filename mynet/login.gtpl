<html>
<head><title></title></head>
<body>
<form enctype="multipart/form-data" action="http://192.168.31.214:9090/upload"
	method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="upload" />
</form>

<form action="http://192.168.31.214:9090/login" method="post">
用户名:<input type="text" name="uname"></br>
密码:<input type="password" name="passwd"></br>
<input type="hidden" name="token" value="{{.}}">


<select name="fruit">
<option value="apple">apple</option>
<option value="pear">pear</option>
<option value="banana">banana</option>
</select></br>


<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女</br>


<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球</br>


<input type="submit" value="登录">

</form>

</body>



</html>