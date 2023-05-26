<script lang="ts">
	import { productList, filteredProductList } from '$lib/stores/products'

	let searchTerm: string = ''

	$: {
		if (searchTerm) {
			filteredProductList.update((l) =>
				l.filter((product) => product.title.toLowerCase().includes(searchTerm.toLocaleLowerCase()))
			)
		} else {
			filteredProductList.set($productList)
		}
	}
</script>

<div class="form-control">
	<input type="text" placeholder="Search" class="input input-ghost" bind:value={searchTerm} />
</div>
