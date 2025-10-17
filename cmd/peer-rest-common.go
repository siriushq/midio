package cmd

const (
	peerRESTVersion       = "v14" // Add GetBucketStats API
	peerRESTVersionPrefix = SlashSeparator + peerRESTVersion
	peerRESTPrefix        = minioReservedBucketPath + "/peer"
	peerRESTPath          = peerRESTPrefix + peerRESTVersionPrefix
)

const (
	peerRESTMethodHealth                 = "/health"
	peerRESTMethodServerInfo             = "/serverinfo"
	peerRESTMethodDriveInfo              = "/driveinfo"
	peerRESTMethodNetInfo                = "/netinfo"
	peerRESTMethodCPUInfo                = "/cpuinfo"
	peerRESTMethodDiskHwInfo             = "/diskhwinfo"
	peerRESTMethodOsInfo                 = "/osinfo"
	peerRESTMethodMemInfo                = "/meminfo"
	peerRESTMethodProcInfo               = "/procinfo"
	peerRESTMethodDispatchNetInfo        = "/dispatchnetinfo"
	peerRESTMethodDeleteBucketMetadata   = "/deletebucketmetadata"
	peerRESTMethodLoadBucketMetadata     = "/loadbucketmetadata"
	peerRESTMethodGetBucketStats         = "/getbucketstats"
	peerRESTMethodServerUpdate           = "/serverupdate"
	peerRESTMethodSignalService          = "/signalservice"
	peerRESTMethodBackgroundHealStatus   = "/backgroundhealstatus"
	peerRESTMethodGetLocks               = "/getlocks"
	peerRESTMethodLoadUser               = "/loaduser"
	peerRESTMethodLoadServiceAccount     = "/loadserviceaccount"
	peerRESTMethodDeleteUser             = "/deleteuser"
	peerRESTMethodDeleteServiceAccount   = "/deleteserviceaccount"
	peerRESTMethodLoadPolicy             = "/loadpolicy"
	peerRESTMethodLoadPolicyMapping      = "/loadpolicymapping"
	peerRESTMethodDeletePolicy           = "/deletepolicy"
	peerRESTMethodLoadGroup              = "/loadgroup"
	peerRESTMethodStartProfiling         = "/startprofiling"
	peerRESTMethodDownloadProfilingData  = "/downloadprofilingdata"
	peerRESTMethodCycleBloom             = "/cyclebloom"
	peerRESTMethodTrace                  = "/trace"
	peerRESTMethodListen                 = "/listen"
	peerRESTMethodLog                    = "/log"
	peerRESTMethodGetLocalDiskIDs        = "/getlocaldiskids"
	peerRESTMethodGetBandwidth           = "/bandwidth"
	peerRESTMethodGetMetacacheListing    = "/getmetacache"
	peerRESTMethodUpdateMetacacheListing = "/updatemetacache"
	peerRESTMethodGetPeerMetrics         = "/peermetrics"
)

const (
	peerRESTBucket         = "bucket"
	peerRESTBuckets        = "buckets"
	peerRESTUser           = "user"
	peerRESTGroup          = "group"
	peerRESTUserTemp       = "user-temp"
	peerRESTPolicy         = "policy"
	peerRESTUserOrGroup    = "user-or-group"
	peerRESTIsGroup        = "is-group"
	peerRESTSignal         = "signal"
	peerRESTProfiler       = "profiler"
	peerRESTTraceErr       = "err"
	peerRESTTraceInternal  = "internal"
	peerRESTTraceStorage   = "storage"
	peerRESTTraceS3        = "s3"
	peerRESTTraceOS        = "os"
	peerRESTTraceThreshold = "threshold"

	peerRESTListenBucket = "bucket"
	peerRESTListenPrefix = "prefix"
	peerRESTListenSuffix = "suffix"
	peerRESTListenEvents = "events"
)
