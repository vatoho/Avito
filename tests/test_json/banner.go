package test_json

const (
	BannerAddOld = `{
		"tag_ids": [123, 12, 1],
		"feature_id": 1,
		"content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
		"is_active": false}`
	BannerAddNew = `{
		"tag_ids": [123, 12, 4],
		"feature_id": 7,
		"content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}",
		"is_active": true}`

	ExpectedAllBanners = `[
  {
    "content": "{\"title\": \"some_title1\", \"text\": \"some_text1\", \"url\": \"some_url1\"}",
    "created_at": "0001-01-01T01:00:00Z",
    "feature_id": 1,
    "id": 1,
    "is_active": true,
    "tag_ids": [
      1,
      2
    ],
    "updated_at": "0001-01-01T01:00:00Z"
  },
  {
    "content": "{\"title\": \"some_title2\", \"text\": \"some_text2\", \"url\": \"some_url2\"}",
    "created_at": "0001-01-01T02:00:00Z",
    "feature_id": 2,
    "id": 2,
    "is_active": false,
    "tag_ids": [
      1,
      4
    ],
    "updated_at": "0001-01-01T02:00:00Z"
  }
]
`

	ExpectedBanner1 = "{\"content\":\"{\\\"title\\\": \\\"some_title1\\\", \\\"text\\\": \\\"some_text1\\\", \\\"url\\\": \\\"some_url1\\\"}\"}"

	ExpectedBannersByFeature1 = `[
  {
    "content": "{\"title\": \"some_title1\", \"text\": \"some_text1\", \"url\": \"some_url1\"}",
    "created_at": "0001-01-01T01:00:00Z",
    "feature_id": 1,
    "id": 1,
    "is_active": true,
    "tag_ids": [
      1,
      2
    ],
    "updated_at": "0001-01-01T01:00:00Z"
  }]
`

	BannerUpdateDuplicate = `{
		"feature_id": 2,
		"content": "{\"title\": \"some_title\", \"text\": \"some_text\", \"url\": \"some_url\"}"
		}`

	BannerUpdateOk = `{
		"feature_id": 4,
		"content": "{\"title\": \"new_title\"}"
		}`

	BannerVersions = `[
  {
    "banner_id": 1,
    "content": "{\"title\": \"some_title3\", \"text\": \"some_text3\", \"url\": \"some_url3\"}",
    "is_active": true,
    "updated_at": "0001-01-01T01:00:00Z",
    "version_id": 1
  },
  {
    "banner_id": 1,
    "content": "{\"title\": \"some_title4\", \"text\": \"some_text4\", \"url\": \"some_url4\"}",
    "is_active": true,
    "updated_at": "0001-01-01T01:00:00Z",
    "version_id": 2
  }
]`
)
