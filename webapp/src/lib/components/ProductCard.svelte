<script lang="ts">
	import type { Product, CartItem } from '$lib/models/types';
	import { shoppingCart, addToCart, plusItem, minusItem } from '$lib/stores/shoppingCart';
	export let product: Product;
	$: currentItem = $shoppingCart.find(
		(item: CartItem): Boolean => (item.product.id === product.id ? true : false)
	)
</script>

<div class="card bordered shadow-2xl bg-neutral-content">
	<figure>
		<img class="object-contain h-72 w-120 scale-50" src={product.image_path} alt={product.title} />
	</figure>
	<div class="card-body">
		<h2 class="card-title" data-testid="title">
			{product.title}
			<span class="badge mx-2 badge-primary badge-outline" data-testid="price"
				>$ {product.price}</span
			>
		</h2>
		<p data-testid="description">
			{product.description}
		</p>
		{#if currentItem == undefined}
			<div class="justify-end card-actions">
				<button
					class="btn btn-primary"
					data-testid="add-{product.id}"
					on:click={() => {
						addToCart(product);
					}}
				>
					Add to cart!
				</button>
			</div>
		{/if}
		{#if currentItem != undefined}
			<div class="justify-end card-actions">
				<div class="btn-group">
					<button
						class="btn btn-sms"
						data-testid="minus-{product.id}"
						on:click={() => {
							minusItem(product);
						}}
					>
						<i class="fa-solid fa-minus" />
					</button>
					<button class="btn btn-sms">{currentItem.quantity}</button>
					<button
						class="btn btn-sms"
						data-testid="plus-{product.id}"
						on:click={() => {
							plusItem(product);
						}}
					>
						<i class="fa-solid fa-plus" />
					</button>
				</div>
			</div>
		{/if}
	</div>
</div>
