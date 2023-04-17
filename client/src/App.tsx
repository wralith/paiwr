import { Route, Routes } from "@solidjs/router"
import { type Component, createSignal } from "solid-js"

import { Navbar } from "./components/Navbar"
import { Home } from "./pages/Home"
import { Login } from "./pages/Login"
import { Logout } from "./pages/Logout"
import { Register } from "./pages/Register"

// TODO: Restricted paths according to auth state

const App: Component = () => {
  const [dark, setDark] = createSignal(false)
  const toggleDark = () => setDark(!dark())

  return (
    <div class={`${dark() ? "dark " : ""}`}>
      <div class="min-h-screen dark:bg-gray-700 dark:text-white">
        <Navbar toggleDark={toggleDark} />
        <Routes>
          <Route path="/" component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/logout" component={Logout} />
          <Route path="/register" component={Register} />
        </Routes>
      </div>
    </div>
  )
}

export default App
