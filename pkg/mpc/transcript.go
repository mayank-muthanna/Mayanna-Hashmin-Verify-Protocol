package mpc

type Transcript struct {
	Views []View
}

func NewTranscript(views []*View) Transcript {

	transcript := Transcript{}

	for _, view := range views {
		transcript.Views = append(transcript.Views, *view)
	}

	return transcript

}
