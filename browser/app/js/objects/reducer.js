import * as actionsObjects from "./actions"
import { SORT_ORDER_ASC } from "../constants"

const removeObject = (list, objectToRemove, lookup) => {
  const idx = list.findIndex(object => lookup(object) === objectToRemove)
  if (idx == -1) {
    return list
  }
  return [...list.slice(0, idx), ...list.slice(idx + 1)]
}

export default (
  state = {
    list: [],
    filter: "",
    listLoading: false,
    sortBy: "",
    sortOrder: SORT_ORDER_ASC,
    currentPrefix: "",
    prefixWritable: false,
    shareObject: {
      show: false,
      object: "",
      url: ""
    },
    checkedList: []
  },
  action
) => {
  switch (action.type) {
    case actionsObjects.SET_LIST:
      return {
        ...state,
        list: action.objects
      }
    case actionsObjects.RESET_LIST:
      return {
        ...state,
        list: []
      }
    case actionsObjects.SET_FILTER:
      return {
        ...state,
        filter: action.filter
      }
    case actionsObjects.SET_LIST_LOADING:
      return {
        ...state,
        listLoading: action.listLoading
      }
    case actionsObjects.REMOVE:
      return {
        ...state,
        list: removeObject(state.list, action.object, object => object.name)
      }
    case actionsObjects.SET_SORT_BY:
      return {
        ...state,
        sortBy: action.sortBy
      }
    case actionsObjects.SET_SORT_ORDER:
      return {
        ...state,
        sortOrder: action.sortOrder
      }
    case actionsObjects.SET_CURRENT_PREFIX:
      return {
        ...state,
        currentPrefix: action.prefix
      }
    case actionsObjects.SET_PREFIX_WRITABLE:
      return {
        ...state,
        prefixWritable: action.prefixWritable
      }
    case actionsObjects.SET_SHARE_OBJECT:
      return {
        ...state,
        shareObject: {
          show: action.show,
          object: action.object,
          url: action.url,
          showExpiryDate: action.showExpiryDate
        }
      }
    case actionsObjects.CHECKED_LIST_ADD:
      return {
        ...state,
        checkedList: [...state.checkedList, action.object]
      }
    case actionsObjects.CHECKED_LIST_REMOVE:
      return {
        ...state,
        checkedList: removeObject(
          state.checkedList,
          action.object,
          object => object
        )
      }
    case actionsObjects.CHECKED_LIST_RESET:
      return {
        ...state,
        checkedList: []
      }
    default:
      return state
  }
}
