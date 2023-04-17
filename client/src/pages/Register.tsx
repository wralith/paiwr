import { useNavigate } from "@solidjs/router"
import { Component } from "solid-js"
import { createStore } from "solid-js/store"

import { Input } from "../components/Input"
import { UserService } from "../services/UserService"

export const Register: Component = () => {
  const [fields, setFields] = createStore<{ username: string; email: string; password: string }>({
    username: "",
    email: "",
    password: "",
  })
  const navigate = useNavigate()

  const onSend = async () => {
    const ok = await UserService.register(fields)
    if (ok) {
      navigate("/login")
    }
  }

  return (
    <div class="flex min-h-[70vh] items-center justify-center">
      <div class="flex w-full flex-col justify-center gap-2 p-4 md:w-1/2 lg:w-1/3">
        <h1 class="pb-4 text-2xl font-semibold">Register</h1>
        <Input
          type="text"
          placeholder="Username"
          value={fields.username}
          onInput={(e) => setFields("username", e.currentTarget.value)}
        />
        <Input
          type="email"
          placeholder="Email"
          value={fields.email}
          onInput={(e) => setFields("email", e.currentTarget.value)}
        />
        <Input
          type="password"
          placeholder="Password"
          value={fields.password}
          onInput={(e) => setFields("password", e.currentTarget.value)}
        />
        <button onClick={onSend} class="btn btn-violet">
          Send
        </button>
      </div>
    </div>
  )
}
