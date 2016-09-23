package goxford

//FaceDetectResponse is the structured response from a call to Detect.
type FaceDetectResponse struct {
	//TODO: Maybe change this to UUID?
	FaceID         string         `json:"faceId"`
	FaceRectangle  FaceRectangle  `json:"faceRectangle"`
	FaceLandmarks  FaceLandmarks  `json:"faceLandmarks"`
	FaceAttributes FaceAttributes `json:"faceAttributes"`
}

//FaceDetectRequest is the structured request to make a call to Detect.
type FaceDetectRequest struct {
	URL string `json:"url"`
}

//FaceRectangle represents the face rectangle detected on an image.
type FaceRectangle struct {
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
	Left   int32 `json:"left"`
	Top    int32 `json:"top"`
}

//FaceLandmarks represents the particular landmarks on a detected face.
type FaceLandmarks struct {
	PupilLeft           Point `json:"pupilLeft"`
	PupilRight          Point `json:"pupilRight"`
	NoseTip             Point `json:"noseTip"`
	MouthLeft           Point `json:"mouthLeft"`
	MouthRight          Point `json:"mouthRight"`
	EyebrowLeftOuter    Point `json:"eyebrowLeftOuter"`
	EyebrowLeftInner    Point `json:"eyebrowLeftInner"`
	EyeLeftOuter        Point `json:"eyeLeftOuter"`
	EyeLeftTop          Point `json:"eyeLeftTop"`
	EyeLeftBottom       Point `json:"eyeLeftBottom"`
	EyeLeftInner        Point `json:"eyeLeftInner"`
	EyebrowRightOuter   Point `json:"eyebrowRightOuter"`
	EyebrowRightInner   Point `json:"eyebrowRightInner"`
	EyeRightOuter       Point `json:"eyeRightOuter"`
	EyeRightTop         Point `json:"eyeRightTop"`
	EyeRightBottom      Point `json:"eyeRightBottom"`
	EyeRightInner       Point `json:"eyeRightInner"`
	NoseRootLeft        Point `json:"noseRootLeft"`
	NoseRootRight       Point `json:"noseRootRight"`
	NoseLeftAlarTop     Point `json:"noseLeftAlarTop"`
	NoseLeftAlaryOutTip Point `json:"noseLeftAlarOutTop"`
	NoseRightAlarOutTip Point `json:"noseRightAlarOutTip"`
	UpperLipTop         Point `json:"upperLipTop"`
	UpperLipBottom      Point `json:"upperLipBottom"`
	UnderLipTop         Point `json:"underLipTop"`
	UnderLipBottom      Point `json:"underLipBottom"`
}

//FaceAttribute represents the attributes of a given face.
type FaceAttributes struct {
	Age        float32    `json:"age"`
	Gender     string     `json:"gender"`
	Smile      float32    `json:"smile"`
	FacialHair FacialHair `json:"facialHair"`
	Glasses    string     `json:"glasses"`
	HeadPose   HeadPose   `json:"headPose"`
}

//FacialHair represents the facial hair detectd on a given face.
type FacialHair struct {
	Mustache  float32 `json:"mustache"`
	Beard     float32 `json:"beard"`
	Sideburns float32 `json:"sideburns"`
}

//HeadPose represents the head pose of a given face.
type HeadPose struct {
	Roll  float32 `json:"roll"`
	Yaw   float32 `json:"yaw"`
	Pitch float32 `json:"pitch"`
}

//Point represents an X and Y coordinate pair. This is primarily used
//to define FaceLandmark locations.
type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
