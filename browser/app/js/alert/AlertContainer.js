import React from "react"
import { connect } from "react-redux"
import Alert from "./Alert"
import * as alertActions from "./actions"

export const AlertContainer = ({ alert, clearAlert }) => {
  if (!alert.message) {
    return ""
  }
  return <Alert {...alert} onDismiss={clearAlert} />
}

const mapStateToProps = state => {
  return {
    alert: state.alert
  }
}

const mapDispatchToProps = dispatch => {
  return {
    clearAlert: () => dispatch(alertActions.clear())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(AlertContainer)
