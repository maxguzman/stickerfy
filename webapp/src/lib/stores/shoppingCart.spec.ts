import { get } from 'svelte/store'
import { cleanShoppingCart, shoppingCart, addToCart, minusItem, plusItem } from "./shoppingCart"

const mockProduct = { product: { id: '1', title: 'a', price: 10 }, quantity: 1 }
const mockCart = [
  { product: { id: '1', title: 'a', price: 10 }, quantity: 1 },
  { product: { id: '2', title: 'b', price: 20 }, quantity: 1 },
]

test("render shopping cart", () => {
  shoppingCart.set(mockCart)
  expect(get(shoppingCart)).toContainEqual(mockProduct)
})

test("minus items to cart", () => {
  let sc = get(shoppingCart)
  minusItem(mockProduct.product)
  sc = get(shoppingCart)
  expect(sc[0].quantity).toEqual(1)
  expect(sc.length).toBe(2)
  minusItem(mockProduct.product)
  minusItem(mockProduct.product)
  expect(sc[0].quantity).toBe(1)
})

test("plus items to cart", () => {
  cleanShoppingCart()
  addToCart(mockProduct.product)
  let sc = get(shoppingCart)
  expect(sc[0].quantity).toBe(1)
  plusItem(mockProduct.product)
  sc = get(shoppingCart)
  expect(sc[0].quantity).toEqual(2)
})

test("clean shopping cart", () => {
  cleanShoppingCart()
  expect(get(shoppingCart)).toEqual([])
})