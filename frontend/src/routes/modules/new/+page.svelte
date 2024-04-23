<form method="POST" class="w-full">
    <label class="form-control grow px-4">
        <div class="label">
            <span class="label-text">Name(link or folder name):</span>
        </div>
        <input type="text" name="name" placeholder="Type here" class="input input-bordered w-full " />
    </label>
	<div class="flex w-full gap-4 px-4">
		<label class="form-control grow">
			<div class="label">
				<span class="label-text">Branch</span>
			</div>
			<input type="text" name="branch" placeholder="Type here" class="input input-bordered w-full " />
		</label>
		<label class="form-control grow">
			<div class="label">
				<span class="label-text">Commit</span>
			</div>
			<input type="text" name="commit" placeholder="Type here" class="input input-bordered w-full" />
		</label>
		<button type="submit" class="btn self-end w-1/12" name="action">Checkout</button>
	</div>
</form>