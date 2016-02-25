<!DOCTYPE html>
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<h1>Login Info</h1>
		<form action="/login" method="POST">
			<div><label>User Name:</label><input name="username" type="text" required /></div>
			<div><label>Password:</label><input name="password" type="password" required /></div>
			<div><input class="button" type="submit" value="Sign In" /></div>
		</form>
	</body>
</html>