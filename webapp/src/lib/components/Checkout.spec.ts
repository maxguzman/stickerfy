/**
 * @jest-environment jsdom
 */

import Checkout from './Checkout.svelte'
import { render, fireEvent } from '@testing-library/svelte'
import { shoppingCart } from '$lib/stores/shoppingCart'

describe('Shopping cart summary', () => {
  it("should display correct sub-total per item", () => {
    const mockCart = [
      { product: { id: '1', title: 'a', price: 10 }, quantity: 5 }
    ]
    shoppingCart.set(mockCart)

    const checkout = render(Checkout)

    const subtotal = checkout.getByTestId('subtotal-1')
    expect(subtotal).toHaveTextContent("50")
  })

  it('should display correct total with multiple items', () => {
    const mockCart = [
      { product: { id: '1', title: 'a', price: 10 }, quantity: 1 },
      { product: { id: '2', title: 'b', price: 5 }, quantity: 2 },
      { product: { id: '3', title: 'c', price: 20 }, quantity: 5 }
    ]
    shoppingCart.set(mockCart)

    const checkout = render(Checkout)

    const total = checkout.getByTitle('total')
    expect(total).toHaveTextContent("120")
  })
})

describe('Checkout process', () => {
  it('should insert the order and receive a confirmation summary', async () => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        json: () => Promise.resolve({error: false, msg:null, order:{}  }),
      }),
    ) as jest.Mock
    const mockCart = [
      { product: { id: '1', title: 'a', price: 10 }, quantity: 5 },
      { product: { id: '1', title: 'a', price: 5 }, quantity: 1 },
    ]
    shoppingCart.set(mockCart)

    const checkout = render(Checkout)
    const button = checkout.getByRole('button', { name: /check out/i })
    await fireEvent.click(button)

    const total = await checkout.findByTestId('orderTotal')
    expect(total).toHaveTextContent('$ 55')
  })

  it('should show an error message when API connection fails', async () => {
    global.fetch = jest.fn(() => Promise.reject()) as jest.Mock
    const mockCart = [
      { product: { id: '1', title: 'a', price: 10 }, quantity: 5 },
    ]
    shoppingCart.set(mockCart)

    const checkout = render(Checkout)
    const button = checkout.getByRole('button', { name: /check out/i })
    await fireEvent.click(button)

    const checkoutButton = await checkout.getByRole('button', { name: /connection error/i })
    expect(checkoutButton).toHaveTextContent('Connection error!')
  })
})

