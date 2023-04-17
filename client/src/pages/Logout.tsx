import { Component } from "solid-js"

import { createAuth } from "../store/auth"

export const Logout: Component = () => {
  const auth = createAuth

  return (
    <div class="flex min-h-[70vh] flex-col items-center justify-center">
      <div class="items-start">
        <button class="btn-lg btn-red" onClick={() => auth.logout()}>
          Logout!
        </button>
        <p>Are you sure?</p>
      </div>
    </div>
  )
}
