/**
 * @jest-environment jsdom
 */

import OpenCart from './OpenCart.svelte'
import { render } from '@testing-library/svelte'
import { shoppingCart } from '$lib/stores/shoppingCart'

describe('Open shopping cart button', () => {
  it('should be disabled when no products in shopping cart', () => {
    shoppingCart.set([])

    const cart = render(OpenCart)

    const button = cart.getByRole('button')
    expect(button).toBeDisabled()
  })

  it('should be enabled and have one product badge when there is one product on cart', () => {
    const mockCart = [{ product: { id: '1', title: '', price: 0 }, quantity: 1 }]
    shoppingCart.set(mockCart)

    const cart = render(OpenCart)

    const button = cart.getByRole('button')
    const badge = cart.getByTestId('badge')
    expect(button).toBeEnabled()
    expect(badge).toHaveTextContent("1")
  })
})