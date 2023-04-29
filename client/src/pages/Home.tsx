import { A, useNavigate } from "@solidjs/router"
import { Component, Show } from "solid-js"

import { createAuth } from "../store/auth"

const Home: Component = () => {
  const navigate = useNavigate()
  const auth = createAuth

  return (
    <div class="flex min-h-[70vh] flex-col items-center justify-center p-4">
      <div class="flex flex-col gap-2  lg:w-1/2">
        <h1 class="pb-4 text-xl font-semibold">Welcome!</h1>
        <p class="text-red-400">
          This site do stuff, some login stuff?{" "}
          <em class="text-sm text-violet-500">hopefully will do more, a bit insecure app it is</em>
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Veniam architecto facilis, iusto
          asperiores provident ipsa, debitis totam enim quam doloremque omnis? Eum, aliquam ex neque
          nihil blanditiis architecto, commodi officiis eius aspernatur quod quisquam! Lorem ipsum
          dolor sit amet consectetur adipisicing elit. Eligendi tempore sit facilis dolore itaque
          corrupti atque consequatur esse molestias voluptatibus ducimus repellendus, deleniti aut
          excepturi dicta assumenda corporis laborum autem dolor fuga. Boring, maybe, who cares?
        </p>
        <Show when={!auth.isLoggedIn()}>
          <button onClick={() => navigate("/login")} class="btn-md btn-red mt-12 self-start">
            <A href="/login">Login!</A>
          </button>
        </Show>
      </div>
    </div>
  )
}

export default Home
