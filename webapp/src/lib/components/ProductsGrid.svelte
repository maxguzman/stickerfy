<script lang="ts">
	import { filteredProductList, productList } from '$lib/stores/products'
	import ProductCard from '$lib/components/ProductCard.svelte'

	const loadProducts = async () => {
		const res = await fetch('/api/products', {
			method: 'GET',
			headers: { 'Content-Type': 'application/json' }
		})
		const data = await res.json()

		if (res.ok) {
			productList.set(data.products)
			return data
		} else {
			throw new Error()
		}
	}
</script>

{#await loadProducts()}
	<div class="flex items-center justify-center pt-4">
		<div
			style="border-top-color:transparent"
			class="w-4 h-4 border-4 border-gray-600 border-solid rounded-full animate-spin mr-2"
		/>
		Loading products..
	</div>
{:then}
	<div class="grid gap-4 grid-cols-3 grid-rows-1">
		{#each $filteredProductList as product}
			<ProductCard {product} />
		{/each}
	</div>
{:catch error}
	<div class="alert alert-error shadow-lg justify-center">
		<div>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="stroke-current flex-shrink-0 h-6 w-6"
				fill="none"
				viewBox="0 0 24 24"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
				/></svg
			>
			<span class="ml-2">Error calling the API: {error.message}</span>
		</div>
	</div>
{/await}
