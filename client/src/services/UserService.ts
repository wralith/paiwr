async function login(fields: { username: string; password: string }, login: (jwt: string) => void) {
  const res = await fetch("http://localhost:8080/users/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(fields),
  })

  const json = await res.json()
  login(json.token)

  return res.ok
}

async function register(fields: { username: string; password: string; email: string }) {
  const res = await fetch("http://localhost:8080/users/register", {
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
