export const SET = "alert/SET"
export const CLEAR = "alert/CLEAR"

export let alertId = 0

export const set = alert => {
  const id = alertId++
  return (dispatch, getState) => {
    if (alert.type !== "danger" || alert.autoClear) {
      setTimeout(() => {
        dispatch({
          type: CLEAR,
          alert: {
            id
          }
        })
      }, 5000)
    }
    dispatch({
      type: SET,
      alert: Object.assign({}, alert, {
        id
      })
    })
  }
}

export const clear = () => {
  return { type: CLEAR }
}
