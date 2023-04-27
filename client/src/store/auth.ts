import { createRoot } from "solid-js"

import { createLocalStorageSignal } from "../signals/createLocalStorageSignal"

type User = { id: string; username: string }

// TODO: Maybe store it somewhere not the memory?
function auth() {
  const [isLoggedIn, setIsLoggedIn] = createLocalStorageSignal(false, "is-logged-in")
  const [user, setUser] = createLocalStorageSignal<User | undefined>(undefined, "user")
  const [jwt, setJwt] = createLocalStorageSignal<string | undefined>(undefined, "user-token")

  const login = (jwt: string) => {
    const tokens = jwt.split(".")
    const userData = JSON.parse(atob(tokens[1])) as {
      name: string
      sub: string
      exp: number
    }
    setJwt(jwt)

    setUser({ id: userData.sub, username: userData.name })
    setIsLoggedIn(true)
  }

  const logout = () => {
    setUser(undefined)
    setIsLoggedIn(false)
    setJwt(undefined)
  }

  return { isLoggedIn, user, login, logout, jwt }
}

export const createAuth = createRoot(auth)
