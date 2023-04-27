import { useNavigate } from "@solidjs/router"
import { Component, JSX } from "solid-js"

import { createAuth } from "../store/auth"

// TODO: maybe redirect some other place

export const OnlyUsers: Component<{ children: JSX.Element }> = (props) => {
  const auth = createAuth
  if (!auth.isLoggedIn()) {
    const navigate = useNavigate()
    navigate("/", { replace: true })
  }

  return <>{props.children}</>
}

export const OnlyGuests: Component<{ children: JSX.Element }> = (props) => {
  const auth = createAuth
  if (auth.isLoggedIn()) {
    const navigate = useNavigate()
    navigate("/", { replace: true })
  }

  return <>{props.children}</>
}
