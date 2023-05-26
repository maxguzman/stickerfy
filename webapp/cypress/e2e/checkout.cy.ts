/// <reference types="cypress" />

describe('Checkout process', () => {
  it('should buy a couple of stickers', () => {
    cy.visit('http://localhost:5173/')
    cy.findByPlaceholderText('Search').type('cat')
    cy.findByRole('button', { name: /add to cart!/i }).click()
    cy.findByTestId('plus-2').click().click()

    cy.findByTestId('badge').findAllByText('1').should('exist')

    cy.findByPlaceholderText('Search').clear()
    cy.findByPlaceholderText('Search').type('ctm')
    cy.findByRole('button', { name: /add to cart!/i }).click()
    cy.findByTestId('plus-4').click()

    cy.findByTestId('badge').findAllByText('2').should('exist')

    cy.findByTestId('openCheckout').click()

    cy.findByTestId('total').then($elem => {
      const total = $elem.text()
      cy.findByRole('button', { name: /check out/i }).click()
      cy.findByTestId('orderTotal').findAllByText('$ ' + total).should('exist')
    })

    cy.findByRole('button', { name: /close/i }).click()
    cy.findByPlaceholderText('Search').clear()
  })
})
