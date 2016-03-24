<p>Test {{.TestName}}</p>
<div>
	<h4>{{.flash.notice}}</h4>
</div>
<form id="questionForm" action="" method="POST">
	<div><label>Question Text:</label><textarea name="question" type="text" required ></textarea></div>
	<div><label>Answer 1:</label><textarea name="answer1" type="text" required ></textarea></div>
	<div><label>Answer 2:</label><textarea name="answer2" type="text" required ></textarea></div>
	<div><label>Answer 3:</label><textarea name="answer3" type="text" required ></textarea></div>
	<div><label>Answer 4:</label><textarea name="answer4" type="text" required ></textarea></div>
	<div>
		<label>Correct Answer #:</label>
		<select name="correctAnswer">
			<option value=""></option>
			<option value="1">1</option>
			<option value="2">2</option>
			<option value="3">3</option>
			<option value="4">4</option>
		</select>
	</div>
	<div>
		<input type="button" value="Previous" name="previous" onclick="previousQuestion()" />
		<input type="button" value="Next" name="next" onclick="nextQuestion()" />
		<input type="button" value="Finish" name="finished" onclick="finishedTest()" />
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