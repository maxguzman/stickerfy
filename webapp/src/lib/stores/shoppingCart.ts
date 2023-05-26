import type { Product, CartItem } from '$lib/models/types'
import { writable } from 'svelte/store'

export const shoppingCart = writable<CartItem[]>([])

let cart: CartItem[] = []

export const addToCart = (product: Product): void => {
  let cartIndex = 0
  if (cart.length > 0) {
    for (const item of cart) {
      if (item.product.id === product.id) {
        item.quantity += 1
        item.product = product
        cart[cartIndex] = item
        shoppingCart.set(cart)
        return
      }
      cartIndex++
    }
  }
  const item: CartItem = { product: product, quantity: 1 }
  cart = [...cart, item]
  shoppingCart.set(cart)
}

export const minusItem = (product: Product): void => {
  let cartIndex = 0
  if (cart.length > 0) {
    for (const item of cart) {
      if (item.product.id === product.id) {
        if (item.quantity > 1) {
          item.quantity -= 1
          cart[cartIndex] = item
          shoppingCart.set(cart)
        } else {
          cart = cart.filter((cartItem) => cartItem.product.id != product.id)
          shoppingCart.set(cart)
        }
        return
      }
      cartIndex++
    }
  }
}

export const plusItem = (product: Product): void => {
  let cartIndex = 0
  if (cart.length > 0) {
    for (const item of cart) {
      if (item.product.id === product.id) {
        item.quantity += 1
        cart[cartIndex] = item
        shoppingCart.set(cart)
        return
      }
      cartIndex++
    }
  }
}

export const cleanShoppingCart = (): void => {
  shoppingCart.set([])
  cart = []
}

export const checkoutVisible = writable<boolean>(false)
