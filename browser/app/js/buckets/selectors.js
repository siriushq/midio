import { createSelector } from "reselect"

const bucketsSelector = state => state.buckets.list
const bucketsFilterSelector = state => state.buckets.filter

export const getFilteredBuckets = createSelector(
  bucketsSelector,
  bucketsFilterSelector,
  (buckets, filter) => buckets.filter(
    bucket => bucket.toLowerCase().indexOf(filter.toLowerCase()) > -1)
)

export const getCurrentBucket = state => state.buckets.currentBucket
