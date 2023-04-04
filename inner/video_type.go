package inner

type VideoMeta struct {
	ResponseContext   *ResponseContext   `json:"responseContext"`
	PlayabilityStatus *PlayabilityStatus `json:"playabilityStatus"`
	StreamingData     *StreamingData     `json:"streamingData"`
	PlaybackTracking  *PlaybackTracking  `json:"playbackTracking"`
	VideoDetails      *VideoDetails      `json:"videoDetails"`
	PlayerConfig      *PlayerConfig      `json:"playerConfig"`
	Storyboards       *Storyboards       `json:"storyboards"`
	Microformat       *Microformat       `json:"microformat"`
	TrackingParams    string             `json:"trackingParams"`
	Attestation       *Attestation       `json:"attestation"`
	Messages          []*Message         `json:"messages"`
	Endscreen         *Endscreen         `json:"endscreen"`
	FrameworkUpdates  *FrameworkUpdates  `json:"frameworkUpdates"`
}

type Attestation struct {
	PlayerAttestationRenderer *PlayerAttestationRenderer `json:"playerAttestationRenderer"`
}

type PlayerAttestationRenderer struct {
	Challenge    string        `json:"challenge"`
	BotguardData *BotguardData `json:"botguardData"`
}

type BotguardData struct {
	Program            string              `json:"program"`
	InterpreterSafeURL *InterpreterSafeURL `json:"interpreterSafeUrl"`
	ServerEnvironment  int64               `json:"serverEnvironment"`
}

type InterpreterSafeURL struct {
	PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
}

type Endscreen struct {
	EndscreenRenderer *EndscreenRenderer `json:"endscreenRenderer"`
}

type EndscreenRenderer struct {
	Elements       []*Element `json:"elements"`
	StartMS        string     `json:"startMs"`
	TrackingParams string     `json:"trackingParams"`
}

type Element struct {
	EndscreenElementRenderer *EndscreenElementRenderer `json:"endscreenElementRenderer"`
}

type EndscreenElementRenderer struct {
	Style             string              `json:"style"`
	Image             *ImageClass         `json:"image"`
	Left              float64             `json:"left"`
	Width             float64             `json:"width"`
	Top               float64             `json:"top"`
	AspectRatio       float64             `json:"aspectRatio"`
	StartMS           string              `json:"startMs"`
	EndMS             string              `json:"endMs"`
	Title             *Title              `json:"title"`
	Metadata          *Description        `json:"metadata"`
	Endpoint          *Endpoint           `json:"endpoint"`
	TrackingParams    string              `json:"trackingParams"`
	ID                string              `json:"id"`
	ThumbnailOverlays []*ThumbnailOverlay `json:"thumbnailOverlays"`
}

type Endpoint struct {
	ClickTrackingParams string                   `json:"clickTrackingParams"`
	CommandMetadata     *EndpointCommandMetadata `json:"commandMetadata"`
	WatchEndpoint       *WatchEndpoint           `json:"watchEndpoint"`
}

type EndpointCommandMetadata struct {
	WebCommandMetadata *PurpleWebCommandMetadata `json:"webCommandMetadata"`
}

type PurpleWebCommandMetadata struct {
	URL         string `json:"url"`
	WebPageType string `json:"webPageType"`
	RootVe      int64  `json:"rootVe"`
}

type WatchEndpoint struct {
	VideoID                            string                              `json:"videoId"`
	WatchEndpointSupportedOnesieConfig *WatchEndpointSupportedOnesieConfig `json:"watchEndpointSupportedOnesieConfig"`
}

type WatchEndpointSupportedOnesieConfig struct {
	HTML5PlaybackOnesieConfig *HTML5PlaybackOnesieConfig `json:"html5PlaybackOnesieConfig"`
}

type HTML5PlaybackOnesieConfig struct {
	CommonConfig *CommonConfig `json:"commonConfig"`
}

type CommonConfig struct {
	URL string `json:"url"`
}

type ImageClass struct {
	Thumbnails []*ThumbnailElement `json:"thumbnails"`
}

type ThumbnailElement struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Description struct {
	SimpleText string `json:"simpleText"`
}

type ThumbnailOverlay struct {
	ThumbnailOverlayTimeStatusRenderer *ThumbnailOverlayTimeStatusRenderer `json:"thumbnailOverlayTimeStatusRenderer"`
}

type ThumbnailOverlayTimeStatusRenderer struct {
	Text  *Title `json:"text"`
	Style string `json:"style"`
}

type Title struct {
	Accessibility *Accessibility `json:"accessibility"`
	SimpleText    string         `json:"simpleText"`
}

type Accessibility struct {
	AccessibilityData *AccessibilityData `json:"accessibilityData"`
}

