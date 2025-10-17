import React from "react"
import { shallow } from "enzyme"
import history from "../../history"
import { BucketList } from "../BucketList"

jest.mock("../../web", () => ({
  LoggedIn: jest
    .fn(() => false)
    .mockReturnValueOnce(true)
    .mockReturnValueOnce(true)
}))

describe("BucketList", () => {
  it("should render without crashing", () => {
    const fetchBuckets = jest.fn()
    shallow(<BucketList filteredBuckets={[]} fetchBuckets={fetchBuckets} />)
  })

  it("should call fetchBuckets before component is mounted", () => {
    const fetchBuckets = jest.fn()
    const wrapper = shallow(
      <BucketList filteredBuckets={[]} fetchBuckets={fetchBuckets} />
    )
    expect(fetchBuckets).toHaveBeenCalled()
  })

  it("should call setBucketList and selectBucket before component is mounted when the user has not loggedIn", () => {
    const setBucketList = jest.fn()
    const selectBucket = jest.fn()
    history.push("/bk1/pre1")
    const wrapper = shallow(
      <BucketList
        filteredBuckets={[]}
        setBucketList={setBucketList}
        selectBucket={selectBucket}
      />
    )
    expect(setBucketList).toHaveBeenCalledWith(["bk1"])
    expect(selectBucket).toHaveBeenCalledWith("bk1", "pre1")
  })
})
