import * as uploadsActions from "./actions"

const add = (files, action) => ({
  ...files,
  [action.slug]: {
    loaded: 0,
    size: action.size,
    name: action.name
  }
})

const updateProgress = (files, action) => ({
  ...files,
  [action.slug]: {
    ...files[action.slug],
    loaded: action.loaded
  }
})

const stop = (files, action) => {
  const newFiles = Object.assign({}, files)
  delete newFiles[action.slug]
  return newFiles
}

export default (state = { files: {}, showAbortModal: false }, action) => {
  switch (action.type) {
    case uploadsActions.ADD:
      return {
        ...state,
        files: add(state.files, action)
      }
    case uploadsActions.UPDATE_PROGRESS:
      return {
        ...state,
        files: updateProgress(state.files, action)
      }
    case uploadsActions.STOP:
      return {
        ...state,
        files: stop(state.files, action)
      }
    case uploadsActions.SHOW_ABORT_MODAL:
      return {
        ...state,
        showAbortModal: action.show
      }
    default:
      return state
  }
}