type AccessibilityData struct {
	Label string `json:"label"`
}

type FrameworkUpdates struct {
	EntityBatchUpdate *EntityBatchUpdate `json:"entityBatchUpdate"`
}

type EntityBatchUpdate struct {
	Mutations []*Mutation `json:"mutations"`
	Timestamp *Timestamp  `json:"timestamp"`
}

type Mutation struct {
	EntityKey string   `json:"entityKey"`
	Type      string   `json:"type"`
	Payload   *Payload `json:"payload"`
}

type Payload struct {
	OfflineabilityEntity *OfflineabilityEntity `json:"offlineabilityEntity"`
}

type OfflineabilityEntity struct {
	Key                     string `json:"key"`
	AddToOfflineButtonState string `json:"addToOfflineButtonState"`
}

type Timestamp struct {
	Seconds string `json:"seconds"`
	Nanos   int64  `json:"nanos"`
}

type Message struct {
	TooltipRenderer *TooltipRenderer `json:"tooltipRenderer"`
}

type TooltipRenderer struct {
	PromoConfig       *PromoConfig     `json:"promoConfig"`
	TargetID          string           `json:"targetId"`
	Text              *Text            `json:"text"`
	DetailsText       *Text            `json:"detailsText"`
	DismissButton     *DismissButton   `json:"dismissButton"`
	SuggestedPosition *DismissStrategy `json:"suggestedPosition"`
	DismissStrategy   *DismissStrategy `json:"dismissStrategy"`
	DwellTimeMS       string           `json:"dwellTimeMs"`
	TrackingParams    string           `json:"trackingParams"`
}

type Text struct {
	Runs []*Run `json:"runs"`
}

type Run struct {
	Text string `json:"text"`
}

type DismissButton struct {
	ButtonRenderer *ButtonRenderer `json:"buttonRenderer"`
}

type ButtonRenderer struct {
	Size           string   `json:"size"`
	Text           *Text    `json:"text"`
	TrackingParams string   `json:"trackingParams"`
	Command        *Command `json:"command"`
}

type Command struct {
	ClickTrackingParams    string                  `json:"clickTrackingParams"`
	CommandExecutorCommand *CommandExecutorCommand `json:"commandExecutorCommand"`
}

type CommandExecutorCommand struct {
	Commands []*AcceptCommand `json:"commands"`
}

type AcceptCommand struct {
	ClickTrackingParams string                        `json:"clickTrackingParams"`
	CommandMetadata     *AcceptCommandCommandMetadata `json:"commandMetadata"`
	FeedbackEndpoint    *FeedbackEndpoint             `json:"feedbackEndpoint"`
}

type AcceptCommandCommandMetadata struct {
	WebCommandMetadata *FluffyWebCommandMetadata `json:"webCommandMetadata"`
}

type FluffyWebCommandMetadata struct {
	SendPost bool   `json:"sendPost"`
	APIURL   string `json:"apiUrl"`
}

type FeedbackEndpoint struct {
	FeedbackToken string     `json:"feedbackToken"`
	UIActions     *UIActions `json:"uiActions"`
}

type UIActions struct {
	HideEnclosingContainer bool `json:"hideEnclosingContainer"`
}

type DismissStrategy struct {
	Type string `json:"type"`
}

type PromoConfig struct {
	PromoID             string           `json:"promoId"`
	ImpressionEndpoints []*AcceptCommand `json:"impressionEndpoints"`
	AcceptCommand       *AcceptCommand   `json:"acceptCommand"`
	DismissCommand      *AcceptCommand   `json:"dismissCommand"`
}

type Microformat struct {
	PlayerMicroformatRenderer *PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
	Thumbnail          *ImageClass  `json:"thumbnail"`
	Embed              *Embed       `json:"embed"`
	Title              *Description `json:"title"`
	Description        *Description `json:"description"`
	LengthSeconds      string       `json:"lengthSeconds"`
	OwnerProfileURL    string       `json:"ownerProfileUrl"`
	ExternalChannelID  string       `json:"externalChannelId"`
	IsFamilySafe       bool         `json:"isFamilySafe"`
	AvailableCountries []string     `json:"availableCountries"`
	IsUnlisted         bool         `json:"isUnlisted"`
	HasYpcMetadata     bool         `json:"hasYpcMetadata"`
	ViewCount          string       `json:"viewCount"`
	Category           string       `json:"category"`
	PublishDate        string       `json:"publishDate"`
	OwnerChannelName   string       `json:"ownerChannelName"`
	UploadDate         string       `json:"uploadDate"`
}

