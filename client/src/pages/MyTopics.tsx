import { createForm, zodForm } from "@modular-forms/solid"
import { Component, createResource, createSignal, For, Show } from "solid-js"
import { z } from "zod"

import { Input } from "../components/Input"
import { createTopicSchema, TopicService } from "../services/TopicService"
import { createAuth } from "../store/auth"

// TODO: Do table or something
const MyTopics: Component = () => {
  const [showForm, setShowForm] = createSignal(false)
  const auth = createAuth
  // if auth.user() === undefined will redirect to "/" so "!" operator is ok in here
  const [topics, { refetch }] = createResource(() =>
    TopicService.findTopicsByOwnerId(auth.user()!.id!)
  )

  const [, { Form, Field }] = createForm<z.input<typeof createTopicSchema>>({
    validate: zodForm(createTopicSchema),
  })

  const onSend = async (input: z.input<typeof createTopicSchema>) => {
    const ok = await TopicService.createTopic(input)
    console.log(input)
    if (ok) {
      refetch()
    }
  }

  return (
    <div class="flex min-h-[70vh] items-center justify-center">
      <div class="flex w-full flex-col items-start justify-center gap-4 p-4 md:w-1/2 lg:w-1/3">
        <button class="btn btn-violet" onClick={() => setShowForm(!showForm())}>
          {showForm() ? "Close Form" : "Crate New Topic"}
        </button>
        <Show when={showForm()}>
          <h1 class="text-lg font-bold">
            {showForm() ? "Crate New Topic" : "Close Create New Topic Form"}
          </h1>
          <Form class="flex w-full flex-col gap-2" onSubmit={(values) => onSend(values)}>
            <Field name="title">
              {(field, props) => (
                <Input
                  type="text"
                  {...props}
                  placeholder="Title"
                  value={field.value}
                  error={field.error}
                  required
                />
              )}
            </Field>
            <Field name="capacity" type="number">
              {(field, props) => (
                <Input
                  type="number"
                  {...props}
                  placeholder="Capacity"
                  value={field.value as number | string}
                  error={field.error}
                  required
                />
              )}
            </Field>
            <Field name="category">
              {(field, props) => (
                <select {...props}>
                  <For
                    each={[
                      { label: "Software", value: "software" },
                      { label: "Social Sciences", value: "social_sciences" },
                      { label: "Other", value: "other" },
                    ]}
                  >
                    {({ label, value }) => (
                      <option value={value} selected={field.value === value}>
                        {label}
                      </option>
                    )}
                  </For>
                </select>
              )}
            </Field>
            <input class="btn btn-violet" type="submit" />
          </Form>
        </Show>
        <div>
          <h1 class="text-lg font-bold">My Topics</h1>
          <div>
            <pre>{JSON.stringify(topics(), null, 2)}</pre>
          </div>
        </div>
      </div>
    </div>
  )
}

export default MyTopics
