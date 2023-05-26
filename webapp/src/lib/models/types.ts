export type Product = {
  id: string
  image_path?: string
  title: string
  description?: string
  price: number
}

export type CartItem = {
  product: Product
  quantity: number
}

export type OrderResponse = {
  order: string
  error: boolean
  msg: string
}

