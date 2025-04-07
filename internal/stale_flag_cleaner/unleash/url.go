package unleash

const (
	featureURL          = "%s/projects/%s/features/%s"
	addFeatureTagURL    = "%s/api/admin/features/%s/tags"
	getStaleFeaturesURL = "%s/api/admin/search/features?state=IS:stale&project=IS:%s"
)
