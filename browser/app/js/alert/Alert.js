import React from "react"
import AlertComponent from "react-bootstrap/lib/Alert"

const Alert = ({ show, type, message, onDismiss }) => (
  <AlertComponent
    className={"alert animated " + (show ? "fadeInDown" : "fadeOutUp")}
    bsStyle={type}
    onDismiss={onDismiss}
  >
    <div className="text-center">{message}</div>
  </AlertComponent>
)

export default Alert
