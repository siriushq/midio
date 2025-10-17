import React from "react"

let InputGroup = ({
  label,
  id,
  name,
  value,
  onChange,
  type,
  spellCheck,
  required,
  readonly,
  autoComplete,
  align,
  className
}) => {
  var input = (
    <input
      id={id}
      name={name}
      value={value}
      onChange={onChange}
      className="ig-text"
      type={type}
      spellCheck={spellCheck}
      required={required}
      autoComplete={autoComplete}
    />
  )
  if (readonly)
    input = (
      <input
        id={id}
        name={name}
        value={value}
        onChange={onChange}
        className="ig-text"
        type={type}
        spellCheck={spellCheck}
        required={required}
        autoComplete={autoComplete}
        disabled
      />
    )
  return (
    <div className={"input-group " + align + " " + className}>
      {input}
      <i className="ig-helpers" />
      <label className="ig-label">{label}</label>
    </div>
  )
}

export default InputGroup
