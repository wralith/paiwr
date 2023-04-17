import { Component } from "solid-js"
import { JSX } from "solid-js/jsx-runtime"

type Props = {
  value: string
  onInput: JSX.EventHandler<HTMLInputElement, InputEvent>
} & JSX.InputHTMLAttributes<HTMLInputElement>

export const Input: Component<Props> = (props) => {
  return (
    <input
      type={props.type}
      placeholder={props.placeholder ?? props.name}
      value={props.value}
      onInput={(e) => props.onInput(e)}
    />
  )
}
