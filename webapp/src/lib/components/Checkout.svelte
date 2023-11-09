<script lang="ts">
	import { shoppingCart, cleanShoppingCart, checkoutVisible } from '$lib/stores/shoppingCart'

	let cartDetailVisible = true
	let checkingOut = false
	let checkoutButtonVisible = true
	let errorButtonVisible = false
	let orderTotal = 0

	$: total = $shoppingCart.reduce((sum, item) => sum + item.quantity * item.product.price, 0)

	const checkOut = async () => {
		checkingOut = true
		const url = '/api/orders'
		try {
			const result = await fetch(url, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ items: $shoppingCart })
			})
			const data = await result.json()
			orderTotal = total
			cleanShoppingCart()
			checkingOut = false
			cartDetailVisible = false
			checkoutButtonVisible = false
			return { response: data }
		} catch (e) {
			checkoutButtonVisible = false
			errorButtonVisible = true
			return e
		}
	}
</script>

<div class="modal {$checkoutVisible ? 'modal-open' : ''}">
	<div class="modal-box">
		{#if cartDetailVisible}
			<div class="overflow-x-auto">
				<table class="table-fixed w-full text-black">
					<thead>
						<tr>
							<th />
							<th class="text-left">Name</th>
							<th class="text-left">Quantity</th>
							<th class="text-right">Subtotal</th>
						</tr>
					</thead>
					<tbody>
						{#each $shoppingCart as item}
							<tr>
								<td>
									<img
										class="scale-140 h-4"
										src={item.product.image_path}
										alt={item.product.id}
									/>
								</td>
								<td>{item.product.title}</td>
								<td>{item.quantity}</td>
								<td class="text-right" data-testid="subtotal-{item.product.id}"
									>$ {item.quantity * item.product.price}</td
								>
							</tr>
						{/each}
					</tbody>
				</table>
				<div class="text-black text-right font-bold">
					<br />
					TOTAL: $
					<span title="total" data-testid="total">{total}</span>
				</div>
			</div>
		{:else}
			<div><i class="fa-solid fa-shopping-bag fa-2x" /></div>
			<div>
				<p><strong>Thank you for your purchase!</strong></p>
				<p>
					Your order for<span class="badge mx-2 badge-primary" data-testid="orderTotal"
						>$ {orderTotal}</span
					>has been entered.
				</p>
				<p>You'll be notified when your package is ready.</p>
			</div>
		{/if}
		<div class="modal-action">
			<button
				class="btn"
				on:click={() => {
					$checkoutVisible = false
					cartDetailVisible = true
					checkoutButtonVisible = true
					errorButtonVisible = false
					checkingOut = false
				}}
			>
				Close
			</button>
			{#if checkoutButtonVisible}
				<button class="btn btn-primary {checkingOut ? 'loading' : ''}" on:click={checkOut}
					>Check out</button
				>
			{/if}
			{#if errorButtonVisible}
				<button
					class="btn btn-error"
					on:click={() => {
						errorButtonVisible = false
						checkoutButtonVisible = true
						checkingOut = false
					}}>Connection error!</button
				>
			{/if}
		</div>
	</div>
</div>
