<div>
	<form id="testForm" action="" method="POST">
		<div><p>Create New Test</p></div>
		<div><label>Test Name:</label><input name="testName" type="text" required /></div>
		<div><label>Test Type:</label><input name="testType" type="text" required /></div>
		<div>
			<input type="button" value="create" name="create" onclick="startAddingQuestions()" />
		</div>
	</form>
</div>
<script>
	form=document.getElementById("testForm");
	function startAddingQuestions() {
	    form.action="/admin/test/startaddingquestions";
	    form.submit();
	}
</script>