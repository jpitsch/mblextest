<div>
	<h1>User Info</h1>
	<p>User Name: {{.UserName}}</p>
	<p>Email: {{.Email}}</p>
</div>
<div>
	<h1>Tests</h1>
	<ul style="list-style-type:none">
		{{range $key, $val := .s}}
    	<li>{{$key}}</li>
    	<li>{{$val}}</li>
    	{{end}}
    </ul>
</div>