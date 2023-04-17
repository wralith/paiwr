import { A } from "@solidjs/router"
import { Component, createEffect, createSignal, For } from "solid-js"

import { createAuth } from "../store/auth"

// TODO: Could be declared in one place and used in App.tsx and here
type Route = { name: string; path: string }
const commonRoutes: Route[] = [{ name: "Home", path: "/" }]

const guestRoutes: Route[] = [
  { name: "Login", path: "/login" },
  { name: "Register", path: "/register" },
]

const authenticatedRoutes: Route[] = [{ name: "Logout", path: "/logout" }]

export const Navbar: Component<{ toggleDark: () => void }> = (props) => {
  const auth = createAuth
  const [routes, setRoutes] = createSignal<Route[]>([])
  createEffect(() => {
    setRoutes([...commonRoutes, ...(auth.isLoggedIn() ? authenticatedRoutes : guestRoutes)])
  })

  return (
    <nav class="flex w-full items-center justify-between p-2 md:justify-around">
      <A class="rounded-lg p-2 text-lg font-bold transition-colors hover:underline" href="/">
        Paiwr
      </A>
      <ul class="flex gap-4">
        <For each={routes()}>
          {(route) => (
            <li>
              <A class="rounded-lg p-2 transition-colors hover:underline" href={route.path}>
                {route.name}
              </A>
            </li>
          )}
        </For>
        <li>
          <span
            class="cursor-pointer rounded-lg p-2 transition-colors hover:underline"
            onClick={() => props.toggleDark()}
          >
            Theme
          </span>
        </li>
      </ul>
    </nav>
  )
}
