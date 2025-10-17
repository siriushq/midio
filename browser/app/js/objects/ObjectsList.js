import React from "react"
import ObjectContainer from "./ObjectContainer"
import PrefixContainer from "./PrefixContainer"

export const ObjectsList = ({ objects }) => {
  const list = objects.map(object => {
    if (object.name.endsWith("/")) {
      return <PrefixContainer object={object} key={object.name} />
    } else {
      return <ObjectContainer object={object} key={object.name} />
    }
  })
  return <div>{list}</div>
}

export default ObjectsList
