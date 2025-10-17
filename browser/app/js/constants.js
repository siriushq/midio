// File for all the browser constants.

// minioBrowserPrefix absolute path.
var p = window.location.pathname
export const minioBrowserPrefix = p.slice(0, p.indexOf("/", 1))

export const READ_ONLY = "readonly"
export const WRITE_ONLY = "writeonly"
export const READ_WRITE = "readwrite"
export const NONE = "none"

export const SHARE_OBJECT_EXPIRY_DAYS = 5
export const SHARE_OBJECT_EXPIRY_HOURS = 0
export const SHARE_OBJECT_EXPIRY_MINUTES = 0

export const ACCESS_KEY_MIN_LENGTH = 3
export const SECRET_KEY_MIN_LENGTH = 8

export const SORT_BY_NAME = "name"
export const SORT_BY_SIZE = "size"
export const SORT_BY_LAST_MODIFIED = "last-modified"

export const SORT_ORDER_ASC = "asc"
export const SORT_ORDER_DESC = "desc"
