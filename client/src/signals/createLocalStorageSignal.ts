import { createSignal, Signal } from "solid-js"

export function createLocalStorageSignal<T>(defaultValue: T, key: string): Signal<T> {
  const initialValue = localStorage.getItem(key)
    ? JSON.parse(localStorage.getItem(key)!)
    : defaultValue

  const [value, setValue] = createSignal<T>(initialValue)

  const setter = ((newValue): T => {
    if (!newValue) {
      localStorage.removeItem(key)
    } else {
      localStorage.setItem(key, JSON.stringify(newValue))
    }
    return setValue(newValue)
  }) as typeof setValue

  return [value, setter]
}
