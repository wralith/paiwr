import { Component, Show } from "solid-js"

import { TopicResponse, TopicService } from "../services/TopicService"

type Props = {
  topic: TopicResponse
  refetch: () => void
  deletable?: boolean
}

// TODO: Put delete button into a popup menu and show some warning before send delete request

export const TopicCard: Component<Props> = (props) => {
  const createdAt = () => new Date(props.topic.created_at)
  const onDelete = async () => {
    await TopicService.deleteTopicById(props.topic.id)
    props.refetch()
  }

  return (
    <div class="flex flex-col gap-6 rounded-md border border-gray-300 p-4 transition-colors hover:bg-gray-100 dark:border-gray-600 dark:hover:bg-gray-600">
      <div>
        <div class="flex items-baseline justify-between">
          <h4 class="text-lg">{props.topic.title}</h4>
          <Show when={props.deletable}>
            <button
              onClick={onDelete}
              class="text-xs text-red-600 dark:text-red-400"
              title={`Delete ${props.topic.title}`}
            >
              Delete
            </button>
          </Show>
        </div>
        <a class="cursor-pointer text-xs font-semibold text-violet-700 hover:underline dark:text-violet-300">
          {props.topic.category}
        </a>
      </div>
      <div class="flex items-baseline justify-between">
        <p class="text-sm">
          <span class="text-gray-400">Capacity:</span> {props.topic.capacity}
        </p>
        <p>{`${createdAt().toLocaleDateString("tr-TR")}`}</p>
      </div>
    </div>
  )
}
