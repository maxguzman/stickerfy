const apiURL = process.env.STICKERFY_SERVICE_URL || "http://localhost:8000"

export const GET = async (): Promise<Response> => {
  try {
    const res = await fetch(apiURL + "/products", {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    })
    const data = await res.json()
    return new Response(JSON.stringify(data), { status: 200 })
  } catch (e) {
    console.log(e)
    return new Response(JSON.stringify({ error: e }), { status: 500 })
  }
}
