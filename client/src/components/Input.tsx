import { JSX, splitProps } from "solid-js"

type Props = {
  name: string
  type: "text" | "email" | "tel" | "password" | "url" | "date"
  label?: string
  placeholder?: string
  value: string | undefined
  error: string
  required?: boolean
  ref: (element: HTMLInputElement) => void
  onInput: JSX.EventHandler<HTMLInputElement, InputEvent>
  onChange: JSX.EventHandler<HTMLInputElement, Event>
  onBlur: JSX.EventHandler<HTMLInputElement, FocusEvent>
}

export function Input(props: Props) {
  const [, inputProps] = splitProps(props, ["value", "label", "error"])
  return (
    <div class="flex flex-col gap-1">
      {props.label && (
        <label for={props.name}>
          {props.label} {props.required && <span>*</span>}
        </label>
      )}
      <input
        {...inputProps}
        id={props.name}
        value={props.value || ""}
        aria-invalid={!!props.error}
        aria-errormessage={`${props.name}-error`}
      />
      {props.error && (
        <div
          class="error-message mt-1 border border-gray-200 p-2 text-sm text-red-400 dark:border-gray-600"
          id={`${props.name}-error`}
        >
          {props.error}
        </div>
      )}
    </div>
  )
}
