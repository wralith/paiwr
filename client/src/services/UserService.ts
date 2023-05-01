import { createFetch } from "../signals/createFetch"

async function login(fields: { username: string; password: string }, login: (jwt: string) => void) {
  const url = new URL("users/login", import.meta.env.VITE_SERVER_BASE_URL)
  const res = await createFetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(fields),
  })

  const json = await res.json()
  login(json.token)

  return res.ok
}

async function register(fields: { username: string; password: string; email: string }) {
  const url = new URL("users/register", import.meta.env.VITE_SERVER_BASE_URL)
  const res = await createFetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(fields),
  })

  return res.ok
}

export const UserService = {
  login,
  register,
}
