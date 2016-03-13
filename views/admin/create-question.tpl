<div>
	<h4>{{.flash.notice}}</h4>
</div>
<form id="questionForm" action="" method="POST">
	<div><label>Question Text:</label><input name="question" type="text" required /></div>
	<div><label>Answer 1:</label><input name="answer1" type="text" required /></div>
	<div><label>Answer 2:</label><input name="answer2" type="text" required /></div>
	<div><label>Answer 3:</label><input name="answer3" type="text" required /></div>
	<div><label>Answer 4:</label><input name="answer4" type="text" required /></div>
	<div>
		<label>Correct Answer:</label>
		<select>
			<option value=""></option>
			<option value="1">1</option>
			<option value="2">2</option>
			<option value="3">3</option>
			<option value="4">4</option>
		</select>
	</div>
	<div>
		<input type="button" value="Next" name"next" onclick="nextQuestion()" />
		<input type="button" value="Next" name"next" onclick="previousQuestion()" />
		<input type="button" value="Next" name"next" onclick="finishedTest()" />
	</div>
</form>
<script>
	form=document.getElementById("questionForm");
	function nextQuestion() {
	    form.action="/admin/test/addnextquestion";
	    form.submit();
	}
	function previousQuestion() {
	    form.action="/admin/test/previousquestion";
	    form.submit();
	}
	function finishedTest() {
		form.action="/admin/test/finalizetest";
		form.submit();
	}
</script>