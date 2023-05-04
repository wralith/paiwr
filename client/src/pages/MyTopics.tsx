import { createForm, zodForm } from "@modular-forms/solid"
import { Component, createResource, createSignal, For, Show } from "solid-js"
import { z } from "zod"

import { Input } from "../components/Input"
import { TopicCard } from "../components/TopicCard"
import { createTopicSchema, TopicService } from "../services/TopicService"
import { createAuth } from "../store/auth"

const MyTopics: Component = () => {
  const [showForm, setShowForm] = createSignal(false)
  const auth = createAuth
  // if auth.user() === undefined will redirect to "/" so "!" operator is ok in here
  const [myTopics, { refetch: refetchMyTopics }] = createResource(() =>
    TopicService.findTopicsByOwnerId(auth.user()!.id!)
  )
  const [involvedTopics, { refetch: refecthInvolvedTopcis }] = createResource(async () =>
    TopicService.excludeUserOwnedTopics(
      await TopicService.findPairedTopicsByUserId(auth.user()!.id!),
      auth.user()!.id!
    )
  )

  const refetch = async () => {
    await refetchMyTopics()
    await refecthInvolvedTopcis()
  }

  const [, { Form, Field }] = createForm<z.input<typeof createTopicSchema>>({
    validate: zodForm(createTopicSchema),
  })

  const onSend = async (input: z.input<typeof createTopicSchema>) => {
    const ok = await TopicService.createTopic(input)
    if (ok) {
      await refetch()
    }
  }

  return (
    <div class="flex min-h-[70vh] items-center justify-center">
      <div class="flex w-full flex-col items-start justify-center gap-4 p-4 md:w-1/2">
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
                  value={field.value as string | undefined}
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
        <Show when={myTopics() && myTopics()!.length > 0}>
          <h1 class="text-lg font-bold">My Topics</h1>
          <div class="grid min-w-full flex-col gap-2 xl:grid-cols-2">
            <For each={myTopics()!}>
              {(topic) => <TopicCard topic={topic} refetch={refetch} deletable />}
            </For>
          </div>
        </Show>
        <Show when={involvedTopics() && involvedTopics()!.length > 0}>
          <h1 class="text-lg font-bold">Involved Topics</h1>
          <div class="grid min-w-full flex-col gap-2 xl:grid-cols-2">
            <For each={involvedTopics()!}>
              {(topic) => <TopicCard topic={topic} refetch={refetch} />}
            </For>
          </div>
        </Show>
      </div>
    </div>
  )
}

export default MyTopics
