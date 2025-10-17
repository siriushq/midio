import React from "react"
import { shallow } from "enzyme"
import { StorageInfo } from "../StorageInfo"

describe("StorageInfo", () => {
  it("should render without crashing", () => {
    shallow(
      <StorageInfo storageInfo={ {used: 60} } fetchStorageInfo={jest.fn()} />
    )
  })

  it("should fetchStorageInfo before component is mounted", () => {
    const fetchStorageInfo = jest.fn()
    shallow(
      <StorageInfo
        storageInfo={ {used: 60} }
        fetchStorageInfo={fetchStorageInfo}
      />
    )
    expect(fetchStorageInfo).toHaveBeenCalled()
  })

  it("should not render anything if used is null", () => {
    const fetchStorageInfo = jest.fn()
    const wrapper = shallow(
      <StorageInfo
      storageInfo={ {used: 0} }
        fetchStorageInfo={fetchStorageInfo}
      />
    )
    expect(wrapper.text()).toBe("")
  })
})
