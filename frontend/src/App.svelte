<script>
	
</script>

<main>
	<div class="welc">
			<div class="title">Welcome</div>
			<div class="suptitle">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</div>
	</div>
	<div class="note_block">
		<div class="title_note">NOTES</div>
		<div class="notes_box">
			<div class="add_notes">
				<h2 class="add_note">Add note</h2>
				<input type="text" class="input_note"> 
				<button class="add_btn"><div class="btn_text">Add</div></button>
			</div>
			<hr class="hr">
			<div class="notes">
				<ul>
					<li> 
						<div class="text">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</div> 
					</li>
				</ul> 
			</div>
		</div> <!--notes_box-->
	</div> <!--note_block-->
</main>

<style>

	/*base*/
	main {
		margin: 0px;
		padding: 0px;
	}

	/*welc*/
	.welc {
		padding-bottom: 100px;
	}

	/*text*/
	.title {
		font-family: 'Bebas Neue', cursive;
		text-align: center;
		font-size: 100px;
		color: #ccc;
	}

	.suptitle {
		font-size: 30px;
		text-align: center;
		max-width: 900px;
		font-family: 'Barlow Condensed', sans-serif;
		color: #ccc;
		margin: auto;
		padding: 50px;
	}

	.title_note {
		font-family: 'Bebas Neue', cursive;
		text-align: center;
		font-size: 100px;
		margin-top: 50px;
		color: #ccc;

		padding-bottom: 30px;
	}

	.add_note {
		padding-top: 30px;
		margin: 0 40px;

		font-size: 35px;
		font-weight: 700;
		color: #333;
		font-family: 'Montserrat', sans-serif;
	}

	.text {
		font-family: 'Barlow Condensed', sans-serif;
		font-weight: bold;
		font-size: 30px;

		max-width: 1000px;
		margin: 50px;
	}

	/*box*/
	.notes_box {
		width: 80%;
		margin: 0 200px;
		background: rgb(221, 221, 231);

		border: solid 3px;
		border-color: black;
		border-radius: 2em;
	}

	/*input*/

	.input_note {
		margin: 25px 30px;
		width: 1100px;
		background: rgb(97, 104, 128);	
		color: rgb(36, 31, 31);
		font-weight: bold;
		height: 40px;

		border: solid 2px;
		border-color: black;
		font-family: 'Montserrat', sans-serif;
	}

	/*add btn*/
	.add_btn {
		border: solid 2px;
		border-radius: 1em;
		border-color: #333;

		background: rgb(47, 206, 47);
		font-size: 22px;
		margin: 0 100px;
	}

	.add_btn:active {
		background: rgb(38, 155, 38);
	}

	.btn_text{
		font-family: 'Barlow Condensed', sans-serif;
		padding: 0px 30px;
		font-weight: bold;
	}

	/*delite btn*/
	.del_btn {
		
	}
	/*line*/
	.hr {
		height: 2px;
		background: black;
		width: 1500px;
		margin-top: 20px;
		margin-bottom: 30px;
		border: #ccc;
	}
</style>