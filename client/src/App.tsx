import { Route, Routes } from "@solidjs/router"
import { type Component, lazy } from "solid-js"

import { OnlyGuests, OnlyUsers } from "./components/Guard"
import { Navbar } from "./components/Navbar"
import { createLocalStorageSignal } from "./signals/createLocalStorageSignal"

const Home = lazy(() => import("./pages/Home"))
const Login = lazy(() => import("./pages/Login"))
const Register = lazy(() => import("./pages/Register"))
const MyTopics = lazy(() => import("./pages/MyTopics"))

const App: Component = () => {
  const [dark, setDark] = createLocalStorageSignal(false, "theme")
  const toggleDark = () => setDark(!dark())

  return (
    <div class={`${dark() ? "dark" : ""}`}>
      <div class="min-h-screen transition-colors dark:bg-gray-700 dark:text-white">
        <Navbar toggleDark={toggleDark} dark={dark()} />
        <Routes>
          <Route path="/" component={Home} />
          <OnlyGuests>
            <Route path="/login" component={Login} />
            <Route path="/register" component={Register} />
          </OnlyGuests>
          <OnlyUsers>
            <Route path="/my-topics" component={MyTopics} />
          </OnlyUsers>
        </Routes>
      </div>
    </div>
  )
}

export default App
