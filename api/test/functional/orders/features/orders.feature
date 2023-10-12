Feature: orders
  In order to use the orders API
  As an API user
  I need to be able to manage orders

  Scenario: should post order
    When I send "POST" request to "/orders" with body:
      """
      {
        "items": [
          {
            "product": {
              "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
              "title": "Product 1",
              "price": 100,
              "image_path": "path/to/file",
              "description": "Description 1"
            },
            "quantity": 1
          },
          {
            "product": {
              "id": "ee665f9a-12d7-40bc-bfa6-0ea0e73db93b",
              "title": "Product 2",
              "price": 200,
              "image_path": "path/to/file",
              "description": "Description 2"
            },
            "quantity": 2
          }
        ]
      }
      """
    Then the response code should be 201

  Scenario: should get orders
    Given the following orders
      """
      [
        {
          "id": "954aacb1-a02c-433f-99c1-9d0a8a745c48",
          "items": [
            {
              "product": {
                "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
                "title": "Product 1",
                "price": 100,
                "image_path": "path/to/file",
                "description": "Description 1"
              },
              "quantity": 1
            }
          ]
        },
        {
          "id": "b0f03340-4ab5-4a6c-96c3-f37a81d8a1ce",
          "items": [
            {
              "product": {
                "id": "ee665f9a-12d7-40bc-bfa6-0ea0e73db93b",
                "title": "Product 2",
                "price": 200,
                "image_path": "path/to/file",
                "description": "Description 2"
              },
              "quantity": 2
            }
          ]
        }
      ]
      """
    When I send "GET" request to "/orders" with body:
      """
      {}
      """
    Then the response code should be 200
    And the response should match json:
      """
      {
        "orders": [
          {
            "id": "954aacb1-a02c-433f-99c1-9d0a8a745c48",
            "items": [
              {
                "product": {
                  "id": "69e3acde-9d90-4eea-8074-8e0d95ed7910",
                  "title": "Product 1",
                  "price": 100,
                  "image_path": "path/to/file",
                  "description": "Description 1"
                },
                "quantity": 1
              }
            ]
          },
          {
            "id": "b0f03340-4ab5-4a6c-96c3-f37a81d8a1ce",
            "items": [
              {
                "product": {
                  "id": "ee665f9a-12d7-40bc-bfa6-0ea0e73db93b",
                  "title": "Product 2",
                  "price": 200,
                  "image_path": "path/to/file",
                  "description": "Description 2"
                },
                "quantity": 2
              }
            ]
          }
        ]
      }
      """
