import { writable } from 'svelte/store'
import type { Product } from '$lib/models/types'

export const productList = writable<Product[]>([])
export const filteredProductList = writable<Product[]>([])
