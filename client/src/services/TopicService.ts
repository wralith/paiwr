import { z } from "zod"

import { createFetch } from "../signals/createFetch"

export type TopicResponse = {
  id: string
  capacity: number
  category: TopicCategories
  owner: string
  parties: string[]
  title: string
  created_at: string
  finished_at: string
  updated_at: string
}

const topicCategoriesEnum = z.enum(["software", "social_sciences", "other"])
export type TopicCategories = z.infer<typeof topicCategoriesEnum>

export const createTopicSchema = z.object({
  capacity: z
    .number({ invalid_type_error: "Capacity should contain valid number" })
    .min(1, "Capacity must be at least 1"),
  category: topicCategoriesEnum,
  title: z.string().min(3, "Title must have 3 or more characters"),
})

type createTopicInput = z.input<typeof createTopicSchema>

async function createTopic(fields: createTopicInput) {
  const url = new URL("topics", import.meta.env.VITE_SERVER_BASE_URL)
  const res = await createFetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(fields),
  })

  return res.ok
}

async function findTopicsByOwnerId(ownerId: string): Promise<TopicResponse[]> {
  let url = new URL("topics/owner/", import.meta.env.VITE_SERVER_BASE_URL)
  url = new URL(ownerId, url)

  const res = await createFetch(url)

  if (res.status === 400) {
    throw new Error("Invalid owner id")
  }

  return res.json()
}

async function findPairedTopicsByUserId(ownerId: string): Promise<TopicResponse[]> {
  let url = new URL("topics/pair/", import.meta.env.VITE_SERVER_BASE_URL)
  url = new URL(ownerId, url)

  const res = await createFetch(url)

  if (res.status === 400) {
    throw new Error("Invalid owner id")
  }

  return res.json()
}

// HACK: Maybe wrap findPairedTopicsByUserId inside this function?
function excludeUserOwnedTopics(topics: TopicResponse[], ownerId: string): TopicResponse[] {
  return topics.filter((t) => t.owner !== ownerId)
}

async function deleteTopicById(topicId: string): Promise<boolean> {
  let url = new URL("topics/", import.meta.env.VITE_SERVER_BASE_URL)
  url = new URL(topicId, url)

  const res = await createFetch(url, {
    method: "DELETE",
  })

  if (res.status === 404) {
    throw new Error("Topic not found")
  }

  return res.ok
}

export const TopicService = {
  createTopic,
  findTopicsByOwnerId,
  findPairedTopicsByUserId,
  excludeUserOwnedTopics,
  deleteTopicById,
}
