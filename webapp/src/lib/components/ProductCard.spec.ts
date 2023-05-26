/**
 * @jest-environment jsdom
 */

import ProductCard from './ProductCard.svelte'
import { render, fireEvent } from '@testing-library/svelte'
import { shoppingCart } from '$lib/stores/shoppingCart'

describe('Product card', () => {
  it('should render the product correctly', () => {
    const mockProduct = { id: '1', image_path: 'image.jpg', title: 'dummy product', description: 'dummy description', price: 10 }

    const card = render(ProductCard, { product: mockProduct })

    expect(card.getByAltText(/dummy product/i)).toBeVisible()
    expect(card.getByTestId('title')).toHaveTextContent('dummy product')
    expect(card.getByTestId('price')).toHaveTextContent('10')
    expect(card.getByTestId('description')).toHaveTextContent('dummy description')
  })

  it('should add a new product, morph to add more/less buttons when pressing Add to cart!', async () => {
    const mockProduct = { id: '1', image_path: 'image.jpg', title: 'product-1', description: 'dummy product', price: 10 }

    const card = render(ProductCard, { product: mockProduct })
    const addButton = card.getByRole('button', { name: /add to cart/i })
    await fireEvent.click(addButton)

    expect(shoppingCart.subscribe.length).toEqual(1)
    expect(await card.getByRole('button', { name: /1/i })).toBeVisible()
  })
})
