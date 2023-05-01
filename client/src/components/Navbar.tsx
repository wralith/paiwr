import { A } from "@solidjs/router"
import { IconMoon, IconSun } from "@tabler/icons-solidjs"
import { Component, createEffect, createSignal, For, Match, Show, Switch } from "solid-js"

import { createAuth } from "../store/auth"

// TODO: Could be declared in one place and used in App.tsx and here
type Route = { name: string; path: string }
const commonRoutes: Route[] = [{ name: "Home", path: "/" }]

const guestRoutes: Route[] = [
  { name: "Login", path: "/login" },
  { name: "Register", path: "/register" },
]

const authenticatedRoutes: Route[] = [{ name: "My Topics", path: "/my-topics" }]

export const Navbar: Component<{ dark: boolean; toggleDark: () => void }> = (props) => {
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
      <ul class="flex items-center gap-4">
        <For each={routes()}>
          {(route) => (
            <li>
              <A class="rounded-lg p-2 transition-colors hover:underline" href={route.path}>
                {route.name}
              </A>
            </li>
          )}
        </For>
        <Show when={auth.isLoggedIn()}>
          <li
            class="rounded-lg p-2 transition-colors hover:underline"
            onClick={() => auth.logout()}
          >
            Logout
          </li>
        </Show>
        <li>
          <span class="cursor-pointer" onClick={() => props.toggleDark()}>
            <Switch>
              <Match when={props.dark}>
                <IconSun size="34" class="p-2 transition-colors hover:text-orange-600" />
              </Match>
              <Match when={!props.dark}>
                <IconMoon size="34" class="p-2 transition-colors hover:text-cyan-600" />
              </Match>
            </Switch>
          </span>
        </li>
      </ul>
    </nav>
  )
}
