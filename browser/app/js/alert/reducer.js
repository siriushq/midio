import * as actionsAlert from "./actions"

const initialState = {
  show: false,
  type: "danger"
}
export default (state = initialState, action) => {
  switch (action.type) {
    case actionsAlert.SET:
      return {
        show: true,
        id: action.alert.id,
        type: action.alert.type,
        message: action.alert.message
      }
    case actionsAlert.CLEAR:
      if (action.alert && action.alert.id != state.id) {
        return state
      } else {
        return initialState
      }
    default:
      return state
  }
}
