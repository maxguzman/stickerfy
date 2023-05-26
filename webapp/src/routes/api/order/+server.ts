import type { RequestEvent } from '@sveltejs/kit'

const apiURL = process.env.STICKERFY_SERVICE_URL || "http://localhost:8000/v1"

export const POST = async (req: RequestEvent): Promise<Response> => {
  try {
    const res = await (fetch(apiURL + "/order", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(await req.request.json().then(body => body)),
    }
    ))
    const data = await res.json()
    return new Response(JSON.stringify(data), { status: 200 })
  } catch (e) {
    console.log(e)
    return new Response(JSON.stringify({ error: e }), { status: 500 })
  }
}
