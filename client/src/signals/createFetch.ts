import { createAuth } from "../store/auth"

// It doesn't seem like a signal though...
export async function createFetch(input: RequestInfo | URL, init?: RequestInit): Promise<Response> {
  const auth = createAuth

  const fetchOpts = {
    ...init,
    headers: {
      ...init?.headers,
      Authorization: auth.isLoggedIn() ? `Bearer ${auth.jwt()}` : "",
    },
  }

  const res = await fetch(input, fetchOpts)
  if (res.status === 401) {
    auth.logout()
  }

  return res
}
