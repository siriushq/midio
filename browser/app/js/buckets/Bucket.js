import React from "react"
import classNames from "classnames"
import BucketDropdown from "./BucketDropdown"

export const Bucket = ({ bucket, isActive, selectBucket }) => {
  return (
    <li
      className={classNames({
        active: isActive
      })}
      onClick={e => {
        e.preventDefault()
        selectBucket(bucket)
      }}
    >
      <a
        href=""
        className={classNames({
          "fesli-loading": false
        })}
      >
        {bucket}
      </a>
      <BucketDropdown bucket={bucket}/>
    </li>
  )
}

export default Bucket
