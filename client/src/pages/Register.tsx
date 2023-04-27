import { createForm, zodForm } from "@modular-forms/solid"
import { useNavigate } from "@solidjs/router"
import { Component } from "solid-js"
import { z } from "zod"

import { Input } from "../components/Input"
import { UserService } from "../services/UserService"

const registerSchema = z.object({
  email: z.string().email("Email is invalid"),
  username: z.string().min(3, "Username must have 3 or more characters."),
  password: z.string().min(6, "Password must have 6 or more characters."),
})

const Register: Component = () => {
  const navigate = useNavigate()

  const [, { Form, Field }] = createForm<z.input<typeof registerSchema>>({
    validate: zodForm(registerSchema),
  })

  const onSend = async (input: z.input<typeof registerSchema>) => {
    const ok = await UserService.register(input)
    if (ok) {
      navigate("/login")
    }
  }

  return (
    <div class="flex min-h-[70vh] items-center justify-center">
      <div class="flex w-full flex-col justify-center gap-2 p-4 md:w-1/2 lg:w-1/3">
        <h1 class="pb-4 text-2xl font-semibold">Register</h1>
        <Form class="flex flex-col gap-2" onSubmit={(values) => onSend(values)}>
          <Field name="username">
            {(field, props) => (
              <Input
                type="text"
                {...props}
                placeholder="Username"
                value={field.value}
                error={field.error}
                required
              />
            )}
          </Field>
          <Field name="email">
            {(field, props) => (
              <Input
                type="email"
                {...props}
                placeholder="Email"
                value={field.value}
                error={field.error}
                required
              />
            )}
          </Field>
          <Field name="password">
            {(field, props) => (
              <Input
                type="password"
                {...props}
                placeholder="Password"
                value={field.value}
                error={field.error}
                required
              />
            )}
          </Field>
          <input class="btn btn-violet" type="submit" />
        </Form>
      </div>
    </div>
  )
}

export default Register
