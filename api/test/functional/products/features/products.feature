Feature: products
  In order to use the products API
  As an API user
  I need to be able to manage products

  Scenario: should get products
    Given the following products
      | id                                   | title     | price | image_path   | description   |
      | 69e3acde-9d90-4eea-8074-8e0d95ed7910 | Product 1 | 100   | path/to/file | Description 1 |
      | ee665f9a-12d7-40bc-bfa6-0ea0e73db93b | Product 2 | 200   | path/to/file | Description 2 |
      | b0f03340-4ab5-4a6c-96c3-f37a81d8a1ce | Product 3 | 300   | path/to/file | Description 3 |
    When I send "GET" request to "/products"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "products": [
          {
            "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
            "title": "Product 1",
            "price": 100,
            "image_path": "path/to/file",
            "description": "Description 1"
          },
          {
            "id": "ee665f9a-12d7-40bc-bfa6-0ea0e73db93b",
            "title": "Product 2",
            "price": 200,
            "image_path": "path/to/file",
            "description": "Description 2"
          },
          {
            "id": "b0f03340-4ab5-4a6c-96c3-f37a81d8a1ce",
            "title": "Product 3",
            "price": 300,
            "image_path": "path/to/file",
            "description": "Description 3"
          }
        ]
      }
      """

  Scenario: should get no products
    Given the following products
      | id | title | price | image_path | description |
    When I send "GET" request to "/products"
    Then the response code should be 404
    And the response should match json:
      """
      {
        "code": 404,
        "message": "there where no products found"
      }
      """

  Scenario: should get product by ID
    Given the following products
      | id                                   | title     | price | image_path   | description   |
      | 69e3acde-9d90-4eea-8074-8e0d95ed7910 | Product 1 | 100   | path/to/file | Description 1 |
    When I send "GET" request to "/products/69e3acde-9d90-4eea-8074-8e0d95ed7910"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
        "title": "Product 1",
        "price": 100,
        "image_path": "path/to/file",
        "description": "Description 1"
      }
      """

  Scenario: should post product
    When I send "POST" request to "/products" with body:
      """
      {
        "title": "Product 1",
        "price": 100,
        "image_path": "path/to/file",
        "description": "Description 1"
      }
      """
    Then the response code should be 201

  Scenario: should update product
    Given the following products
      | id                                   | title     | price | image_path   | description   |
      | 69e3acde-9d90-4eea-8074-8e0d95ed7910 | Product 1 | 100   | path/to/file | Description 1 |
    When I send "PUT" request to "/products" with body:
      """
      {
        "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
        "title": "Product 2",
        "price": 200,
        "image_path": "path/to/file",
        "description": "Description 2"
      }
      """
    Then the response code should be 200

  Scenario: should delete product
    Given the following products
      | id                                   | title     | price | image_path   | description   |
      | 69e3acde-9d90-4eea-8074-8e0d95ed7910 | Product 1 | 100   | path/to/file | Description 1 |
    When I send "DELETE" request to "/products" with body:
      """
      {
        "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
        "title": "Product 1",
        "price": 100,
        "image_path": "path/to/file",
        "description": "Description 1"
      }
      """
    Then the response code should be 200
