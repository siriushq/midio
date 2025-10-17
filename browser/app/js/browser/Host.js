import React from "react"

export const Host = () => (
  <div className="fes-host">
    <i className="fas fa-globe-americas" />
    <a href="/">{window.location.host}</a>
  </div>
)

export default Host
