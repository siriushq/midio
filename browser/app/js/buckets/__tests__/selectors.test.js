import { getFilteredBuckets, getCurrentBucket } from "../selectors"

describe("getFilteredBuckets", () => {
  let state
  beforeEach(() => {
    state = {
      buckets: {
        list: ["test1", "test11", "test2"]
      }
    }
  })

  it("should return all buckets if no filter specified", () => {
    state.buckets.filter = ""
    expect(getFilteredBuckets(state)).toEqual(["test1", "test11", "test2"])
  })

  it("should return all matching buckets if filter is specified", () => {
    state.buckets.filter = "test1"
    expect(getFilteredBuckets(state)).toEqual(["test1", "test11"])
  })
})
