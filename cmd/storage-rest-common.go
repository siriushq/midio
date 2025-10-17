package cmd

const (
	storageRESTVersion       = "v31" // Added RenameData with fileInfo()
	storageRESTVersionPrefix = SlashSeparator + storageRESTVersion
	storageRESTPrefix        = minioReservedBucketPath + "/storage"
)

const (
	storageRESTMethodHealth      = "/health"
	storageRESTMethodDiskInfo    = "/diskinfo"
	storageRESTMethodNSScanner   = "/nsscanner"
	storageRESTMethodMakeVol     = "/makevol"
	storageRESTMethodMakeVolBulk = "/makevolbulk"
	storageRESTMethodStatVol     = "/statvol"
	storageRESTMethodDeleteVol   = "/deletevol"
	storageRESTMethodListVols    = "/listvols"

	storageRESTMethodAppendFile     = "/appendfile"
	storageRESTMethodCreateFile     = "/createfile"
	storageRESTMethodWriteAll       = "/writeall"
	storageRESTMethodWriteMetadata  = "/writemetadata"
	storageRESTMethodUpdateMetadata = "/updatemetadata"
	storageRESTMethodDeleteVersion  = "/deleteversion"
	storageRESTMethodReadVersion    = "/readversion"
	storageRESTMethodRenameData     = "/renamedata"
	storageRESTMethodCheckParts     = "/checkparts"
	storageRESTMethodCheckFile      = "/checkfile"
	storageRESTMethodReadAll        = "/readall"
	storageRESTMethodReadFile       = "/readfile"
	storageRESTMethodReadFileStream = "/readfilestream"
	storageRESTMethodListDir        = "/listdir"
	storageRESTMethodDeleteFile     = "/deletefile"
	storageRESTMethodDeleteVersions = "/deleteverions"
	storageRESTMethodRenameFile     = "/renamefile"
	storageRESTMethodVerifyFile     = "/verifyfile"
	storageRESTMethodWalkDir        = "/walkdir"
)

const (
	storageRESTVolume         = "volume"
	storageRESTVolumes        = "volumes"
	storageRESTDirPath        = "dir-path"
	storageRESTFilePath       = "file-path"
	storageRESTForceDelMarker = "force-delete-marker"
	storageRESTVersionID      = "version-id"
	storageRESTReadData       = "read-data"
	storageRESTTotalVersions  = "total-versions"
	storageRESTSrcVolume      = "source-volume"
	storageRESTSrcPath        = "source-path"
	storageRESTDstVolume      = "destination-volume"
	storageRESTDstPath        = "destination-path"
	storageRESTOffset         = "offset"
	storageRESTLength         = "length"
	storageRESTCount          = "count"
	storageRESTPrefixFilter   = "prefix"
	storageRESTForwardFilter  = "forward"
	storageRESTRecursive      = "recursive"
	storageRESTReportNotFound = "report-notfound"
	storageRESTBitrotAlgo     = "bitrot-algo"
	storageRESTBitrotHash     = "bitrot-hash"
	storageRESTDiskID         = "disk-id"
	storageRESTForceDelete    = "force-delete"
)
