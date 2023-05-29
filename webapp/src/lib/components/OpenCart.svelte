<script lang="ts">
	import { shoppingCart, checkoutVisible } from '$lib/stores/shoppingCart'
</script>

<button
	class="btn"
	data-testid="openCheckout"
	on:click={() =>
		$shoppingCart.length > 0 ? ($checkoutVisible = true) : ($checkoutVisible = false)}
	disabled={!$shoppingCart.length}
>
	<i class="fa-solid fa-cart-shopping fa-2x" />
	{#if $shoppingCart.length > 0}
		<span class="badge badge-secondary" data-testid="badge">{$shoppingCart.length}</span>
	{/if}
</button>
