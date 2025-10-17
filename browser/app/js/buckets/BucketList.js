import React from "react"
import { connect } from "react-redux"
import { Scrollbars } from "react-custom-scrollbars"
import InfiniteScroll from "react-infinite-scroller"
import * as actionsBuckets from "./actions"
import { getFilteredBuckets } from "./selectors"
import BucketContainer from "./BucketContainer"
import web from "../web"
import history from "../history"
import { pathSlice } from "../utils"

export class BucketList extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      page: 1
    }
    this.loadNextPage = this.loadNextPage.bind(this)
  }
  componentDidUpdate(prevProps) {
    if (this.props.filter !== prevProps.filter) {
      this.setState({
        page: 1
      })
    }
  }
  componentWillMount() {
    const { fetchBuckets, setBucketList, selectBucket } = this.props
    if (web.LoggedIn()) {
      fetchBuckets()
    } else {
      const { bucket, prefix } = pathSlice(history.location.pathname)
      if (bucket) {
        setBucketList([bucket])
        selectBucket(bucket, prefix)
      } else {
        history.replace("/login")
      }
    }
  }
  loadNextPage() {
    this.setState({
      page: this.state.page + 1
    })
  }
  render() {
    const { filteredBuckets } = this.props
    const visibleBuckets = filteredBuckets.slice(0, this.state.page * 100)
    return (
      <div className="fesl-inner">
        <Scrollbars
          renderTrackVertical={props => <div className="scrollbar-vertical" />}
        >
          <InfiniteScroll
            pageStart={0}
            loadMore={this.loadNextPage}
            hasMore={filteredBuckets.length > visibleBuckets.length}
            useWindow={false}
            element="div"
            initialLoad={false}
          >
            <ul>
              {visibleBuckets.map(bucket => (
                <BucketContainer key={bucket} bucket={bucket} />
              ))}
            </ul>
          </InfiniteScroll>
        </Scrollbars>
      </div>
    )
  }
}

const mapStateToProps = state => {
  return {
    filteredBuckets: getFilteredBuckets(state),
    filter: state.buckets.filter
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchBuckets: () => dispatch(actionsBuckets.fetchBuckets()),
    setBucketList: buckets => dispatch(actionsBuckets.setList(buckets)),
    selectBucket: (bucket, prefix) =>
      dispatch(actionsBuckets.selectBucket(bucket, prefix))
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BucketList)