type Embed struct {
	IframeURL string `json:"iframeUrl"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
}

type PlayabilityStatus struct {
	Status          string      `json:"status"`
	PlayableInEmbed bool        `json:"playableInEmbed"`
	Miniplayer      *Miniplayer `json:"miniplayer"`
	ContextParams   string      `json:"contextParams"`
}

type Miniplayer struct {
	MiniplayerRenderer *MiniplayerRenderer `json:"miniplayerRenderer"`
}

type MiniplayerRenderer struct {
	PlaybackMode string `json:"playbackMode"`
}

type PlaybackTracking struct {
	VideostatsPlaybackURL                   *URL    `json:"videostatsPlaybackUrl"`
	VideostatsDelayplayURL                  *URL    `json:"videostatsDelayplayUrl"`
	VideostatsWatchtimeURL                  *URL    `json:"videostatsWatchtimeUrl"`
	PtrackingURL                            *URL    `json:"ptrackingUrl"`
	QoeURL                                  *URL    `json:"qoeUrl"`
	AtrURL                                  *URL    `json:"atrUrl"`
	VideostatsScheduledFlushWalltimeSeconds []int64 `json:"videostatsScheduledFlushWalltimeSeconds"`
	VideostatsDefaultFlushIntervalSeconds   int64   `json:"videostatsDefaultFlushIntervalSeconds"`
}

type URL struct {
	BaseURL                 string `json:"baseUrl"`
	ElapsedMediaTimeSeconds int64  `json:"elapsedMediaTimeSeconds"`
}

type PlayerConfig struct {
	AudioConfig           *AudioConfig           `json:"audioConfig"`
	StreamSelectionConfig *StreamSelectionConfig `json:"streamSelectionConfig"`
	MediaCommonConfig     *MediaCommonConfig     `json:"mediaCommonConfig"`
	WebPlayerConfig       *WebPlayerConfig       `json:"webPlayerConfig"`
}

type AudioConfig struct {
	LoudnessDB              float64 `json:"loudnessDb"`
	PerceptualLoudnessDB    float64 `json:"perceptualLoudnessDb"`
	EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
}

type MediaCommonConfig struct {
	DynamicReadaheadConfig *DynamicReadaheadConfig `json:"dynamicReadaheadConfig"`
}

type DynamicReadaheadConfig struct {
	MaxReadAheadMediaTimeMS int64 `json:"maxReadAheadMediaTimeMs"`
	MinReadAheadMediaTimeMS int64 `json:"minReadAheadMediaTimeMs"`
	ReadAheadGrowthRateMS   int64 `json:"readAheadGrowthRateMs"`
}

type StreamSelectionConfig struct {
	MaxBitrate string `json:"maxBitrate"`
}

type WebPlayerConfig struct {
	UseCobaltTvosDash       bool                     `json:"useCobaltTvosDash"`
	WebPlayerActionsPorting *WebPlayerActionsPorting `json:"webPlayerActionsPorting"`
}

type WebPlayerActionsPorting struct {
	GetSharePanelCommand        *GetSharePanelCommand        `json:"getSharePanelCommand"`
	SubscribeCommand            *SubscribeCommand            `json:"subscribeCommand"`
	UnsubscribeCommand          *UnsubscribeCommand          `json:"unsubscribeCommand"`
	AddToWatchLaterCommand      *AddToWatchLaterCommand      `json:"addToWatchLaterCommand"`
	RemoveFromWatchLaterCommand *RemoveFromWatchLaterCommand `json:"removeFromWatchLaterCommand"`
}

type AddToWatchLaterCommand struct {
	ClickTrackingParams  string                                      `json:"clickTrackingParams"`
	CommandMetadata      *AcceptCommandCommandMetadata               `json:"commandMetadata"`
	PlaylistEditEndpoint *AddToWatchLaterCommandPlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type AddToWatchLaterCommandPlaylistEditEndpoint struct {
	PlaylistID string          `json:"playlistId"`
	Actions    []*PurpleAction `json:"actions"`
}

type PurpleAction struct {
	AddedVideoID string `json:"addedVideoId"`
	Action       string `json:"action"`
}

type GetSharePanelCommand struct {
	ClickTrackingParams                 string                               `json:"clickTrackingParams"`
	CommandMetadata                     *AcceptCommandCommandMetadata        `json:"commandMetadata"`
	WebPlayerShareEntityServiceEndpoint *WebPlayerShareEntityServiceEndpoint `json:"webPlayerShareEntityServiceEndpoint"`
}

type WebPlayerShareEntityServiceEndpoint struct {
	SerializedShareEntity string `json:"serializedShareEntity"`
}

type RemoveFromWatchLaterCommand struct {
	ClickTrackingParams  string                                           `json:"clickTrackingParams"`
	CommandMetadata      *AcceptCommandCommandMetadata                    `json:"commandMetadata"`
	PlaylistEditEndpoint *RemoveFromWatchLaterCommandPlaylistEditEndpoint `json:"playlistEditEndpoint"`
}

type RemoveFromWatchLaterCommandPlaylistEditEndpoint struct {
	PlaylistID string          `json:"playlistId"`
	Actions    []*FluffyAction `json:"actions"`
}

type FluffyAction struct {
	Action         string `json:"action"`
	RemovedVideoID string `json:"removedVideoId"`
}

type SubscribeCommand struct {
	ClickTrackingParams string                        `json:"clickTrackingParams"`
	CommandMetadata     *AcceptCommandCommandMetadata `json:"commandMetadata"`
	SubscribeEndpoint   *SubscribeEndpoint            `json:"subscribeEndpoint"`
}

type SubscribeEndpoint struct {
	ChannelIDS []string `json:"channelIds"`
	Params     string   `json:"params"`
}

type UnsubscribeCommand struct {
	ClickTrackingParams string                        `json:"clickTrackingParams"`
	CommandMetadata     *AcceptCommandCommandMetadata `json:"commandMetadata"`
	UnsubscribeEndpoint *SubscribeEndpoint            `json:"unsubscribeEndpoint"`
}

type ResponseContext struct {
	VisitorData                     string                           `json:"visitorData"`
	ServiceTrackingParams           []*ServiceTrackingParam          `json:"serviceTrackingParams"`
	MainAppWebResponseContext       *MainAppWebResponseContext       `json:"mainAppWebResponseContext"`
	WebResponseContextExtensionData *WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type MainAppWebResponseContext struct {
	LoggedOut bool `json:"loggedOut"`
}

type ServiceTrackingParam struct {
	Service string   `json:"service"`
	Params  []*Param `json:"params"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type WebResponseContextExtensionData struct {
	HasDecorated bool `json:"hasDecorated"`
}

type Storyboards struct {
	PlayerStoryboardSpecRenderer *PlayerStoryboardSpecRenderer `json:"playerStoryboardSpecRenderer"`
}

type PlayerStoryboardSpecRenderer struct {
	Spec string `json:"spec"`
}

type StreamingData struct {
	ExpiresInSeconds string    `json:"expiresInSeconds"`
	Formats          []*Format `json:"formats"`
	AdaptiveFormats  []*Format `json:"adaptiveFormats"`
	ProbeURL         string    `json:"probeUrl"`
}

type Format struct {
	Itag             int64      `json:"itag"`
	URL              string     `json:"url"`
	MIMEType         string     `json:"mimeType"`
	Bitrate          int64      `json:"bitrate"`
	Width            int64      `json:"width,omitempty"`
	Height           int64      `json:"height,omitempty"`
	InitRange        *Range     `json:"initRange,omitempty"`
	IndexRange       *Range     `json:"indexRange,omitempty"`
	LastModified     string     `json:"lastModified"`
	ContentLength    string     `json:"contentLength,omitempty"`
	Quality          string     `json:"quality"`
	FPS              int64      `json:"fps,omitempty"`
	QualityLabel     string     `json:"qualityLabel,omitempty"`
	ProjectionType   string     `json:"projectionType"`
	AverageBitrate   int64      `json:"averageBitrate,omitempty"`
	ApproxDurationMS string     `json:"approxDurationMs"`
	ColorInfo        *ColorInfo `json:"colorInfo,omitempty"`
	HighReplication  bool       `json:"highReplication,omitempty"`
	AudioQuality     string     `json:"audioQuality,omitempty"`
	AudioSampleRate  string     `json:"audioSampleRate,omitempty"`
	AudioChannels    int64      `json:"audioChannels,omitempty"`
	LoudnessDB       float64    `json:"loudnessDb,omitempty"`
}

type ColorInfo struct {
	Primaries               string `json:"primaries"`
	TransferCharacteristics string `json:"transferCharacteristics"`
	MatrixCoefficients      string `json:"matrixCoefficients"`
}

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type VideoDetails struct {
	VideoID           string      `json:"videoId"`
	Title             string      `json:"title"`
	LengthSeconds     string      `json:"lengthSeconds"`
	Keywords          []string    `json:"keywords"`
	ChannelID         string      `json:"channelId"`
	IsOwnerViewing    bool        `json:"isOwnerViewing"`
	ShortDescription  string      `json:"shortDescription"`
	IsCrawlable       bool        `json:"isCrawlable"`
	Thumbnail         *ImageClass `json:"thumbnail"`
	AllowRatings      bool        `json:"allowRatings"`
	ViewCount         string      `json:"viewCount"`
	Author            string      `json:"author"`
	IsPrivate         bool        `json:"isPrivate"`
	IsUnpluggedCorpus bool        `json:"isUnpluggedCorpus"`
	IsLiveContent     bool        `json:"isLiveContent"`
}
