import JSONrpc from "../jsonrpc"

describe("jsonrpc", () => {
  it("should fail with invalid endpoint", done => {
    try {
      let jsonRPC = new JSONrpc({
        endpoint: "htt://localhost:9000",
        namespace: "Test"
      })
    } catch (e) {
      done()
    }
  })
  it("should succeed with valid endpoint", () => {
    let jsonRPC = new JSONrpc({
      endpoint: "http://localhost:9000/webrpc",
      namespace: "Test"
    })
    expect(jsonRPC.version).toEqual("2.0")
    expect(jsonRPC.host).toEqual("localhost")
    expect(jsonRPC.port).toEqual("9000")
    expect(jsonRPC.path).toEqual("/webrpc")
    expect(jsonRPC.scheme).toEqual("http")
  })
})
