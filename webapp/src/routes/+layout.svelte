<script lang="ts">
	import '../app.css'
	import OpenCart from '$lib/components/OpenCart.svelte'
	import SearchBox from '$lib/components/SearchBox.svelte'
	import Checkout from '$lib/components/Checkout.svelte'

	let title: string = "Stickerfy!"
</script>

<div class="container mx-auto px-4 py-4">
	<div class="navbar mb-2 shadow-lg bg-neutral text-neutral-content rounded-box">
		<div class="flex-1 hidden px-2 mx-2 lg:flex">
			<span class="text-lg font-bold"> {title} </span>
		</div>
		<div class="flex-1 lg:flex-none">
			<SearchBox />
		</div>
		<div class="w-6" />
		<div class="flex-none">
			<OpenCart />
		</div>
	</div>
	<div class="container mx-auto">
		<slot />
	</div>
</div>

<Checkout />
