package goxford

//FaceDetectResponse is the structured response from a call to Detect.
type FaceDetectResponse struct {
	//TODO: Maybe change this to UUID?
	FaceID         string         `json:"faceId,omitempty"`
	FaceRectangle  FaceRectangle  `json:"faceRectangle,omitempty"`
	FaceLandmarks  FaceLandmarks  `json:"faceLandmarks,omitempty"`
	FaceAttributes FaceAttributes `json:"faceAttributes,omitempty"`
}

//FaceDetectRequest is the structured request to make a call to Detect.
type FaceDetectRequest struct {
	URL string `json:"url,omitempty"`
}

//FaceRectangle represents the face rectangle detected on an image.
type FaceRectangle struct {
	Width  int32 `json:"width,omitempty"`
	Height int32 `json:"height,omitempty"`
	Left   int32 `json:"left,omitempty"`
	Top    int32 `json:"top,omitempty"`
}

//FaceLandmarks represents the particular landmarks on a detected face.
type FaceLandmarks struct {
	PupilLeft           Point `json:"pupilLeft,omitempty"`
	PupilRight          Point `json:"pupilRight,omitempty"`
	NoseTip             Point `json:"noseTip,omitempty"`
	MouthLeft           Point `json:"mouthLeft,omitempty"`
	MouthRight          Point `json:"mouthRight,omitempty"`
	EyebrowLeftOuter    Point `json:"eyebrowLeftOuter,omitempty"`
	EyebrowLeftInner    Point `json:"eyebrowLeftInner,omitempty"`
	EyeLeftOuter        Point `json:"eyeLeftOuter,omitempty"`
	EyeLeftTop          Point `json:"eyeLeftTop,omitempty"`
	EyeLeftBottom       Point `json:"eyeLeftBottom,omitempty"`
	EyeLeftInner        Point `json:"eyeLeftInner,omitempty"`
	EyebrowRightOuter   Point `json:"eyebrowRightOuter,omitempty"`
	EyebrowRightInner   Point `json:"eyebrowRightInner,omitempty"`
	EyeRightOuter       Point `json:"eyeRightOuter,omitempty"`
	EyeRightTop         Point `json:"eyeRightTop,omitempty"`
	EyeRightBottom      Point `json:"eyeRightBottom,omitempty"`
	EyeRightInner       Point `json:"eyeRightInner,omitempty"`
	NoseRootLeft        Point `json:"noseRootLeft,omitempty"`
	NoseRootRight       Point `json:"noseRootRight,omitempty"`
	NoseLeftAlarTop     Point `json:"noseLeftAlarTop,omitempty"`
	NoseLeftAlaryOutTip Point `json:"noseLeftAlarOutTop,omitempty"`
	NoseRightAlarOutTip Point `json:"noseRightAlarOutTip,omitempty"`
	UpperLipTop         Point `json:"upperLipTop,omitempty"`
	UpperLipBottom      Point `json:"upperLipBottom,omitempty"`
	UnderLipTop         Point `json:"underLipTop,omitempty"`
	UnderLipBottom      Point `json:"underLipBottom,omitempty"`
}

//FaceAttribute represents the attributes of a given face.
type FaceAttributes struct {
	Age        float32    `json:"age,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Smile      float32    `json:"smile,omitempty"`
	FacialHair FacialHair `json:"facialHair,omitempty"`
	Glasses    string     `json:"glasses,omitempty"`
	HeadPose   HeadPose   `json:"headPose,omitempty"`
}

//FacialHair represents the facial hair detectd on a given face.
type FacialHair struct {
	Mustache  float32 `json:"mustache,omitempty"`
	Beard     float32 `json:"beard,omitempty"`
	Sideburns float32 `json:"sideburns,omitempty"`
}

//HeadPose represents the head pose of a given face.
type HeadPose struct {
	Roll  float32 `json:"roll,omitempty"`
	Yaw   float32 `json:"yaw,omitempty"`
	Pitch float32 `json:"pitch,omitempty"`
}

//Point represents an X and Y coordinate pair. This is primarily used
//to define FaceLandmark locations.
type Point struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
}
