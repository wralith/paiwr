import { A } from "@solidjs/router"
import { Component, createSignal } from "solid-js"

const Home: Component = () => {
  const [c, setC] = createSignal(0)

  return (
    <div class="flex min-h-[70vh] flex-col items-center justify-center p-12">
      <h1 class="text-xl font-semibold">Hello Solid</h1>
      <p class="p-4 text-4xl font-bold">{c()}</p>
      <div class="flex gap-2">
        <button class="btn btn-violet" onClick={() => setC(c() + 1)}>
          Increase
        </button>
        <button class="btn btn-red" onClick={() => setC(c() - 1)}>
          Decrease
        </button>
      </div>
      <button class="btn-lg btn-red mt-24">
        <A href="/login">Login!</A>
      </button>
    </div>
  )
}

export default Home
